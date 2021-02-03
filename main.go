package main

import (
	"goe/app"
)

func main() {
	appServer := &app.App{Host: "0.0.0.0", Port: "8989"}
	appServer.Start()
}
