package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/landaire/recwatch"
	"gopkg.in/fsnotify.v1"
	"os"
	"path/filepath"
	"strings"
)

var (
	Log           *logrus.Logger
	WatchExit     chan bool
	watchRoot     string
	ModifiedFiles chan string
	DeletedFiles  chan string
)

func init() {
	Log = logrus.New()
	WatchExit = make(chan bool)
	ModifiedFiles = make(chan string, 100)
	DeletedFiles = make(chan string, 100)
}

func Watch(path string) {
	Log.Debugf("Adding %s to watcher...\n", path)
	watcher, err := recwatch.NewRecursiveWatcher(path)
	if err != nil {
		Log.Fatalf("can not initialize fsnotify watcher: %v\n", err)
	}
	watchRoot = path
	for {
		select {
		case event := <-watcher.Events:
			Log.Debugln("Event: ", event)
			if strings.TrimSpace(event.Name) == "" {
				Log.Debugln("Got an empty string... not touching that")
				continue
			}
			path, err := filePathFromEvent(&event)
			if err != nil {
				Log.Errorln(err)
				continue
			}
			if event.Op&fsnotify.Write == fsnotify.Write {
				Log.Debugln("Write event received:", event.Name)
				ModifiedFiles <- path
			} else if event.Op&fsnotify.Remove == fsnotify.Remove {
				Log.Debugln("Remove event received:", event.Name)
				DeletedFiles <- path
			} else if event.Op&fsnotify.Create == fsnotify.Create {
				Log.Debugln("Create event received:", event.Name)
				ModifiedFiles <- path
				stat, err := os.Stat(path)
				if err != nil {
					Log.Errorln(err)
					continue
				}
				if stat.IsDir() {
					if err := watcher.Add(path); err != nil {
						Log.Debugln("Can not watch folder:", err)
					}
				}
			} else if event.Op&fsnotify.Rename == fsnotify.Rename {
				Log.Debugln("Rename event received:", event.Name)
				DeletedFiles <- path
			}
		case <-WatchExit:
			Log.Debugln("WatchExit signle received -- shutting down watcher")
			goto _cleanup
		}
	}
_cleanup:
	if err := watcher.Close(); err != nil {
		Log.Errorln("Error shutting down watcher", err)
	}
	Log.Debugln("Exiting Watcher")
	WatchExit <- true
}
func filePathFromEvent(event *fsnotify.Event) (path string, err error) {
	defer func() {
		path = filepath.Clean(path)
	}()
	if filepath.IsAbs(event.Name) {
		path = event.Name
		return
	}
	path, err = filepath.Abs(event.Name)
	return
}

func main() {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "verbose",
			Usage: "Enable verbose logging",
		},
	}
	app.Action = watch
	app.Name = "watcher"
	app.Run(os.Args)
}
