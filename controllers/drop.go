package controller

import (
	"net/http"
	"qaz_latin/db"
)

func DropTables(w http.ResponseWriter, r *http.Request) {
	db.Drop()
	w.WriteHeader(200)
}
