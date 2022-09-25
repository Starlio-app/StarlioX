package controllers

import (
	"github.com/Redume/EveryNasa/api/utils"
	"github.com/Redume/EveryNasa/functions"
	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
	"net/http"
	"os"
	"os/user"
	"strings"
)

var SetStartup = func(w http.ResponseWriter, r *http.Request) {
	u, err := user.Current()
	if err != nil {
		functions.Logger(err.Error())
	}

	dir, err := os.Getwd()
	if err != nil {
		functions.Logger(err.Error())
	}

	dir = strings.Replace(dir, "\\", "\\\\", -1) + "\\EveryNasa.exe"

	err = makeLnk(dir, strings.Replace(u.HomeDir, "\\", "\\\\", -1)+"\\AppData\\Roaming\\Microsoft\\Windows\\Start Menu\\Programs\\Startup\\EveryNasa.lnk")
	if err != nil {
		functions.Logger(err.Error())
	}

	utils.Respond(w, utils.Message(true, "The settings have been applied successfully."))
}

var RemoveStartup = func(w http.ResponseWriter, r *http.Request) {
	u, err := user.Current()
	if err != nil {
		functions.Logger(err.Error())
	}

	err = os.Remove(strings.Replace(u.HomeDir, "\\", "\\\\", -1) + "\\AppData\\Roaming\\Microsoft\\Windows\\Start Menu\\Programs\\Startup\\EveryNasa.lnk")
	if err != nil {
		functions.Logger(err.Error())
	}

	utils.Respond(w, utils.Message(true, "The settings have been applied successfully."))
}

func makeLnk(src, dst string) error {
	err := ole.CoInitializeEx(0, ole.COINIT_APARTMENTTHREADED|ole.COINIT_SPEED_OVER_MEMORY)
	if err != nil {
		functions.Logger(err.Error())
	}

	oleShellObject, err := oleutil.CreateObject("WScript.Shell")
	if err != nil {
		functions.Logger(err.Error())
	}

	defer oleShellObject.Release()
	wshell, err := oleShellObject.QueryInterface(ole.IID_IDispatch)
	if err != nil {
		functions.Logger(err.Error())
	}

	defer wshell.Release()
	cs, err := oleutil.CallMethod(wshell, "CreateShortcut", dst)
	if err != nil {
		functions.Logger(err.Error())
	}

	idispatch := cs.ToIDispatch()
	_, err = oleutil.PutProperty(idispatch, "TargetPath", src)
	if err != nil {
		functions.Logger(err.Error())
	}

	dir, err := os.Getwd()
	if err != nil {
		functions.Logger(err.Error())
	}

	_, err = oleutil.PutProperty(idispatch, "WorkingDirectory", dir)
	if err != nil {
		functions.Logger(err.Error())
	}

	_, err = oleutil.CallMethod(idispatch, "Save")
	if err != nil {
		functions.Logger(err.Error())
	}

	return nil
}
