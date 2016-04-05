package main

import (
	"errors"
	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/fsnotify/fsnotify"
	"os"
	"path/filepath"
	"strings"
)

type Event fsnotify.Event

type RecursiveWatcher struct {
	*fsnotify.Watcher
	Files   chan string
	Folders chan string
}

var (
	Log            *logrus.Logger
	WatchRoot      string
	ModifiedFiles  chan string
	DeletedFiles   chan string
	WatchExit      chan bool
	FileManageExit chan bool
)

func (e Event) String() string {
	return fsnotify.Event(e).String()
}

func NewRecursiveWatcher(path string) (*RecursiveWatcher, error) {
	folders := Subfolders(path)
	if len(folders) == 0 {
		return nil, errors.New("No folder to watch.")
	}
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}
	rw := &RecursiveWatcher{Watcher: watcher}
	rw.Files = make(chan string, 10)
	rw.Folders = make(chan string, len(folders))
	for _, folder := range folders {
		if err = rw.AddFolder(folder); err != nil {
			return nil, err
		}
	}
	return rw, nil
}

func (watcher *RecursiveWatcher) AddFolder(folder string) error {
	if err := watcher.Add(folder); err != nil {
		return err
	}
	watcher.Folders <- folder
	return nil
}

func Subfolders(path string) (paths []string) {
	filepath.Walk(path, func(newPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			name := info.Name()
			if IgnoreFile(name) && name != "." && name != ".." {
				return filepath.SkipDir
			}
			paths = append(paths, newPath)
		}
		return nil
	})
	return paths
}

func IgnoreFile(name string) bool {
	return strings.HasPrefix(name, ".") || strings.HasPrefix(name, "_")
}

func Watch(path string) {
	Log.Debugf("Adding %s to watcher\n", path)
	watcher, err := NewRecursiveWatcher(path)
	if err != nil {
		Log.Fatalf("Could not initialize fsnotify watcher: %v\n", err)
	}
	WatchRoot = path
	for {
		select {
		case event := <-watcher.Events:
			Log.Debugln("Event:", event)
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
						Log.Debugln("Couldn't watch folder:", err)
					}
				}
			} else if event.Op&fsnotify.Rename == fsnotify.Rename {
				Log.Debugln("Rename event received:", event.Name)
				DeletedFiles <- path
			}
		case <-WatchExit:
			Log.Debugln("WatchExit signal received -- shutting down watcher")
			goto _cleanup
		}
	}
_cleanup:
	if err := watcher.Close(); err != nil {
		Log.Errorln("Error shutting down watcher:", err)
	}

	Log.Debugln("Exiting Watcher")
	WatchExit <- true
}

func filePathFromEvent(event *fsnotify.Event) (path string, err error) {
	defer func() { path = filepath.Clean(path) }()

	if filepath.IsAbs(event.Name) {
		path = event.Name
		return
	}

	path, err = filepath.Abs(event.Name)
	return
}

func FileManage() {
	for {
		select {
		case file := <-ModifiedFiles:
			Log.Infoln("Modified file:", file)
		case file := <-DeletedFiles:
			Log.Infoln("Delete file:", file)
		case <-FileManageExit:
			Log.Debugln("FileManage exit signal received")
			goto _cleanup
		}
	}
_cleanup:
	close(ModifiedFiles)
	close(DeletedFiles)
	FileManageExit <- true
}
func main() {
	app := cli.NewApp()
	app.Action = watch
	app.Name = "Watch"
	app.Usage = "Watch folder"
	app.Run(os.Args)
}

func watch(context *cli.Context) {
	argc := len(context.Args())
	if argc < 1 || argc > 2 {
		context.App.Command("help").Run(context)
		return
	}

	watchDir, err := os.Getwd()
	if err != nil {
		Log.Fatalln(err)
	}
	if argc > 1 {
		watchDir = context.Args()[0]
	}

	Log.Debugln("Starting watcher goroutine with watchdir", watchDir)
	go Watch(watchDir)
	go FileManage()
	<-WatchExit
	Log.Debug("Exiting")
}
