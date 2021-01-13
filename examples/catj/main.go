package main

import (
	"log"
	"os"
	"time"

	commandjsonio "github.com/go-marshaltemabu/go-commandjsonio"
)

type helloData struct {
	Text  string
	Int64 int64
}

func main() {
	if len(os.Args) < 2 {
		log.Print("Usage: ./catj /bin/cat ...")
		return
	}
	cmdArgs := os.Args[1:]
	dataIn := &helloData{
		Text:  "Hello World",
		Int64: time.Now().Unix(),
	}
	var dataOut helloData
	err := commandjsonio.RunCommandJSON(nil, cmdArgs, nil, "", 10*time.Second, dataIn, &dataOut)
	if nil != err {
		log.Printf("ERROR: cannot run command: %v", err)
	}
	log.Printf("INFO: output = %#v", &dataOut)
}
