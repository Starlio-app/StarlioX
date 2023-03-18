package controllers

import (
	"database/sql"
	"github.com/Redume/EveryNasa/utils"
	"github.com/gofiber/fiber/v2"
)

var GetFavorites = func(c *fiber.Ctx) error {
	FavoriteTitle := c.Query("title")

	type Favorite struct {
		Title       string `json:"title"`
		Explanation string `json:"explanation"`
		Copyright   string `json:"copyright"`
		Date        string `json:"date"`
		URL         string `json:"url"`
		HDURL       string `json:"hdurl"`
		MediaType   string `json:"media_type"`
	}

	var title, explanation, copyright, date, url, hdurl, media_type string
	db := utils.GetDatabase()

	if FavoriteTitle != "" {
		queryFavorites, err := db.Query("SELECT * FROM favorite WHERE title LIKE ?", FavoriteTitle)
		if err != nil {
			utils.Logger(err.Error())
		}

		defer func(query *sql.Rows) {
			err := query.Close()
			if err != nil {
				utils.Logger(err.Error())
			}
		}(queryFavorites)

		for queryFavorites.Next() {
			err := queryFavorites.Scan(&title, &explanation, &copyright, &date, &url, &hdurl, &media_type)
			if err != nil {
				utils.Logger(err.Error())
			}

			return c.JSON(fiber.Map{
				"title":       title,
				"explanation": explanation,
				"copyright":   copyright,
				"date":        date,
				"url":         url,
				"hdurl":       hdurl,
				"media_type":  media_type,
			})
		}
	} else {
		queryFavorite, err := db.Query("SELECT * FROM favorite")
		if err != nil {
			utils.Logger(err.Error())
		}

		defer func(query *sql.Rows) {
			err := query.Close()
			if err != nil {
				utils.Logger(err.Error())
			}
		}(queryFavorite)

		var favorites []Favorite
		for queryFavorite.Next() {
			err := queryFavorite.Scan(&title, &explanation, &copyright, &date, &url, &hdurl, &media_type)
			if err != nil {
				utils.Logger(err.Error())
			}

			favorites = append(favorites, Favorite{
				Title:       title,
				Explanation: explanation,
				Copyright:   copyright,
				Date:        date,
				URL:         url,
				HDURL:       hdurl,
				MediaType:   media_type})
		}

		return c.JSON(favorites)
	}

	return c.SendString("No favorites found")
}

var AddFavorite = func(c *fiber.Ctx) error {
	title := c.FormValue("title")
	explanation := c.FormValue("explanation")
	copyright := c.FormValue("copyright")
	date := c.FormValue("date")
	url := c.FormValue("url")
	hdurl := c.FormValue("hdurl")
	media_type := c.FormValue("media_type")

	if title == "" && explanation == "" && date == "" && url == "" && hdurl == "" && media_type == "" {
		utils.Respond(c, utils.Message(false, "All fields are required"))
		return nil
	}

	db := utils.GetDatabase()

	_, err := db.Exec("INSERT INTO favorite (title, explanation, copyright, date, url, hdurl, media_type) VALUES (?, ?, ?, ?, ?, ?, ?)",
		title,
		explanation,
		copyright,
		date,
		url,
		hdurl,
		media_type)

	if err != nil {
		utils.Logger(err.Error())
	}

	utils.Respond(c, utils.Message(true, "Favorite added"))
	return nil
}

var DeleteFavorite = func(c *fiber.Ctx) error {
	title := c.FormValue("title")

	db := utils.GetDatabase()

	_, err := db.Exec("DELETE FROM favorite WHERE title = ?", title)
	if err != nil {
		utils.Logger(err.Error())
	}

	utils.Respond(c, utils.Message(true, "Favorite deleted"))
	return nil
}
