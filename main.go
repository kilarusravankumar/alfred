package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/fsnotify/fsnotify"
) 


func main(){
    fmt.Println("Alfred will watchover your diretory and run go build after an file write event.")
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
                    fmt.Println("triggered event: ", event)
                    if event.Has(fsnotify.Write){
                        if isNotGitAndExeFile(event.Name) {
                            runBuildCommand()
                            fmt.Printf("%s file is modified", event.Name)

                        }
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
        fmt.Println(err.Error())
    }

    <- make(chan struct{})

}


func runBuildCommand(){
    fmt.Println("Building go application.")
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
    if strings.Contains(filename, ".git") && strings.Contains(filename, ".exe"){
        return false
    }
   return true 
}
