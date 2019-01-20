package main

import (
	"os"
	"flag"
	"time"

	"godi/app"
	"godi/framework/dependency"
	"godi/framework/trace"
	_ "godi/imports"

	"github.com/labstack/gommon/log"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "c", "main/app_rdb.conf", "dependency configration file path")
}

func main() {
	flag.Parse()

	if err := dependency.Load(configPath); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	if err := dependency.Create(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	if err := dependency.Inject(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	trace.Dump()

	appInstance := dependency.Get("app:app").(*app.App)
	appInstance.Start()

	time.Sleep(10 * time.Second)
}