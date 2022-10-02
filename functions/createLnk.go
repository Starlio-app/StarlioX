package functions

import (
	"os"

	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

func CreateLnk(src, dst string) error {
	err := ole.CoInitializeEx(0, ole.COINIT_APARTMENTTHREADED|ole.COINIT_SPEED_OVER_MEMORY)
	if err != nil {
		Logger(err.Error())
	}

	oleShellObject, err := oleutil.CreateObject("WScript.Shell")
	if err != nil {
		Logger(err.Error())
	}

	defer oleShellObject.Release()
	wshell, err := oleShellObject.QueryInterface(ole.IID_IDispatch)
	if err != nil {
		Logger(err.Error())
	}

	defer wshell.Release()
	cs, err := oleutil.CallMethod(wshell, "CreateShortcut", dst)
	if err != nil {
		Logger(err.Error())
	}

	idispatch := cs.ToIDispatch()
	_, err = oleutil.PutProperty(idispatch, "TargetPath", src)
	if err != nil {
		Logger(err.Error())
	}

	dir, err := os.Getwd()
	if err != nil {
		Logger(err.Error())
	}

	_, err = oleutil.PutProperty(idispatch, "WorkingDirectory", dir)
	if err != nil {
		Logger(err.Error())
	}

	_, err = oleutil.CallMethod(idispatch, "Save")
	if err != nil {
		Logger(err.Error())
	}

	return nil
}
