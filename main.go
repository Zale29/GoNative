package main

import (
	"Go/NativeQuery/db"
	"Go/NativeQuery/routes"
)

func main() {
	db.Init()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":1234"))
}
