package nosql

import (
	"godi/framework/dependency"
	"godi/framework/trace"
	"fmt"
)

func init() {
	dependency.FactoryRegister(Factory)
}

type nosql struct {
	m map[string]string
}

func Factory() interface{} {
	return &nosql{make(map[string]string)}
}

func (o *nosql) Get(key string) string {
	return o.m[key]
}
func (o *nosql) Set(key string, val string) error {
	o.m[key] = val
	return nil
}
func (o *nosql) Delete(key string)             {}
func (o *nosql) Update(key string, val string) {}
func (o *nosql) Name() string {
	return trace.MyName()
}

func (o *nosql) Init(){
	fmt.Println(trace.MyName(), "init")
}

func (o *nosql) Start(){
	fmt.Println(trace.MyName(), "start")
}