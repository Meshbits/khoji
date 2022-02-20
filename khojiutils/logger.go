package khojiutils

import (
	"log"
	"os"
)

var (
	// Log ...
	Log *log.Logger
	// KhojiLogFile stores the stdout/stderr output log
	KhojiLogFile = "./khoji.log"
)

func init() {
	// set location of log file
	var logpath = KhojiLogFile

	// flag.Parse()
	var file, err1 = os.Create(logpath)

	if err1 != nil {
		panic(err1)
	}
	Log = log.New(file, "", log.LstdFlags|log.Lshortfile)
	// Log.Println("LogFile : " + logpath)
}
