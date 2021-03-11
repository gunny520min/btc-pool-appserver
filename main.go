package main

import (
	"btc-pool-appserver/application/app"
	"btc-pool-appserver/application/library/log"
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("NumCPU: %d, GOMAXPROCS: %d\n", runtime.NumCPU(), runtime.GOMAXPROCS(-1))
	//initial app
	app.Init("web", "./")
	defer app.Exit()

	if err := app.Start(); err != nil {
		log.Error(err)
		panic(err.Error())
	}
}
