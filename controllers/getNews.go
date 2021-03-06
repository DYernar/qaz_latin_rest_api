package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"qaz_latin/db"
	model "qaz_latin/models"
)

func GetNews(w http.ResponseWriter, r *http.Request) {
	Verify(w, r)

	news := db.GetAllNews()

	json, _ := json.Marshal(&news)

	w.WriteHeader(200)
	w.Write([]byte(json))
}

func AddNews(w http.ResponseWriter, r *http.Request) {
	Verify(w, r)

	var news model.News

	err := json.NewDecoder(r.Body).Decode(&news)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	db.InsertNews(news)

	w.WriteHeader(200)
}
