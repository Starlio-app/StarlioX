"""
   Copyright 2022 Redume

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
"""


from bs4 import BeautifulSoup
from win10toast import ToastNotifier
from pystray import MenuItem, Menu, Icon
from PIL import Image
from elevate import elevate

import json
import requests
import ctypes
import os
import urllib
import schedule
import winreg as reg
import getpass
import sys

from config import Config


class Nasa(Config):
    def __init__(self):
        self.url = "https://apod.nasa.gov/apod/"
        self.headers = {'User-Agent': 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_3) '
                                      'AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.149 Safari/537.36'}
        self.photoName = "everydayphotonasa.jpg"
        self.toaster = ToastNotifier()
        self.state = json.loads(
                                Config.get_setting(self,
                                                   path="config.ini",
                                                   section="Settings",
                                                   setting="autorun")
                                .lower())

    def resource_path(self, relative_path):
        base_path = getattr(sys, '_MEIPASS', os.path.dirname(os.path.abspath(__file__)))
        return os.path.join(base_path, relative_path)

    def autorun(self):
        self.state = not MenuItem.checked

        if not ctypes.windll.shell32.IsUserAnAdmin() != 0:
            elevate(show_console=False)

        if self.state is False:
            print("Автозапуск включен.")
            path = os.path.dirname(os.path.realpath(__file__))
            address = os.path.join(path, "start.bat")
            key_value = "SOFTWARE\Microsoft\Windows\CurrentVersion\Run"
            user = getpass.getuser()
            key = reg.OpenKey(reg.HKEY_LOCAL_MACHINE, key_value, 0, reg.KEY_ALL_ACCESS)
            reg.SetValueEx(key, user, 0, reg.REG_SZ, address)
            reg.CloseKey(key)
            Config.update_setting(self,
                                  path="config.ini",
                                  section="Settings",
                                  setting="autorun",
                                  value="True")

            self.toaster.show_toast("EveryDayPhotoNasa",
                                    "Программа добавлена в автозапуск.",
                                    duration=3,
                                    icon_path=self.resource_path("nasa.ico"))
        else:
            print("Программа удалена из автозапуска.")
            key_value = "SOFTWARE\Microsoft\Windows\CurrentVersion\Run"
            user = getpass.getuser()
            key = reg.OpenKey(reg.HKEY_LOCAL_MACHINE, key_value, 0, reg.KEY_ALL_ACCESS)
            reg.DeleteValue(key, user)
            reg.CloseKey(key)
            Config.update_setting(self,
                                  path="config.ini",
                                  section="Settings",
                                  setting="autorun",
                                  value="False")

            self.toaster.show_toast("EveryDayPhotoNasa",
                                    "Программа удалена из автозапуска.",
                                    duration=3,
                                    icon_path=self.resource_path("nasa.ico"))

    def tray(self):
        tray = Icon("EveryDayPhotoNasa", title="EveryDayPhotoNasa",
                    icon=Image.open(self.resource_path("nasa.ico")),
                    menu=Menu(
                        MenuItem("Автозапуск",
                                 self.autorun,
                                 checked=lambda item: self.state),
                        MenuItem("Выход",
                                 self.tray_close),
                    ))
        tray.run()

    def tray_close(self, tray):
        tray.visible = False
        sys.exit(0)

    def download_photo(self):
        try:
            full_page = requests.get(self.url, headers=self.headers)
            soup = BeautifulSoup(full_page.content, 'html.parser')
            lnk = str
            for link in soup.select("img"):
                lnk = link["src"]

            img = urllib.request.urlopen(self.url + lnk).read()
            out = open(self.photoName, "wb")
            out.write(img)
            out.close()
            self.set_wallpaper()
        except requests.exceptions.ConnectionError:
            return self.toaster.show_toast("EveryDayPhotoNasa",
                                           "Не получилось подключится к сайту, проверьте подключение к интернету.",
                                           duration=3,
                                           icon_path=self.resource_path("nasa.ico"))

    def set_wallpaper(self):
        print(self.state)
        path = os.path.abspath(self.photoName)
        ctypes.windll.user32.SystemParametersInfoW(20, 0, path, 0)
        self.toaster.show_toast("EveryDayPhotoNasa",
                                "Обои установлены.",
                                duration=3,
                                icon_path=self.resource_path("nasa.ico"))


nasa = Nasa()
nasa.download_photo()

schedule.every().day.at("00:30").do(nasa.download_photo)
while True:
    schedule.run_pending()
    nasa.tray()
