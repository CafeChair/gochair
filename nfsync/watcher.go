package main

import (
	"github.com/codegangsta/cli"
	"gopkg.in/fsnotify.v1"
	"log"
	"os"
	"strings"
)

var watchExit chan bool

func watch(path string) {
	log.Printf("Adding %s to watcher\n", path)
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatalf("Could not initialize fsnotify watcher: %v\n", err)
	}
	defer watcher.Close()

	go func() {
		for {
			select {
			case event := <-watcher.Events:
				log.Printf("Event:", event)
				if strings.TrimSpace(event.Name) == "" {
					log.Println("Got an empty string... not touching that")
					continue
				}
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Printf("Write event received: %s", event.Name)
				} else if event.Op&fsnotify.Remove == fsnotify.Remove {
					log.Printf("Remove event received: %s", event.Name)
				} else if event.Op&fsnotify.Create == fsnotify.Create {
					log.Printf("Create event received: %s", event.Name)
				} else if event.Op&fsnotify.Rename == fsnotify.Rename {
					log.Printf("Rename event received: %s", event.Name)
				}
			case err := <-watcher.Errors:
				log.Printf("Error %s\n", err)
			}
		}
	}()
	err = watcher.Add(path)
	if err != nil {
		log.Fatalln(err)
	}
	<-watchExit
}

func Watcher(context *cli.Context) {
	argc := len(context.Args())
	if argc < 1 || argc > 2 {
		context.App.Command("help").Run(context)
		return
	}
	watchDir, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	if argc > 1 {
		watchDir = context.Args()[0]
	}
	watch(watchDir)
	log.Println("Exiting...")
}

func main() {
	app := cli.NewApp()
	app.Action = Watcher
	app.Run(os.Args)
	watch("/tmp")
}
