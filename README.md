<p align="center"><img src="web/static/image/icons/banner.png" alt="EveryNasa banner" title="EveryNasa"></p>

The program takes a picture from the NASA [website](https://apod.nasa.gov/apod) every day
and sets it as a background for your workspace.

---

### Contents
- Solution or answers to possible problems [Windows](#windows) / [Linux](#linux)
- [Build project](#build-project)

---

| OS      | Status      | Latest version | Download                                                                            |
|---------|-------------|----------------|-------------------------------------------------------------------------------------|
| Windows | Available   | 1.6            | [.exe](https://github.com/Redume/EveryNasa/releases/download/v1.6/EveryNasa.exe)    |
| Debian  | Available   | 1.6            | [Binary file](https://github.com/Redume/EveryNasa/releases/download/v1.6/EveryNasa) |
| Android | Available   | 1.2.0          | [Google Play](https://play.google.com/store/apps/details?id=ru.murzify.everynasa)   |
| MacOS   | Unavailable |                |                                                                                     |
| iOS     | Unavailable |                |                                                                                     |

---

### Solution or answers to possible problems
#### Windows
<details>
<summary></summary>
    <li>To make all functions work correctly, install the program anywhere except Program Files(x86) / Program Files</li>
</details>

#### Linux
<details>
<summary></summary>

- If you have a mistake with `ayatana-appindicator3-0.1`

    <details>
        <summary><b>Debian / Ubuntu / Mint</b></summary>
        <details>
            <summary><b>KDE Plasma</b></summary>

  ```shell
  $ sudo apt install gir1.2-appindicator3-0.1
  ```

  </details>
  <details>
  <summary><b>GNOME</b></summary>

  - Install the package
  ```shell
  $ sudo apt install gnome-shell-extension-appindicator
  ```

  - Open `Tweaks`
  - Go to `Extensions`
  - Enable `Kstatusnotifieritem/appindicator support`
</details>
</details>
</details>

### Build project

<details>
<summary></summary>

- Close the repository
```shell
$ git clone https://github.com/Redume/EveryNasa.git
```
- Change directory
```shell
$ cd EveryNasa
```
- Build project
<details>
<summary><b>Windows</b></summary>

```shell
$ go build -o EveryNasa.exe -ldflags = "-H windowsgui"
```

</details>
<details>
<summary><b>Linux</b></summary>

```shell
$ go build -o EveryNasa
```
</details>
</details>