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

    @staticmethod
    def autorun():
        path = os.path.dirname(os.path.realpath(__file__))
        address = os.path.join(path, "main.py")
        key_value = "Software/Microsoft/Windows/CurrentVersion/Run"
        user = getpass.getuser()
        key = reg.OpenKey(reg.HKEY_LOCAL_MACHINE, key_value, 0, reg.KEY_ALL_ACCESS)
        reg.SetValueEx(key, user, 0, reg.REG_SZ, address)
        reg.CloseKey(key)
        print("Программа добавлена в автозапуск")

    def download_photo(self):
        full_page = requests.get(self.url, headers=self.headers)
        soup = BeautifulSoup(full_page.content, 'html.parser')
        lnk = str
        for link in soup.select("img"):
            lnk = link["src"]
            print(f"Сохарняю картинку — {self.url + lnk}")

        img = urllib.request.urlopen(self.url + lnk).read()
        out = open(self.photoName, "wb")
        out.write(img)
        out.close()

    def set_wallpaper(self):
        path = os.path.abspath(self.photoName)
        print("Установлено фоновое изображение")
        ctypes.windll.user32.SystemParametersInfoW(20, 0, path, 0)

    def start(self):
        self.download_photo()
        self.set_wallpaper()
        print("Можно закрывать программу")
        if ctypes.windll.shell32.IsUserAnAdmin() != 0:
            elevate(show_console=False, graphical=False)
            self.autorun()
            print("Программа добавлена в автозапуск, можете закрывать программу")


if __name__ == "__main__":
    nasa = Nasa()
    nasa.start()
    schedule.every(1).day.do(nasa.start)
    while True:
        schedule.run_pending()
