package controller

import (
	"fmt"
	"net/http"
	"qaz_latin/db"

	"gopkg.in/gin-gonic/gin.v1/json"
)

func GetAllUserScore(w http.ResponseWriter, r *http.Request) {
	Verify(w, r)

	ret := db.GetAllUser()

	fmt.Println(ret)
	json, _ := json.Marshal(&ret)

	w.WriteHeader(200)

	w.Write([]byte(json))
}
