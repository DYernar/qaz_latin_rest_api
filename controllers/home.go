package controller

import (
	"encoding/json"
	"net/http"
	"qaz_latin/db"
)

func GetData(w http.ResponseWriter, r *http.Request) {
	Verify(w,r)

	token := r.Header.Get("token")

	if token == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	user := db.GetUserFromToken(token)

	if user.Username == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json, _ := json.Marshal(&user)
	w.WriteHeader(200)
	w.Write([]byte(string(json)))
}
