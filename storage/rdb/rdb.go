package rdb

import (
	"gofra/framework/dependency"
	"gofra/framework/trace"
)

func init() {
	dependency.FactoryRegister(Factory)
}

type rdb struct {
	m           map[string]string
	Connections int
}

func Factory() interface{} {
	return &rdb{m: make(map[string]string)}
}

func (o *rdb) Get(key string) string {
	return o.m[key]
}
func (o *rdb) Set(key string, val string) error {
	o.m[key] = val
	return nil
}
func (o *rdb) Delete(key string)             {}
func (o *rdb) Update(key string, val string) {}
func (o *rdb) Name() string {
	return trace.MyName()
}
