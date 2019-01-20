package app

import (
	"gofra/framework/dependency"
	"gofra/framework/trace"
	"gofra/storage"
	"time"
)

// 어플리케이션 구조체의 선언
type app struct {
	property string
	Storage  storage.Storage
	Name     string
	Value    int
	Dur      time.Duration
	T        time.Time
}

// init() 함수에서 콜백을 등록한다.
func init() {
	dependency.FactoryRegister(Factory)
}

func Factory() interface{} {
	return &app{}
}

// 4. 기능
func (a *app) Func() {
	go a.RepeatPrint()
}

func (a *app) RepeatPrint(){
	for {
		a.Storage.Name()
	}
}

func StorageName() string {
	obj := dependency.Get(trace.MyName())
	a := obj.(*app)
	return a.Storage.Name()
}
