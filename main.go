package main

import "Go/NativeQuery/routes"

func main() {
	e := routes.Init()

	e.Logger.Fatal(e.Start(":1234"))
}
