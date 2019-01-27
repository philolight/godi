package dependency

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

var Value Parser

func init() {
	parser := MultiParser{
		m: map[reflect.Kind]Parser{
			reflect.Ptr:     newPtrParser().Value,
			reflect.Bool:    boolValue,
			reflect.Int:     intValue,
			reflect.Int8:    int8Value,
			reflect.Int16:   int16Value,
			reflect.Int32:   int32Value,
			reflect.Int64:   int64Value,
			reflect.Uint:    uintValue,
			reflect.Uint8:   uint8Value,
			reflect.Uint16:  uint16Value,
			reflect.Uint32:  uint32Value,
			reflect.Uint64:  uint64Value,
			reflect.Float32: float32Value,
			reflect.Float64: float64Value,
			reflect.Map:     mapValue,
			reflect.Slice:   sliceValue,
			reflect.String:  stringValue,
		},
	}
	Value = parser.Value
}

type Parser func(ftype reflect.Type, value string) (*reflect.Value, error)

type MultiParser struct {
	m map[reflect.Kind]Parser
}

func NewParser() *MultiParser {
	return &MultiParser{
		m: map[reflect.Kind]Parser{
			reflect.Ptr:     newPtrParser().Value,
			reflect.Bool:    boolValue,
			reflect.Int:     intValue,
			reflect.Int8:    int8Value,
			reflect.Int16:   int16Value,
			reflect.Int32:   int32Value,
			reflect.Int64:   int64Value,
			reflect.Uint:    uintValue,
			reflect.Uint8:   uint8Value,
			reflect.Uint16:  uint16Value,
			reflect.Uint32:  uint32Value,
			reflect.Uint64:  uint64Value,
			reflect.Float32: float32Value,
			reflect.Float64: float64Value,
			reflect.Map:     mapValue,
			reflect.Slice:   sliceValue,
			reflect.String:  stringValue,
		},
	}
}

func (m *MultiParser) Value(ftype reflect.Type, value string) (*reflect.Value, error) {
	Value, ok := m.m[ftype.Kind()]
	if !ok {
		return nil, fmt.Errorf("failed to set value")
	}

	return Value(ftype, value)
}

type PtrParser struct {
	m map[reflect.Kind]Parser
}

func (p *PtrParser) Value(ftype reflect.Type, value string) (*reflect.Value, error) {
	Value, ok := p.m[ftype.Elem().Kind()]
	if !ok {
		return nil, fmt.Errorf("no value parser for : %s", ftype.String())
	}

	return Value(ftype, value)
}

func newPtrParser() *PtrParser {
	return &PtrParser{
		m: map[reflect.Kind]Parser{
			reflect.Bool:    boolPtrValue,
			reflect.Int:     intPtrValue,
			reflect.Int8:    int8PtrValue,
			reflect.Int16:   int16PtrValue,
			reflect.Int32:   int32PtrValue,
			reflect.Int64:   int64PtrValue,
			reflect.Uint:    uintPtrValue,
			reflect.Uint8:   uint8PtrValue,
			reflect.Uint16:  uint16PtrValue,
			reflect.Uint32:  uint32PtrValue,
			reflect.Uint64:  uint64PtrValue,
			reflect.Float32: float32PtrValue,
			reflect.Float64: float64PtrValue,
			reflect.Map:     mapPtrValue,
			reflect.Slice:   slicePtrValue,
			reflect.String:  stringPtrValue,
		},
	}
}

func asRef(v reflect.Value) *reflect.Value {
	return &v
}

func boolValue(ftype reflect.Type, value string) (*reflect.Value, error) {
	v, err := strconv.ParseBool(value)
	if err != nil {
		return nil, err
	}
	return asRef(reflect.ValueOf(v)), nil
}

func parseInt(ftype reflect.Type, value string) (int, error) {
	v, err := strconv.ParseInt(value, 10, int(ftype.Size()*8))
	if err != nil {
		return 0, err
	}
	return int(v), nil
}

func parseInt8(ftype reflect.Type, value string) (int8, error) {
	v, err := strconv.ParseInt(value, 10, int(ftype.Size()*8))
	if err != nil {
		return 0, err
	}
	return int8(v), nil
}

