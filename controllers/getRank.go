package controller

import (
	"net/http"
	"qaz_latin/db"
	"strconv"

	"gopkg.in/gin-gonic/gin.v1/json"
)

func GetRankById(w http.ResponseWriter, r *http.Request) {
	Verify(w, r)

	tokenString := r.Header.Get("token")

	if tokenString == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	r.ParseForm()

	gameid, err := strconv.Atoi(r.FormValue(r.FormValue("gameid")))

	if err != nil {
		print(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	usersRankDescending := db.GetUsersRank(gameid)

	json, _ := json.Marshal(&usersRankDescending)

	w.WriteHeader(200)

	w.Write([]byte(json))
}