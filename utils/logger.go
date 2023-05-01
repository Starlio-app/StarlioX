package utils

import (
	"fmt"
	"gopkg.in/toast.v1"
	"os"
	"strconv"
	"time"
)

func Log(text string) {
	timestamps := time.Now()
	logFile := string(timestamps.Format("02.01.06")) + ".log"

	if !FileExists("log") {
		os.Mkdir("log", 0755)
	}
	if !FileExists("log/" + logFile) {
		file, _ := os.Create("log/" + logFile)

		defer func(file *os.File) {
			file.Close()
		}(file)
	}

	file, _ := os.OpenFile("log/"+logFile, os.O_APPEND|os.O_WRONLY, 0600)
	fileText := timestamps.Format("15:04:05") + " | " + text + "\n"

	file.Write([]byte(fileText))
	fmt.Println(fileText)

	dir, _ := os.Getwd()
	dir = dir + "\\log\\" + logFile
	fmt.Println("Error has been save to " + dir)

	Notify(
		"There was an error",
		"An unknown error occurred, the error was recorded in the log file",
		toast.Action{Type: "protocol", Label: "View log file", Arguments: dir},
	)
}

func CheckLogs() {
	if !FileExists("log") {
		return
	}

	files, err := os.ReadDir("log")
	if err != nil {
		panic(err)
	}

	for _, f := range files {
		fileTimestamps, _ := strconv.Atoi(f.Name()[:2])
		timestamps := time.Now().Day()

		if timestamps-fileTimestamps >= 7 {
			os.Remove("log/" + f.Name())
		}
	}
}
func FileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