func parseInt16(ftype reflect.Type, value string) (int16, error) {
	v, err := strconv.ParseInt(value, 10, int(ftype.Size()*8))
	if err != nil {
		return 0, err
	}
	return int16(v), nil
}

func parseInt32(ftype reflect.Type, value string) (int32, error) {
	v, err := strconv.ParseInt(value, 10, int(ftype.Size()*8))
	if err != nil {
		return 0, err
	}
	return int32(v), nil
}

func parseInt64(ftype reflect.Type, value string) (int64, error) {
	v, err := strconv.ParseInt(value, 10, int(ftype.Size()*8))
	if err != nil {
		return 0, err
	}
	return int64(v), nil
}

func parseUint(ftype reflect.Type, value string) (uint, error) {
	v, err := strconv.ParseUint(value, 10, int(ftype.Size()*8))
	if err != nil {
		return 0, err
	}
	return uint(v), nil
}

func parseUint8(ftype reflect.Type, value string) (uint8, error) {
	v, err := strconv.ParseUint(value, 10, int(ftype.Size()*8))
	if err != nil {
		return 0, err
	}
	return uint8(v), nil
}

func parseUint16(ftype reflect.Type, value string) (uint16, error) {
	v, err := strconv.ParseUint(value, 10, int(ftype.Size()*8))
	if err != nil {
		return 0, err
	}
	return uint16(v), nil
}

func parseUint32(ftype reflect.Type, value string) (uint32, error) {
	v, err := strconv.ParseUint(value, 10, int(ftype.Size()*8))
	if err != nil {
		return 0, err
	}
	return uint32(v), nil
}

func parseUint64(ftype reflect.Type, value string) (uint64, error) {
	v, err := strconv.ParseUint(value, 10, int(ftype.Size()*8))
	if err != nil {
		return 0, err
	}
	return uint64(v), nil
}

func intValue(ftype reflect.Type, value string) (*reflect.Value, error) {
	v, err := parseInt(ftype, value)
	if err != nil {
		return nil, err
	}
	return asRef(reflect.ValueOf(v)), nil
}

func int8Value(ftype reflect.Type, value string) (*reflect.Value, error) {
	v, err := parseInt8(ftype, value)
	if err != nil {
		return nil, err
	}
	return asRef(reflect.ValueOf(v)), nil
}

func int16Value(ftype reflect.Type, value string) (*reflect.Value, error) {
	v, err := parseInt16(ftype, value)
	if err != nil {
		return nil, err
	}
	return asRef(reflect.ValueOf(v)), nil
}

func int32Value(ftype reflect.Type, value string) (*reflect.Value, error) {
	v, err := parseInt32(ftype, value)
	if err != nil {
		return nil, err
	}
	return asRef(reflect.ValueOf(v)), nil
}

func int64Value(ftype reflect.Type, value string) (*reflect.Value, error) {
	v, err := parseInt64(ftype, value)
	if err != nil {
		return nil, err
	}
	return asRef(reflect.ValueOf(v)), nil
}

func uintValue(ftype reflect.Type, value string) (*reflect.Value, error) {
	v, err := parseUint(ftype, value)
	if err != nil {
		return nil, err
	}
	return asRef(reflect.ValueOf(v)), nil
}

func uint8Value(ftype reflect.Type, value string) (*reflect.Value, error) {
	v, err := parseUint8(ftype, value)
	if err != nil {
		return nil, err
	}
	return asRef(reflect.ValueOf(v)), nil
}

func uint16Value(ftype reflect.Type, value string) (*reflect.Value, error) {
	v, err := parseUint16(ftype, value)
	if err != nil {
		return nil, err
	}
	return asRef(reflect.ValueOf(v)), nil
}

func uint32Value(ftype reflect.Type, value string) (*reflect.Value, error) {
	v, err := parseUint32(ftype, value)
	if err != nil {
		return nil, err
	}
	return asRef(reflect.ValueOf(v)), nil
}

