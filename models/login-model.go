package models

import (
	"Go/NativeQuery/db"
	"Go/NativeQuery/helper"
	"database/sql"
	"fmt"
)

type User struct {
	Username string `json:"username"`
}

func Checklogin(Username, password string) (bool, error) {
	var obj User
	var pwd string

	con := db.CreateCon()

	sqlStatement := "SELECT * FROM useraccess WHERE username = ?"

	err := con.QueryRow(sqlStatement, Username).Scan(&obj.Username, &pwd)

	if err == sql.ErrNoRows {
		fmt.Println("Data Tidak Tersedia!")
		return false, err
	}

	if err != nil {
		fmt.Println("Error Login")
		return false, err
	}

	match, err := helper.CheckGeneratePassword(password, pwd)
	if !match {
		fmt.Println("Password Yang Anda Masukan Salah!!!")
		return false, err
	}

	return true, nil

}
