package main

import (
	"fmt"
	"log"
	"os"
)

var LOGFILE = "./custom.log"

func main() {
	file, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	ilog := log.New(file, "customLogLineNumber ", log.LstdFlags)

	ilog.SetFlags(log.LstdFlags | log.Lshortfile)
	ilog.Println("Test")
	ilog.Println("Another test")
	for i := 0; i < 8; i++ {
		ilog.Println("Another test", i)
	}
}
