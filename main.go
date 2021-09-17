package main

import (
	"fmt"
	"github.com/sugiantodenny01/apibook/db"
	"github.com/sugiantodenny01/apibook/routes"
)

func main() {
	db.Initial()
	app := routes.Init()
	app.Listen(":12345")
	fmt.Println("yey running")
}
