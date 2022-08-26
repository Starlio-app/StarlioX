<p align="center"><img src="./src/icons/EveryNASA-banner.png" alt="EveryNasa banner"></p>

## Solution or answers to possible problems
- [Windows](#windows)
- [Debian](#linux)

---
## **Feature**
- Rewrite, if possible, into the GoLang programming language
- Make gui
- Make autorun and other settings

The program takes a picture from the NASA [website](https://apod.nasa.gov/apod) every day 
and sets it as a background for your workspace. 
Unfortunately, the addition of autorun is missing due to technical errors.

| OS      	| Status      	| Download link                                                                        	|
|---------	|-------------	|--------------------------------------------------------------------------------------	|
| Windows 	| Available   	| [Download](https://github.com/Redume/EveryNasa/releases/download/v1.6/EveryNasa.exe) 	|
| Debian  	| Available   	| [Download](https://github.com/Redume/EveryNasa/releases/download/v1.6/EveryNasa)     	|
| Andorid 	| Dev         	| [Google Play](https://play.google.com/store/apps/details?id=ru.murzify.everynasa)    	|
| MacOS   	| Unavailable 	|                                                                                      	|
| iOS     	| Unavailable 	|                                                                                      	|
## Windows

[❗] If you get a notification that `EveryNasa` is a virus, 
then disable the antivirus because it mistakenly believes, 
and also deletes the file, in the near future I'm trying to solve this problem

---

## Linux

### [❗] Only the Debian distribution was tested

#### Gnome

1. Install `gnome-shell-extension-appindicator`
```shell
$ sudo apt install gnome-shell-extension-appindicator
```
2. Search `tweaks` in your `Activities` screen
3. Switch `Kstatusnotifieritem/appindicator support` on

---

#### Kde Plasma

Install `gir1.2-appindicator3-0.1`

```shell
$ sudo apt install gir1.2-appindicator3-0.1
```
