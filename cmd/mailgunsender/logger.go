package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

var date string = fmt.Sprintf("%d-%d-%d", time.Now().Year(), time.Now().Month(), time.Now().Day())

// Naive log writing function. Writes in the standard output aswell as in the file.
func WriteLog(level string, message ...any) {
	fileInfo, _ := openLogFile(fmt.Sprintf("./%s.log", date))

	fileLog := log.New(fileInfo, fmt.Sprintf("%s: ", level), log.Ldate|log.Ltime)
	log.SetOutput(fileInfo)
	fileLog.Print(message...)

	stdoutLog := log.New(os.Stdout, fmt.Sprintf("%s: ", level), log.Ldate|log.Ltime)
	log.SetOutput(os.Stdout)
	stdoutLog.Print(message...)
	defer fileInfo.Close()
}

// Opens log file. Deffering moved to writer function.
func openLogFile(path string) (*os.File, error) {
	logFile, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	return logFile, nil
}
