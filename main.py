from bs4 import BeautifulSoup
import requests

import ctypes
import os
import urllib
import schedule


class Nasa:

    def __init__(self):
        self.url = "https://apod.nasa.gov/apod/"
        self.headers = {'User-Agent': 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_3) '
                                      'AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.149 Safari/537.36'}
        self.photoName = "everydayphotonasa.jpg"

    def download_photo(self):
        full_page = requests.get(self.url, headers=self.headers)
        soup = BeautifulSoup(full_page.content, 'html.parser')
        lnk = str
        for link in soup.select("img"):
            lnk = link["src"]
            print(f"Скачиваю картинку — {self.url + lnk}")

        img = urllib.request.urlopen(self.url + lnk).read()
        out = open(self.photoName, "wb")
        out.write(img)
        out.close()

    def set_wallpaper(self):
        path = os.path.abspath(self.photoName)
        ctypes.windll.user32.SystemParametersInfoW(20, 0, path, 0)

    def start(self):
        self.download_photo()
        self.set_wallpaper()


if __name__ == "__main__":
    schedule.every(1).days.do(Nasa().start)
    while True:
        schedule.run_pending()
