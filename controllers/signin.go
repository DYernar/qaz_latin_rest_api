package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"qaz_latin/db"
	model "qaz_latin/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("my_secret_key")
var AppToken = "Qazaq latin app"

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func Signin(w http.ResponseWriter, r *http.Request) {
	Verify(w, r)
	var userdata model.User

	err := json.NewDecoder(r.Body).Decode(&userdata)
	if err != nil {
		fmt.Println(r.Body)
		fmt.Println("error: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !db.UserExists(userdata) {
		db.InsertUser(userdata)
	}

	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &Claims{
		Username: userdata.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	db.InsertToken(userdata, tokenString)

	res := model.SigninResponse{Status: 200, Token: tokenString, User: userdata}
	json, err := json.Marshal(res)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
	w.Write([]byte(string(json)))
}
