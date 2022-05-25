package main

import (
	"Tortuga/internal"
	"fmt"
)

func main() {
	app, err := internal.BuildApp()

	if err != nil {
		fmt.Printf(err.Error())
	}

	err = app.StartHttp()

	if err != nil {
		fmt.Printf(err.Error())
	}
}
