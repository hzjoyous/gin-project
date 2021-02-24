package main

import (
	"thousand-hands-server/bootstrap"
)

func main() {

	app := bootstrap.Initialize()

	app.Run()
}
