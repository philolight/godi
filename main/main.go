package main

import (
	"fmt"
	"gofra/app"
	"gofra/framework/dependency"
	"gofra/framework/property"
	"gofra/framework/trace"
	_ "gofra/imports"
	"os"

	"github.com/labstack/gommon/log"
)

func main() {
	if err := dependency.Load("app.conf"); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	dependency.Create()
	dependency.Inject()

	if err := property.Load("filename"); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	if err := property.Init(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	trace.Dump()

	fmt.Println("app.Storage = ", app.StorageName())
}
