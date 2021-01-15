package main

import (
	"context"
	"log"
	"os"

	commandjsonio "github.com/go-marshaltemabu/go-commandjsonio"
)

type helloData struct {
	InputData  *helloData `json:"input"`
	CmdArgs    []string   `json:"cmd"`
	ClockSec   float64    `json:"clock"`
	CycleCount int32
}

func main() {
	if len(os.Args) < 2 {
		log.Print("Usage: ./echoj /bin/cat ...")
		return
	}
	exePath := os.Args[1]
	cmdArgs := os.Args[2:]
	dataBuf := helloData{
		CmdArgs: []string{"Hello", "World"},
	}
	cmd := commandjsonio.NewCommandJSONReaderWriter(context.Background(), exePath, cmdArgs, nil, "")
	if err := cmd.Start(); nil != err {
		log.Fatalf("ERROR: start failed: %v", err)
		return
	}
	for cycleCnt := int32(0); cycleCnt < 5; cycleCnt++ {
		dataBuf.CycleCount = cycleCnt
		if err := cmd.Write(&dataBuf); nil != err {
			log.Fatalf("ERROR: write failed (cycle=%d): %v", cycleCnt, err)
		}
		log.Printf("INFO: written cycle %d.", cycleCnt)
		dataBuf = helloData{}
		if err := cmd.Read(&dataBuf); nil != err {
			log.Fatalf("ERROR: read failed (cycle=%d): %v", cycleCnt, err)
		}
		log.Printf("INFO: output (cycle %d) = %#v; %#v", cycleCnt, &dataBuf, dataBuf.InputData)
	}
	if err := cmd.CloseStdin(); nil != err {
		log.Printf("ERROR: close stdin failed: %v", err)
	}
	if err := cmd.Wait(); nil != err {
		log.Printf("ERROR: wait failed: %v", err)
	}
}
