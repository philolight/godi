package dependency

import (
	"encoding/json"
	"fmt"
	"math"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type test struct {
	B    bool
	I    int
	I8   int8
	I16  int16
	I32  int32
	I64  int64
	Ui   uint
	Ui8  uint8
	Ui16 uint16
	Ui32 uint32
	Ui64 uint64
	F32  float32
	F64  float64
	S    string
	Is   []int
	Im   map[int]string
	Imc  map[int][]string
	Imc2 map[string][]int
}

func TestTypeParser(t *testing.T) {
	sut := NewParser()

	to := test{}

	ro := test{
		B:    true,
		I:    math.MinInt32,
		I8:   math.MinInt8,
		I16:  math.MinInt16,
		I32:  math.MinInt32,
		I64:  math.MinInt64,
		Ui:   math.MaxUint32,
		Ui8:  math.MaxUint8,
		Ui16: math.MaxUint16,
		Ui32: math.MaxUint32,
		Ui64: math.MaxUint64,
		F32:  math.MaxFloat32,
		F64:  math.MaxFloat64,
		S:    "str",
		Is:   []int{1, 2, 3, 4},
		Im: map[int]string{
			1: "{[,1",
			2: "{[,2",
			3: "3,}]",
			4: "4,}]",
		},
		Imc: map[int][]string{
			1: {"[a]", "{b}"},
			2: {"d,", "e"},
			3: {"g", "h"},
		},
		Imc2: map[string][]int{
			"a":   {1, 2},
			"[b]": {1, 3},
			",c":  {5, 6},
		},
	}

	b, err := json.Marshal(ro)
	fmt.Println(string(b))

	var (
		sb    = "true"
		si    = "-2147483648"
		si8   = "-128"
		si16  = "-32768"
		si32  = "-2147483648"
		si64  = "-9223372036854775808"
		sui   = "4294967295"
		sui8  = "255"
		sui16 = "65535"
		sui32 = "4294967295"
		sui64 = "18446744073709551615"
		sf32  = "3.4028235e+38"
		sf64  = "1.7976931348623157e+308"
		s     = "str"
		sis   = "[1, 2, 3, 4]"
		sim   = `{1:"{[,1", 2:"{[,2", 3:"3,}]", 4:"4,}]"}`
		simc  = `{1:["[a]","{b}"],2:["d,","e"],3:["g","h"]}`
		simc2 = `{",c":[5,6],"[b]":[1,3],"a":[1,2]}`
	)

	rt := reflect.ValueOf(&to).Elem()

	f := rt.FieldByName("B")
	v, err := sut.Value(f.Type(), sb)
	assert.Nil(t, err)
	f.Set(*v)
	assert.Equal(t, to.B, ro.B)

	f = rt.FieldByName("I")
	v, err = sut.Value(f.Type(), si)
	assert.Nil(t, err)
	f.Set(*v)
	assert.Equal(t, to.I, ro.I)

	f = rt.FieldByName("I8")
	v, err = sut.Value(f.Type(), si8)
	assert.Nil(t, err)
	f.Set(*v)
	assert.Equal(t, to.I8, ro.I8)

	f = rt.FieldByName("I16")
	v, err = sut.Value(f.Type(), si16)
	assert.Nil(t, err)
	f.Set(*v)
	assert.Equal(t, to.I16, ro.I16)

	f = rt.FieldByName("I32")
	v, err = sut.Value(f.Type(), si32)
	assert.Nil(t, err)
	f.Set(*v)
	assert.Equal(t, to.I32, ro.I32)

	f = rt.FieldByName("I64")
	v, err = sut.Value(f.Type(), si64)
	assert.Nil(t, err)
	f.Set(*v)
	assert.Equal(t, to.I64, ro.I64)

	f = rt.FieldByName("Ui")
	v, err = sut.Value(f.Type(), sui)
	assert.Nil(t, err)
	f.Set(*v)
	assert.Equal(t, to.Ui, ro.Ui)

	f = rt.FieldByName("Ui8")
	v, err = sut.Value(f.Type(), sui8)
	assert.Nil(t, err)
	f.Set(*v)
	assert.Equal(t, to.Ui8, ro.Ui8)

	f = rt.FieldByName("Ui16")
	v, err = sut.Value(f.Type(), sui16)
	assert.Nil(t, err)
	f.Set(*v)
	assert.Equal(t, to.Ui16, ro.Ui16)

	f = rt.FieldByName("Ui32")
	v, err = sut.Value(f.Type(), sui32)
	assert.Nil(t, err)
	f.Set(*v)
	assert.Equal(t, to.Ui32, ro.Ui32)

	f = rt.FieldByName("Ui64")
	v, err = sut.Value(f.Type(), sui64)
	assert.Nil(t, err)
	f.Set(*v)
	assert.Equal(t, to.Ui64, ro.Ui64)

	f = rt.FieldByName("F32")
	v, err = sut.Value(f.Type(), sf32)
	assert.Nil(t, err)
	f.Set(*v)
	assert.Equal(t, to.F32, ro.F32)

	f = rt.FieldByName("F64")
	v, err = sut.Value(f.Type(), sf64)
	assert.Nil(t, err)
	f.Set(*v)
	assert.Equal(t, to.F64, ro.F64)

	f = rt.FieldByName("S")
	v, err = sut.Value(f.Type(), s)
	assert.Nil(t, err)
	f.Set(*v)
	assert.Equal(t, to.S, ro.S)

	f = rt.FieldByName("Is")
	v, err = sut.Value(f.Type(), sis)
	assert.Nil(t, err)
	f.Set(*v)
	assert.True(t, reflect.DeepEqual(to.Is, ro.Is), fmt.Sprintf("%+v", to.Is))

	f = rt.FieldByName("Im")
	v, err = sut.Value(f.Type(), sim)
	assert.Nil(t, err)
	f.Set(*v)
	assert.True(t, reflect.DeepEqual(to.Im, ro.Im))

	f = rt.FieldByName("Imc")
	v, err = sut.Value(f.Type(), simc)
	assert.Nil(t, err)
	f.Set(*v)
	assert.True(t, reflect.DeepEqual(to.Imc, ro.Imc))

	f = rt.FieldByName("Imc2")
	v, err = sut.Value(f.Type(), simc2)
	assert.Nil(t, err)
	f.Set(*v)
	assert.True(t, reflect.DeepEqual(to.Imc2, ro.Imc2))
}
