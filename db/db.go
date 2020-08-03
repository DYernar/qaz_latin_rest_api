package db

import (
	"database/sql"
	"fmt"
	model "qaz_latin/models"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const host = "ec2-50-17-90-177.compute-1.amazonaws.com"
const port = "5432"
const user = "ogubyitzqmixyk"
const password = "d667be7dedcbc80818844c162250dfbb7cd629b00322f98b22e956ce710fdf7d"
const dbname = "denjk42qaepkv9"

const dbConf = "host=" + host + " port=" + port + " user=" + user + " password=" + password + " dbname=" + dbname + " sslmode=disable"

func Connect() (*sql.DB, error) {
	db, err := sql.Open("postgres", dbConf)
	if err != nil {
		fmt.Print("\n\n\n", err, "\n\n\n")
	}
	_, err = db.Exec("create table if not exists users(userid serial primary key, username varchar, email varchar, name varchar, score int, token varchar)")
	_, err = db.Exec("create table if not exists news(newsid serial primary key, title varchar, text varchar, img varchar)")
	_, err = db.Exec("create table if not exists select_word(userid int, score int)")

	return db, err
}

func Drop() {
	db, _ := Connect()
	db.Exec("DROP TABLE users")
	db.Exec("DROP TABLE select_word")
	db.Close()
}

func InsertUser(user model.User) bool {
	db, err := Connect()

	if err != nil {
		db.Close()
		return false
	}

	_, err = db.Query("insert into users (username, email, name, score, token) values ($1, $2, $3, $4, $5)", user.Username, user.Email, user.Name, 0, "")

	if err != nil {
		db.Close()
		return false
	}

	db.Close()
	return true
}

func UserExists(user model.User) bool {
	db, err := Connect()

	if err != nil {
		fmt.Println(err)
		defer db.Close()
		return true
	}

	row, err := db.Query("select username from users where username=$1 and email=$2", user.Username, user.Email)

	if err != nil {
		fmt.Println("error is : ", err)
		defer db.Close()
		return true
	}

	username := ""

	for row.Next() {
		row.Scan(&username)
		break
	}

	if username == "" {
		defer db.Close()
		return false
	}

	defer db.Close()
	return true
}

func UpdateScore(userid int, score int) {
	db, err := Connect()
	if err != nil {
		return
	}

	row, err := db.Query("select score from users where userid=$1", userid)

	if err != nil {
		fmt.Println(err)
		db.Close()
		return
	}

	defer row.Close()

	highest := 0

	for row.Next() {
		row.Scan(&highest)
		break
	}

	if highest < score {
		_, err = db.Query("update users set score=$1 where userid=$2 ", score, userid)
		if err != nil {
			fmt.Println(err)
			db.Close()
			return
		}
	} else {
		db.Close()
		return
	}
}

func InsertToken(user model.User, token string) {
	db, err := Connect()

	if err != nil {
		db.Close()
		return
	}

	_, err = db.Query("update users set token=$1 where email=$2 ", token, user.Email)

	if err != nil {
		fmt.Println(err)
		db.Close()
		return
	}

	defer db.Close()
}

func GetToken(user model.User) string {
	db, err := Connect()
	if err != nil {
		db.Close()
		return ""
	}

	row, err := db.Query("select token from users where email=$1", user.Email)

	if err != nil {
		fmt.Println(err)
		db.Close()
		return ""
	}

	defer row.Close()

	token := ""

	for row.Next() {
		row.Scan(&token)
		break
	}
	db.Close()
	return token
}

func GetUserFromToken(token string) model.User {
	var ret model.User
	db, err := Connect()
	if err != nil {
		db.Close()
		return ret
	}

	row, err := db.Query("select userid, username, email, name, score from users where token=$1", token)

	if err != nil {
		fmt.Println(err)
		db.Close()
		return ret
	}

	defer row.Close()

	var user model.User

	for row.Next() {
		row.Scan(&user.ID, &user.Username, &user.Email, &user.Name, &user.Score)
		break
	}
	db.Close()
	return user
}
