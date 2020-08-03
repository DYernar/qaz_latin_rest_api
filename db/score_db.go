package db

import (
	"fmt"
)

func AddScore(userid int, score int) {
	db, err := Connect()

	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = db.Query("insert into select_word (userid, score) values ($1, $2)", userid, score)

	if err != nil {
		fmt.Println(err)
	}
	db.Close()
}

func GetScores(userid int) []int {
	var ret []int
	db, err := Connect()

	if err != nil {
		fmt.Println(err)
		db.Close()
		return ret
	}

	row, err := db.Query("select score from select_word where userid=$1", userid)

	if err != nil {
		fmt.Println(err)
		db.Close()
		return ret
	}

	defer row.Close()

	for row.Next() {
		var n int
		row.Scan(&n)
		ret = append(ret, n)
	}

	db.Close()
	return ret
}

func InsertScore(userid, gameid, score int) {
	db, err := Connect()

	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = db.Query("insert into game_results (gameid, userid, score) values ($1, $2, $3)", gameid, userid, score)

	if err != nil {
		fmt.Println(err)
	}
	db.Close()
}

func GetScore(userid int, gameid int) []int {
	var ret []int
	db, err := Connect()

	if err != nil {
		fmt.Println(err)
		db.Close()
		return ret
	}

	row, err := db.Query("select score from game_results where gameid=$1 userid=$2", gameid, userid)

	if err != nil {
		fmt.Println(err)
		db.Close()
		return ret
	}

	defer row.Close()

	for row.Next() {
		var n int
		row.Scan(&n)
		ret = append(ret, n)
	}

	db.Close()
	return ret
}
