package functions

import (
	"encoding/json"
	"os"
	"time"

	"github.com/jasonlvhit/gocron"
	"github.com/reujab/wallpaper"
	"github.com/rodkranz/fetch"
	"gopkg.in/yaml.v2"
)

func GetWallpaper() string {
	file, err := os.ReadFile("config.yaml")
	if err != nil {
		Logger(err.Error())
	}
	data := make(map[interface{}]interface{})

	err = yaml.Unmarshal(file, &data)
	if err != nil {
		Logger(err.Error())
	}

	type WallpaperStruct struct {
		Hdurl      string `json:"hdurl"`
		Url        string `json:"url"`
		Media_type string `json:"media_type"`
	}

	client := fetch.NewDefault()
	res, err := client.Get("https://api.nasa.gov/planetary/apod?api_key="+data["apiKey"].(string), nil)
	if err != nil {
		Logger(err.Error())
	}

	body, err := res.ToString()
	if err != nil {
		Logger(err.Error())
	}

	var Wallpaper WallpaperStruct
	err = json.Unmarshal([]byte(body), &Wallpaper)
	if err != nil {
		Logger(err.Error())
	}

	if Wallpaper.Media_type == "video" {
		Wallpaper.Url = Wallpaper.Url[30 : len(Wallpaper.Url)-6]
		return "https://img.youtube.com/vi/" + Wallpaper.Url + "/0.jpg"
	}

	return Wallpaper.Hdurl
}

func SetWallpaper() {
	err := wallpaper.SetFromURL(GetWallpaper())
	if err != nil {
		Logger(err.Error())
	}
}

func StartWallpaper() {
	type Autostart struct {
		Wallpaper int `json:"wallpaper"`
	}

	client := fetch.NewDefault()
	res, err := client.Get("http://localhost:8080/api/get/settings", nil)
	if err != nil {
		Logger(err.Error())
	}

	body, err := res.ToString()
	if err != nil {
		Logger(err.Error())
	}

	var AutostartSetWallpaper Autostart
	err = json.Unmarshal([]byte(body), &AutostartSetWallpaper)
	if err != nil {
		Logger(err.Error())
	}

	if AutostartSetWallpaper.Wallpaper == 1 {
		times := time.Now()
		t := time.Date(times.Year(), times.Month(), times.Day(), 4, 50, times.Second(), times.Nanosecond(), time.UTC)

		SetWallpaper()

		err = gocron.Every(1).Day().From(&t).Do(SetWallpaper)
		if err != nil {
			Logger(err.Error())
		}

		gocron.Start()
	}
}
