package main

import (
	"github.com/dudyali/echo-rest/db"
	"github.com/dudyali/echo-rest/routes"
)

func main() {
	db.Init()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":1234"))
}
