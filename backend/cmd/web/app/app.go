package app

import (
	"log"
	"os"
)

var (
	ErrorLog *log.Logger
	InfoLog  *log.Logger
)

func Init() {
	InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	ErrorLog = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
}

func Close() {
}
