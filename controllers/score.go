package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"qaz_latin/db"
	"strconv"
)

func UpdateScore(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("appToken") != AppToken {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tokenString := r.Header.Get("token")

	if tokenString == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user := db.GetUserFromToken(tokenString)

	if user.Username == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	r.ParseForm()

	newScore, err := strconv.Atoi(r.FormValue("score"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	db.UpdateScore(user.ID, newScore)
	db.AddScore(user.ID, newScore)

	user = db.GetUserFromToken(tokenString)
	json, _ := json.Marshal(&user)
	w.WriteHeader(200)
	w.Write([]byte(json))
}

func GetScores(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("appToken") != AppToken {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tokenString := r.Header.Get("token")

	if tokenString == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user := db.GetUserFromToken(tokenString)

	scores := db.GetScores(user.ID)

	js, _ := json.Marshal(scores)

	w.WriteHeader(200)

	w.Write(js)
}

func SaveScore(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("appToken") != AppToken {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tokenString := r.Header.Get("token")

	if tokenString == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	r.ParseForm()

	score, err := strconv.Atoi(r.FormValue("score"))
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	gameid, err := strconv.Atoi(r.FormValue("game"))

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user := db.GetUserFromToken(tokenString)

	db.InsertScore(user.ID, gameid, score)

	w.WriteHeader(200)
}

func GetScore(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("appToken") != AppToken {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tokenString := r.Header.Get("token")

	if tokenString == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	r.ParseForm()

	gameid, err := strconv.Atoi(r.FormValue("gameid"))

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user := db.GetUserFromToken(tokenString)

	scores := db.GetScore(user.ID, gameid)

	js, _ := json.Marshal(scores)

	w.WriteHeader(200)

	w.Write(js)
}
