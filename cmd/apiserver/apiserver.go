package main

import (
	"spreadTask/cmd/apiserver/app"
)

func main() {
	server := app.NewApiServer()

	server.Run()
}
