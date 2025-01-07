package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/fsnotify/fsnotify"
) 

func main(){
    logger = InitLogger()
    logger.Info("Alfred 0.2 will watchover your diretory and run go build after an file write event.")
    
    watcher, err := fsnotify.NewWatcher()
    if err != nil {
        panic(err)
    }

    go func(){
        for{
           select{
                case event, ok := <- watcher.Events:
                    if !ok {
                        return 
                    }
                    logger.Debug("triggered event: "+ event.Name)
                    logger.Debug(fmt.Sprintf("%v is a go file", isNotGitAndExeFile(event.Name)))
                    if isNotGitAndExeFile(event.Name) {
                        runBuildCommand()
                        logger.Debug(fmt.Sprintf("%s file is modified", event.Name))

                    }
                case err, ok := <- watcher.Events:
                    if !ok {
                        return
                    }
                    fmt.Println(err)
           } 
        } 

    }()
   
    if err := watcher.Add(getCurrentDirectory()); err != nil {
        logger.Error(err.Error())
    }

    <- make(chan struct{})

}


func runBuildCommand(){
    logger.Info("Building go application.")
    cmd := exec.Command("go","build")
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

func isNotGitAndExeFile(filename string) bool{
    if strings.Contains(filename, ".git") || strings.Contains(filename, ".exe"){
        return false
    }
   return true 
}
