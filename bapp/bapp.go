package bapp

import (
	"gofra/framework/dependency"
	"gofra/framework/trace"
	"gofra/storage"
)

// 어플리케이션 구조체의 선언
type BApp struct {
	property string
	Storage  storage.Storage
}

// init() 함수에서 콜백을 등록한다.
func init() {
	dependency.FactoryRegister(Factory)
}

func Factory() interface{} {
	return &BApp{}
}

// 4. 기능
func (a *BApp) Func() {
	// do something
}

func StorageName() string {
	obj := dependency.Get(trace.MyName())
	b := obj.(*BApp)
	return b.Storage.Name()
}
