from win10toast import ToastNotifier

from bs4 import BeautifulSoup
from elevate import elevate
import requests
import ctypes
import os
import urllib
import schedule
import winreg as reg
import getpass


class Nasa:

    def __init__(self):
        self.url = "https://apod.nasa.gov/apod/"
        self.headers = {'User-Agent': 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_3) '
                                      'AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.149 Safari/537.36'}
        self.photoName = "everydayphotonasa.jpg"
        self.toaster = ToastNotifier()

    def autorun(self):
        path = os.path.dirname(os.path.realpath(__file__))
        address = os.path.join(path, "main.exe")
        key_value = "Software/Microsoft/Windows/CurrentVersion/Run"
        user = getpass.getuser()
        key = reg.OpenKey(reg.HKEY_LOCAL_MACHINE, key_value, 0, reg.KEY_ALL_ACCESS)
        reg.SetValueEx(key, user, 0, reg.REG_SZ, address)
        reg.CloseKey(key)
        self.toaster.show_toast("EveryDayPhotoNasa",
                                "Программа добавлена в автозапуск.",
                                duration=5,
                                icon_path=None)

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
                                           duration=5,
                                           icon_path=None)

    def set_wallpaper(self):
        path = os.path.abspath(self.photoName)
        ctypes.windll.user32.SystemParametersInfoW(20, 0, path, 0)
        self.toaster.show_toast("EveryDayPhotoNasa",
                                "Обои поставлены.",
                                duration=5,
                                icon_path=None)

    def start(self):
        self.download_photo()
        if ctypes.windll.shell32.IsUserAnAdmin() != 0:
            elevate(show_console=False, graphical=False)
            self.autorun()


if __name__ == "__main__":
    nasa = Nasa()
    nasa.start()
    schedule.every(1).day.do(nasa.start)
    while True:
        schedule.run_pending()
