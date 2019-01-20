package app

import (
	"gofra/framework/dependency"
	"gofra/framework/trace"
	"gofra/storage"
	"time"
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

func (a *App) Start() {
	go a.RepeatPrint()
}

func (a *App) RepeatPrint(){
	for {
		a.Storage.Name()
	}
}

func StorageName() string {
	obj := dependency.Get(trace.MyName())
	a := obj.(*App)
	return a.Storage.Name()
}