package functions

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/rodkranz/fetch"
)

func Logger(text string) {
	now := time.Now()

	file := string(now.Format("02.01.2006")) + ".log"

	if getDatabase() == 0 {
		return
	}

	if !FileExists("logs") {
		err := CreateFolder("logs")
		if err != nil {
			panic(err)
		}
	}

	if !FileExists("./logs/" + file) {
		err := CreateFile(file)
		if err != nil {
			panic(err)
		}
	}

	f, err := os.OpenFile("logs/"+file, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

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

	_, err = f.Write([]byte(now.Format("15:04:05") + " | " + text + " [" + file + "] [" + lineString + "]\n"))
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
	file, err := os.Create("logs/" + name)
	if err != nil {
		return nil
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	return nil
}

func CreateFolder(name string) error {
	err := os.Mkdir(name, 0755)
	if err != nil {
		return err
	}
	return nil
}

func getDatabase() int {
	client := fetch.NewDefault()
	res, err := client.Get("http://localhost:3000/api/get/settings", nil)
	if err != nil {
		panic(err)
	}

	body, err := res.ToString()
	if err != nil {
		panic(err)
	}

	type DatabaseStruct struct {
		Save_logg int `json:"save_logg"`
	}

	var Database DatabaseStruct
	err = json.Unmarshal([]byte(body), &Database)
	if err != nil {
		panic(err)
	}

	return Database.Save_logg
}
