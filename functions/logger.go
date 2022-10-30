package functions

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"
)

func Logger(text string) {
	if !FileExists("logger.log") {
		err := CreateFile("logger.log")
		if err != nil {
			panic(err)
		}
	}
	f, err := os.OpenFile("logger.log", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	now := time.Now()

	_, file, line, ok := runtime.Caller(1)

	if !ok {
		file = "???"
	}

	dir, err := os.Getwd()
	if err != nil {
		Logger(err.Error())
	}

	if strings.Contains(file, dir) {
		file = strings.Replace(file, dir, "", -1)
	}

	var lineString string
	if ok {
		lineString = fmt.Sprintf("%d", line)
	}

	_, err = f.Write([]byte(now.Format("Mon Jan 2 15:04:05 2006") + " | " + text + " [" + file + "] [" + lineString + "]\n"))
	if err != nil {
		panic(err)
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

func CreateFile(name string) error {
	fo, err := os.Create(name)
	if err != nil {
		return err
	}
	defer func() {
		err := fo.Close()
		if err != nil {
			return
		}
	}()
	return nil
}
