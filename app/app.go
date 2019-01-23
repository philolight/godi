package app

import (
	"godi/framework/dependency"
	"godi/framework/trace"
	"godi/storage"
	"time"
	"fmt"
)

// Declaration of Application
type App struct {
	property string
	Storage  storage.Storage
	Name     string
	Value    int
	Dur      time.Duration
	T        time.Time
}

// Factory method register in init()
func init() {
	dependency.FactoryRegister(Factory)
}

// Factory method of app
func Factory() interface{} {
	return &App{}
}

func (a *App) Init(){
	fmt.Println(trace.MyName(), "Init")
	fmt.Printf("%+v\n", a)
}

func (a *App) Start() {
	fmt.Println(trace.MyName(), "Start")
	go a.RepeatPrint()
}

func (a *App) RepeatPrint(){
	for {
		fmt.Println("my storage = ", a.Storage.Name())
		time.Sleep(500 * time.Millisecond)
	}
}

func StorageName() string {
	obj := dependency.Get(trace.MyName())
	a := obj.(*App)
	return a.Storage.Name()
}