from pystray import MenuItem, Menu, Icon
from PIL import Image

import notify2
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

    @staticmethod
    def resource_path(relative_path):
        base_path = getattr(sys, '_MEIPASS', os.path.dirname(os.path.abspath(__file__)))
        return os.path.join(base_path, relative_path)

    def tray(self):
        tray = Icon("EveryNasa",
                    title="EveryNasa",
                    icon=Image.open(self.resource_path("./icons/icon.ico")),
                    menu=Menu(
                        MenuItem("Выход",
                                 self.kill_program),
                    ))
        tray.run()

    def kill_program(self):
        for proc in psutil.process_iter():
            if proc.name() == "EveryNasa":
                proc.kill()

    @staticmethod
    def notify(title, message):
        notify2.init("EveryNasa")
        notice = notify2.Notification(title, message)
        notice.show()
        return

    def main(self):
        wall_check = Wallpaper.check(self)
        Wallpaper.download(self)
        wall_set = Wallpaper.set()
        self.notify("EveryNasa", wall_check or wall_set)


nasa = Nasa()
nasa.main()

schedule.every(3).hours.do(nasa.main)
while True:
    schedule.run_pending()
    nasa.tray()
    time.sleep(1)
