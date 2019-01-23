package app

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"godi/bapp"
	_ "godi/bapp"
	"godi/framework/dependency"
	"godi/framework/trace"
	_ "godi/storage/nosql"
	_ "godi/storage/rdb"

	"github.com/stretchr/testify/assert"
)

func TestApp(t *testing.T) {
	dependency.Set("app:app", "Storage", "storage.rdb:rdb")
	dependency.Set("app:app", "Name", "Application")
	dependency.Set("app:app", "Value", "987654321")
	dependency.Set("app:app", "Dur", "10s")
	dependency.Set("app:app", "T", "2006-02-03T15:04:05Z")
	dependency.Set("BApp", "Storage", "storage.rdb:rdb")
	dependency.Set("storage.rdb:rdb", "Connections", "100")

	err := dependency.Create()
	assert.Nil(t, err)
	err = dependency.Inject()
	assert.Nil(t, err)

	trace.Dump()

	assert.Equal(t, "storage.rdb:rdb", StorageName())

	a := dependency.Get("app:app").(*App)
	assert.Equal(t, "Application", a.Name)
	assert.Equal(t, 987654321, a.Value)
	assert.Equal(t, time.Second*10, a.Dur)

	ti, err := time.Parse("2006-01-02T15:04:05Z", "2006-02-03T15:04:05Z")
	assert.Nil(t, err)
	assert.Equal(t, time.Duration(0), ti.Sub(a.T))

	assert.Equal(t, "storage.rdb:rdb", bapp.StorageName())
}

func TestAppType(t *testing.T) {
	sut := App{}

	rsut := reflect.ValueOf(sut)

	fmt.Println("rsut = ", rsut)

	fmt.Println(rsut.Type())
	fmt.Println("rsut.Type().Kind() = ", rsut.Type().Kind())
}

func TestTime(t *testing.T) {
	ti := time.Now()

	rti := reflect.ValueOf(ti)

	fmt.Println("rti = ", rti)

	fmt.Println("rsut.Type() = ", rti.Type())
	fmt.Println("rsut.Type().Kind() = ", rti.Type().Kind())
}

func TestDuration(t *testing.T) {
	str := "10s"

	d, err := time.ParseDuration(str)
	assert.Nil(t, err)
	assert.Equal(t, time.Second*10, d)

	dur := time.Duration(time.Second * 10)

	rdur := reflect.ValueOf(dur)

	fmt.Println("rdur = ", rdur)

	fmt.Println("type = ", rdur.Type())
	fmt.Println("rsut.Type().Kind() = ", rdur.Type().Kind())
}

func TestLoad(t *testing.T) {
	err := dependency.Load("app_test.conf")
	assert.Nil(t, err)

	err = dependency.Create()
	assert.Nil(t, err)
	err = dependency.Inject()
	assert.Nil(t, err)

	trace.Dump()

	assert.Equal(t, "storage.rdb:rdb", StorageName())

	a := dependency.Get("app:app").(*App)
	assert.Equal(t, "Application", a.Name)
	assert.Equal(t, 987654321, a.Value)
	assert.Equal(t, time.Second*10, a.Dur)

	ti, err := time.Parse("2006-01-02T15:04:05Z", "2006-02-03T15:04:05Z")
	assert.Nil(t, err)
	assert.Equal(t, time.Duration(0), ti.Sub(a.T))

	assert.Equal(t, "storage.rdb:rdb", bapp.StorageName())
}