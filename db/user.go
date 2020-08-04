package db

import model "qaz_latin/models"

func GetAllUser() []model.User {
	var ret []model.User
	db, err := Connect()

	if err != nil {
		db.Close()
		return ret
	}

	row, err := db.Query("select * from users")

	for row.Next() {
		var user model.User
		row.Scan(&user.ID, &user.Username, &user.Email, &user.Name, &user.Score, &user.Token)
		ret = append(ret, user)
	}

	db.Close()
	return ret

}

func GetUserById(usreid int) model.User {
	var ret model.User
	db, err := Connect()

	if err != nil {
		db.Close()
		return ret
	}

	row, err := db.Query("select * from users where userid=$1", usreid)

	for row.Next() {
		row.Scan(&ret.ID, &ret.Username, &ret.Email, &ret.Name, &ret.Score, &ret.Token)
		break
	}

	db.Close()
	return ret

}
