package functions

import (
	"encoding/json"
	"time"

	"github.com/jasonlvhit/gocron"
	"github.com/reujab/wallpaper"
	"github.com/rodkranz/fetch"
)

func GetWallpaper(date string) string {
	type WallpaperStruct struct {
		Hdurl      string `json:"hdurl"`
		Url        string `json:"url"`
		Media_type string `json:"media_type"`
	}

	client := fetch.NewDefault()
	res, err := client.Get("https://api.nasa.gov/planetary/apod?api_key=1gI9G84ZafKDEnrbydviGknReOGiVK9jqrQBE3et&date="+
		date,
		nil)
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
		Wallpaper.Url = "https://img.youtube.com/vi/" + Wallpaper.Url[30:41] + "/maxresdefault.jpg"
		return "https://img.youtube.com/vi/" + Wallpaper.Url[30:41] + "/maxresdefault.jpg"
	}

	return Wallpaper.Hdurl
}

func SetWallpaper() {
	times := time.Now()
	t := time.Date(times.Year(),
		times.Month(),
		times.Day(),
		times.Hour(),
		times.Minute(),
		times.Second(),
		times.Nanosecond(),
		time.UTC)

	err := wallpaper.SetFromURL(GetWallpaper(t.Format("2006-01-02")))
	if err != nil {
		t = t.AddDate(0, 0, -1)
		err = wallpaper.SetFromURL(GetWallpaper(t.Format("2006-01-02")))
		if err != nil {
			Logger(err.Error())
		}
	}
}

func StartWallpaper() {
	type Autostart struct {
		Wallpaper int `json:"wallpaper"`
	}

	client := fetch.NewDefault()
	res, err := client.Get("http://localhost:3000/api/get/settings", nil)
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
		t := time.Date(times.Year(),
			times.Month(),
			times.Day(),
			4,
			50,
			times.Second(),
			times.Nanosecond(),
			time.UTC)

		SetWallpaper()

		err = gocron.Every(1).Day().From(&t).Do(SetWallpaper)
		if err != nil {
			Logger(err.Error())
		}

		gocron.Start()
	}
}
