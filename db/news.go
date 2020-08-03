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

func GetAllNews() []model.News {
	var ret []model.News
	db, err := Connect()

	if err != nil {
		fmt.Println(err)
		db.Close()
		return ret
	}

	row, err := db.Query("select title, text, img from news")

	if err != nil {
		fmt.Println(err)
		db.Close()
		return ret
	}

	defer row.Close()

	for row.Next() {
		var n model.News
		row.Scan(&n.Title, &n.Text, &n.Image)
		ret = append(ret, n)
	}

	db.Close()
	return ret
}