func uint64Value(ftype reflect.Type, value string) (*reflect.Value, error) {
	v, err := parseUint64(ftype, value)
	if err != nil {
		return nil, err
	}
	return asRef(reflect.ValueOf(v)), nil
}

func float32Value(ftype reflect.Type, value string) (*reflect.Value, error) {
	v, err := parseFloat32(ftype, value)
	if err != nil {
		return nil, err
	}
	return asRef(reflect.ValueOf(v)), nil
}

func parseFloat32(ftype reflect.Type, value string) (float32, error) {
	v, err := strconv.ParseFloat(value, int(ftype.Size()*8))
	if err != nil {
		return 0, err
	}

	return float32(v), nil
}

func float64Value(ftype reflect.Type, value string) (*reflect.Value, error) {
	v, err := parseFloat64(ftype, value)
	if err != nil {
		return nil, err
	}
	return asRef(reflect.ValueOf(v)), nil
}

func parseFloat64(ftype reflect.Type, value string) (float64, error) {
	v, err := strconv.ParseFloat(value, int(ftype.Size()*8))
	if err != nil {
		return 0, err
	}

	return float64(v), nil
}

func stringValue(ftype reflect.Type, value string) (*reflect.Value, error) {
	return asRef(reflect.ValueOf(value)), nil
}

func mapValue(ftype reflect.Type, value string) (*reflect.Value, error) {
	sp := mapParser{}
	vmap, err := sp.parse(ftype.Key(), ftype.Elem(), value)

	if err != nil {
		return nil, err
	}

	m := reflect.MakeMap(ftype)
	for k, v := range vmap {
		m.SetMapIndex(k, *v)
	}

	return asRef(m), nil
}

type mapParser struct{}

func (s *mapParser) parse(ktype reflect.Type, vtype reflect.Type, value string) (map[reflect.Value]*reflect.Value, error) {
	if err := removeParen(&value, "{", "}"); err != nil {
		return nil, err
	}

	ret := make(map[reflect.Value]*reflect.Value, 0)

	arr, err := split(value, rune(','))
	if err != nil {
		return nil, err
	}

	for _, str := range arr {
		str = strings.Trim(str, CutSet)

		subArr, err := split(str, rune(':'))
		if err != nil {
			return nil, err
		}

		if len(subArr) != 2 {
			return nil, fmt.Errorf("no key : value")
		}

		if ktype.Kind() == reflect.String {
			removeParen(&subArr[0], `"`, `"`)
		}

		if vtype.Kind() == reflect.String {
			removeParen(&subArr[1], `"`, `"`)
		}

		key, err := Value(ktype, subArr[0])
		if err != nil {
			return nil, err
		}

		value, err := Value(vtype, subArr[1])
		if err != nil {
			return nil, err
		}

		ret[*key] = value
	}

	return ret, nil
}

func removeParen(str *string, pre string, suf string) error {
	*str = strings.Trim(*str, CutSet)
	if !strings.HasPrefix(*str, pre) {
		return fmt.Errorf("slice format error : no prefix %s target %+v", pre, *str)
	}

	if !strings.HasSuffix(*str, suf) {
		return fmt.Errorf("slice format error : no suffix %s target %+v", suf, *str)
	}

	*str = strings.Trim((*str)[len(pre):len(*str)-len(suf)], CutSet)
	return nil
}

func sliceValue(ftype reflect.Type, value string) (*reflect.Value, error) {
	sp := sliceParser{}
	arr, err := sp.parse(ftype.Elem(), value)

	if err != nil {
		return nil, err
	}

	s := reflect.MakeSlice(ftype, len(arr), len(arr))
	for i, v := range arr {
		s.Index(i).Set(*v)
	}

	return asRef(s), nil
}

type sliceParser struct{}

func (s *sliceParser) parse(ftype reflect.Type, value string) ([]*reflect.Value, error) {
	if err := removeParen(&value, "[", "]"); err != nil {
		return nil, err
	}

	ret := make([]*reflect.Value, 0)
	arr, err := split(value, rune(','))

	if err != nil {
		return nil, err
	}

	for _, str := range arr {
		str = strings.Trim(str, CutSet)
		if ftype.Kind() == reflect.String {
			err := removeParen(&str, `"`, `"`)
			if err != nil {
				return nil, err
			}
		}
		v, err := Value(ftype, str)
		if err != nil {
			return nil, err
		}

		ret = append(ret, v)
	}

	return ret, nil
}

