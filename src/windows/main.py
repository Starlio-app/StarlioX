from win10toast import ToastNotifier
from infi.systray import SysTrayIcon

import schedule
import os
import sys
import psutil
import time

path = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))
sys.path.insert(0, path)

# noinspection PyUnresolvedReferences
from functions.wallpaper import Wallpaper


class Nasa(Wallpaper):
    def __init__(self):
        super().__init__()
        self.toaster = ToastNotifier()

    @staticmethod
    def resource_path(relative_path):
        base_path = getattr(sys, '_MEIPASS', os.path.dirname(os.path.abspath(__file__)))
        return os.path.join(base_path, relative_path)

    @staticmethod
    def kill_program(tray):
        for proc in psutil.process_iter():
            if proc.name() == "EveryNasa.exe":
                proc.kill()

    def tray(self):
        tray = SysTrayIcon(
            self.resource_path("./icons/icon.ico"),
            "EveryNasa",
            on_quit=self.kill_program,
        )
        tray.start()

    def main(self):
        self.tray()
        wall_check = Wallpaper.check(self)
        Wallpaper.download(self)
        wall_set = Wallpaper.set()
        self.toaster.show_toast("EveryNasa",
                                wall_check or wall_set,
                                duration=4,
                                icon_path=self.resource_path("./icons/icon.ico"))


nasa = Nasa()
nasa.main()

schedule.every(3).hours.do(nasa.main)
while True:
    schedule.run_pending()
    time.sleep(1)
