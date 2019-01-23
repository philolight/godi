package bapp

import (
	"godi/framework/dependency"
	"godi/storage"
	"fmt"
)

type BApp struct {
	property string
	Storage  storage.Storage
}

func init() {
	dependency.FactoryWithName(Factory, "BApp")
}

func Factory() interface{} {
	return &BApp{}
}

func (a *BApp) Func() {
	// do something
}

func StorageName() string {
	obj := dependency.Get("BApp")
	b := obj.(*BApp)
	return b.Storage.Name()
}

func (a *BApp) Init(){
	fmt.Println("BApp Init")
}

func (a *BApp) Start() {
	fmt.Println("BApp Start")
}