func split(value string, by rune) ([]string, error) {
	sp := stringParser{}
	if err := sp.parse(value, by); err != nil {
		return nil, err
	}

	return sp.blocks, nil
}

func boolPtrValue(ftype reflect.Type, value string) (*reflect.Value, error) {
	v, err := strconv.ParseBool(value)
	if err != nil {
		return nil, err
	}
	return asRef(reflect.ValueOf(&v)), nil
}

func intPtrValue(ftype reflect.Type, value string) (*reflect.Value, error) {
	v, err := parseInt(ftype, value)
	if err != nil {
		return nil, err
	}
	return asRef(reflect.ValueOf(&v)), nil
}

func int8PtrValue(ftype reflect.Type, value string) (*reflect.Value, error) {
	v, err := parseInt8(ftype, value)
	if err != nil {
		return nil, err
	}
	return asRef(reflect.ValueOf(&v)), nil
}

func int16PtrValue(ftype reflect.Type, value string) (*reflect.Value, error) {
	v, err := parseInt16(ftype, value)
	if err != nil {
		return nil, err
	}
	return asRef(reflect.ValueOf(&v)), nil
}

func int32PtrValue(ftype reflect.Type, value string) (*reflect.Value, error) {
	v, err := parseInt32(ftype, value)
	if err != nil {
		return nil, err
	}
	return asRef(reflect.ValueOf(&v)), nil
}

func int64PtrValue(ftype reflect.Type, value string) (*reflect.Value, error) {
	v, err := parseInt64(ftype, value)
	if err != nil {
		return nil, err
	}
	return asRef(reflect.ValueOf(&v)), nil
}

func uintPtrValue(ftype reflect.Type, value string) (*reflect.Value, error) {
	v, err := parseUint(ftype, value)
	if err != nil {
		return nil, err
	}
	return asRef(reflect.ValueOf(&v)), nil
}

func uint8PtrValue(ftype reflect.Type, value string) (*reflect.Value, error) {
	v, err := parseUint8(ftype, value)
	if err != nil {
		return nil, err
	}
	return asRef(reflect.ValueOf(&v)), nil
}

func uint16PtrValue(ftype reflect.Type, value string) (*reflect.Value, error) {
	v, err := parseUint16(ftype, value)
	if err != nil {
		return nil, err
	}
	return asRef(reflect.ValueOf(&v)), nil
}

func uint32PtrValue(ftype reflect.Type, value string) (*reflect.Value, error) {
	v, err := parseUint32(ftype, value)
	if err != nil {
		return nil, err
	}
	return asRef(reflect.ValueOf(&v)), nil
}

func uint64PtrValue(ftype reflect.Type, value string) (*reflect.Value, error) {
	v, err := parseUint64(ftype, value)
	if err != nil {
		return nil, err
	}
	return asRef(reflect.ValueOf(&v)), nil
}

func float32PtrValue(ftype reflect.Type, value string) (*reflect.Value, error) {
	v, err := parseFloat32(ftype, value)
	if err != nil {
		return nil, err
	}
	return asRef(reflect.ValueOf(&v)), nil
}

func float64PtrValue(ftype reflect.Type, value string) (*reflect.Value, error) {
	v, err := parseFloat64(ftype, value)
	if err != nil {
		return nil, err
	}
	return asRef(reflect.ValueOf(&v)), nil
}

func stringPtrValue(ftype reflect.Type, value string) (*reflect.Value, error) {
	return asRef(reflect.ValueOf(&value)), nil
}

func mapPtrValue(ftype reflect.Type, value string) (*reflect.Value, error) {
	return asRef(reflect.ValueOf(value)), nil
}

func slicePtrValue(ftype reflect.Type, value string) (*reflect.Value, error) {
	return asRef(reflect.ValueOf(value)), nil
}
