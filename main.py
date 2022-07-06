import math
from bs4 import BeautifulSoup
import requests

import ctypes
import os
import urllib

from PIL import Image
from screeninfo import get_monitors


class Nasa:

    def __init__(self):
        self.url = "https://apod.nasa.gov/apod/"
        self.headers = {'User-Agent': 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_3) '
                                      'AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.149 Safari/537.36'}
        self.photoName = "everydayphotonasa.jpg"
        self.screen_size = (self.get_size_screen())

    def download_photo(self):
        full_page = requests.get(self.url, headers=self.headers)
        soup = BeautifulSoup(full_page.content, 'html.parser')
        for link in soup.select("img"):
            lnk = link["src"]
            print(f"Сохарняю картинку — {self.url + lnk}")

        img = urllib.request.urlopen(self.url + lnk).read()
        out = open(self.photoName, "wb")
        out.write(img)
        out.close()

    def get_size_photo(self):
        image = Image.open(self.photoName)
        width, height = image.size
        return width, height

    def get_size_screen(self):
        for m in get_monitors():
            width = m.width
            height = m.height
            return width, height

    def set_wallpaper(self):
        path = os.path.abspath(self.photoName)
        ctypes.windll.user32.SystemParametersInfoW(20, 0, path, 0)

    def aspect_ratio(self, width, height):
        k = math.gcd(width, height)
        return width // k, height // k

    def valid_image(self):
        w_image, h_image = self.get_size_photo()
        w_image, h_image = self.aspect_ratio(w_image, h_image)
        w_screen, h_screen = self.aspect_ratio(self.screen_size[0], self.screen_size[1])
        return w_image == w_screen and h_image == h_screen

    def crop_image(self):
        size_screen = self.get_size_screen()
        img = Image.open(self.photoName)
        img_crop = img.crop((0, 0, size_screen[0], size_screen[1]))
        img_crop.save(self.photoName)

    def start(self):
        self.download_photo()
        if not self.valid_image():
            self.crop_image()

        self.set_wallpaper()


if __name__ == "__main__":
    nasa = Nasa()
    nasa.start()
