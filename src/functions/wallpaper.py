from bs4 import BeautifulSoup

import requests
import urllib
import sys
import ctypes
import os

from .config import Config


class Wallpaper(Config):
	def __init__(self):
		super().__init__()
		self.url = "https://apod.nasa.gov/apod/"
		self.config_path = os.path.join(os.path.dirname(os.path.abspath(__file__)), "config.ini")

		self.full_page = requests.get(self.url)
		self.soup = BeautifulSoup(self.full_page.content, 'html.parser')
		self.lnk = None

	def check(self):
		config = Config.get_setting(self, path=self.config_path, section="Config", setting="wallpaper-link")

		for link in self.soup.select("img"):
			self.lnk = link["src"]

		if self.lnk is None:
			return "On the NASA website, they posted not a picture but a video, " \
				"unfortunately, it will not work to install wallpaper"

		if config == self.lnk:
			return
		else:
			Config.update_setting(
				self,
				path=self.config_path,
				section="Config",
				setting="wallpaper-link",
				value=self.lnk
			)

	def download(self):
		try:
			img = urllib.request.urlopen(self.url + self.lnk).read()
			with open("everyNASA.jpg", "wb") as file:
				file.write(img)
		except requests.exceptions.ConnectionError:
			return "Connection error, please try again later."

	@staticmethod
	def set():
		if sys.platform == "win32":
			ctypes.windll.user32.SystemParametersInfoW(20, 0, os.path.abspath("everyNASA.jpg"), 0)
		elif sys.platform == "linux":
			desk_env = os.environ.get("DESKTOP_SESSION")
			if desk_env == "gnome":
				os.system(
					"gsettings set org.gnome.desktop.background "
					"picture-uri 'file://{}'"
					.format(
						os.path.abspath(
							"everyNASA.jpg"
						)
					)
				)
			elif desk_env == "plasma":
				import dbus
				jscript = """
					var allDesktops = desktops();
					print (allDesktops);
					for (i=0;i<allDesktops.length;i++) {
						d = allDesktops[i];
						d.wallpaperPlugin = "%s";
						d.currentConfigGroup = Array("Wallpaper", "%s", "General");
						d.writeConfig("Image", "file://%s")
				}
				"""
				bus = dbus.SessionBus()
				plasma = dbus.Interface(
					bus.get_object(
						'org.kde.plasmashell',
						'/PlasmaShell'),
					dbus_interface='org.kde.PlasmaShell'
				)

				plasma.evaluateScript(jscript % ("org.kde.image", "org.kde.image", os.path.abspath("everyNASA.jpg")))
			else:
				return "Your desktop environment is not supported."
		else:
			return "Your operating system is not supported."

		return "The wallpaper is installed."
