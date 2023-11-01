package db

import (
	"Go/NativeQuery/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

var err error

func Init() {

	conf := config.GetConfig()

	ConnectionString := conf.DB_USERNAME + " : " + conf.DB_PASSWORD + "@tcp(" + conf.DB_HOST + ":" + conf.DB_PORT + ")/" + conf.DB_NAME

	db, err = sql.Open("mysql", ConnectionString)

	if err != nil {
		panic("connectionString Error..")
	}

	// err = db.Ping()
	// if err != nil {
	// 	panic("DB Invalid")
	// }
}

func CreateCon() *sql.DB {
	return db
}
