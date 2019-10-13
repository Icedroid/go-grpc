package main

import (
	"flag"

)

var appName = flag.String("appName", "go-grpc", "set app name.")

func main() {
	flag.Parse()

	app, err := CreateApp(*appName)
	if err != nil {
		panic(err)
	}

	if err := app.Start(); err != nil {
		panic(err)
	}

	app.AwaitSignal()
}
