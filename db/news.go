package db

import (
	"fmt"
	model "qaz_latin/models"
)

func InsertNews(news model.News) {
	db, err := Connect()

	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = db.Query("insert into news (title, text, img) values ($1, $2, $3)", news.Title, news.Text, news.Image)

	if err != nil {
		fmt.Println(err)
	}
	db.Close()
}
