package db

import (
	"database/sql"
)

var db *sql.DB

var err error

func Init() {
	conf := config.getConfig()

	connectionString := conf.DB_USERNAME + " : " + conf.DB_PASSWORD + " @ " + "tcp(" + conf.DB_HOST + ":" + conf.DB_PORT + ")/" + conf.DB_NAME

	db, err := sql.Open("mysql", connectionString)

	if err != nil {
		panic("connectionString Error")
	}

	err = db.Ping()
	if err != nil {
		panic("DB Invalid")
	}
}

func CreateCon() *sql.DB {
	return db
}
