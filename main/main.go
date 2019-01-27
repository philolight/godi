package main

import (
	"flag"
	"os"
	"time"

	"godi/framework/dependency"
	"godi/framework/trace"
	_ "godi/imports"

	"fmt"

	"github.com/labstack/gommon/log"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "c", "main/app_rdb.conf", "dependency configration file path")
	//flag.StringVar(&configPath, "c", "main/app_rdb.conf", "dependency configration file path")
}

func main() {
	flag.Parse()

	if err := dependency.Load(configPath); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	if err := dependency.Inject(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	trace.Dump()

	dependency.ObjectDiagram()

	fmt.Println(">>> Init")
	dependency.Call("Init")

	fmt.Println(">>> Start")
	dependency.Call("Start")

	time.Sleep(3 * time.Second)
}
