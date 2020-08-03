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
		return ret
	}

	row, err := db.Query("select score from select_word")

	if err != nil {
		fmt.Println(err)
		return ret
	}

	defer row.Close()

	for row.Next() {
		var n int
		row.Scan(&n)
		ret = append(ret, n)
	}

	return ret
}
