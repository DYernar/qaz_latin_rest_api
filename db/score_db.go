package db

import (
	"fmt"
	model "qaz_latin/models"
	"sort"
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

	row, err := db.Query("select score from game_results where gameid=$1 and userid=$2", gameid, userid)

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

func GetUsersRank(gameid int) []model.User {
	var ret []model.User
	db, err := Connect()
	if err != nil {
		fmt.Println(err)
		return ret
	}

	row, err := db.Query("select userid, score from game_results where gameid=$1", gameid)

	if err != nil {
		fmt.Println(err)
		return ret
	}

	defer row.Close()

	m := make(map[int]int)

	for row.Next() {
		var userid int
		var score int
		if m[userid] < score {
			m[userid] = score
		}
	}

	for userid, s := range m {
		user := GetUserById(userid)
		user.Score = s
		ret = append(ret, user)
	}

	sort.Slice(ret, func(i, j int) bool {
		return ret[i].Score < ret[j].Score
	})
	return ret
}
