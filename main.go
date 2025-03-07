package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/fsnotify/fsnotify"
)

func main() {
	logger = InitLogger()
	logger.Info("Alfred 0.2.1 will watchover your diretory and run go build automatically")
	config := ParseConfig()
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic(err)
	}

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				logger.Debug("triggered event: " + event.Name)
				logger.Debug(fmt.Sprintf("%v is a go file", isWatchedFile(event.Name, config)))
				if isWatchedFile(event.Name, config) {
					runBuildCommand(config)
					logger.Debug(fmt.Sprintf("%s file is modified", event.Name))

				}
			case eventErr, ok := <-watcher.Errors:
				if !ok {
					return
				}
				logger.Error(eventErr.Error())
			}
		}

	}()

	if err := watcher.Add(getCurrentDirectory()); err != nil {
		logger.Error(err.Error())
	}

	<-make(chan struct{})

}

func runBuildCommand(config Config) {
	logger.Info("Building go application.")
	var cmd *exec.Cmd
    commandPath := strings.Split(config.BuildCommand," ")
	if config.BuildName == "" {
        cmd = exec.Command(commandPath[0],commandPath[1:]...)
	} else {
        cmd = exec.Command(commandPath[0],commandPath[1:]...)
	}
	if err := cmd.Run(); err != nil {
		fmt.Println(err.Error())
	}
}

func getCurrentDirectory() string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return dir
}

func isWatchedFile(filename string, config Config) bool {
    isWatched := false
    for _, ext := range config.WatchFiles {
        if strings.Contains(filename, ext) {
            isWatched = true
        }
    }
	return isWatched 
}
