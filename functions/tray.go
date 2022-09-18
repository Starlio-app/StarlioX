package functions

import (
	"fmt"
	"github.com/getlantern/systray"
	"github.com/shirou/gopsutil/v3/process"
	"github.com/skratchdot/open-golang/open"
	"os"
)

func Tray() {
	systray.SetIcon(GetIcon("web/static/image/icons/everynasa.ico"))
	systray.SetTitle("EveryNasa")
	systray.SetTooltip("EveryNasa")
	ui := systray.AddMenuItem("Open Ui", "Open Ui")
	exit := systray.AddMenuItem("Quit", "Quit the whole app")

	for {
		select {
		case <-ui.ClickedCh:
			err := open.Run("http://localhost:4662")
			if err != nil {
				panic(err)
			}
		case <-exit.ClickedCh:
			Quit()
		}
	}
}

func Quit() {
	err := KillProcess("EveryNasa.exe")
	if err != nil {
		panic(err)
	}

	systray.Quit()
}

func KillProcess(name string) error {
	processes, err := process.Processes()
	if err != nil {
		panic(err)
	}

	for _, p := range processes {
		n, ProccErr := p.Name()
		if ProccErr != nil {
			panic(ProccErr)
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
		fmt.Print(err)
	}
	return b
}
