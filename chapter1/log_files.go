package main

import (
	"log"
	"log/syslog"
	"os"
	"path/filepath"
)

func main() {
	program_name := filepath.Base(os.Args[0])
	sys_log, err := syslog.New(syslog.LOG_INFO|syslog.LOG_LOCAL7, program_name)

	if err != nil {
		log.Fatalln(err)
	} else {
		log.SetOutput(sys_log)
	}

	log.Println("Test logging!!!")
}
