package functions

import (
	"fmt"
	"os"

	"github.com/getlantern/systray"
	"github.com/shirou/gopsutil/v3/process"
	"github.com/skratchdot/open-golang/open"
)

func Tray() {
	systray.SetIcon(GetIcon("web/static/image/icons/icon.ico"))
	systray.SetTitle("EveryNasa")
	systray.SetTooltip("EveryNasa")
	ui := systray.AddMenuItem("Open UI", "Open UI")
	exit := systray.AddMenuItem("Quit", "Quit the whole app")

	for {
		select {
		case <-ui.ClickedCh:
			err := open.Run("http://localhost:4662")
			if err != nil {
				Logger(err.Error())
			}
		case <-exit.ClickedCh:
			Quit()
		}
	}
}

func Quit() {
	err := KillProcess("EveryNasa.exe")
	if err != nil {
		Logger(err.Error())
	}

	systray.Quit()
}

func KillProcess(name string) error {
	processes, err := process.Processes()
	if err != nil {
		Logger(err.Error())
	}

	for _, p := range processes {
		n, err := p.Name()
		if err != nil {
			Logger(err.Error())
		}

		if n == name {
			return p.Kill()
		}
	}
	return fmt.Errorf("process not found")
}

func GetIcon(s string) []byte {
	b, err := os.ReadFile(s)
	if err != nil {
		Logger(err.Error())
	}
	return b
}
