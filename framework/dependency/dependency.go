package dependency

import (
	"fmt"
	"godi/framework/trace"
	"reflect"
	"strconv"
	"time"
)

type Factory func() interface{}

type dependency struct {
	binds     []bind
	factories map[string]Factory
	beans     map[string]interface{}
}

type bind struct {
	client  string
	field   string
	subject string
}

var defaultDependency = &dependency{
	factories: make(map[string]Factory),
	beans:     make(map[string]interface{}),
}

func FactoryRegister(f Factory) {
	trace.Trace()
	defaultDependency.factories[trace.PathFileWithPC(1)] = f
}

func Load(filename string) error {
	return parseConfigFile(filename)
}

func Create() error {
	return defaultDependency.Create()
}

func Inject() error {
	return defaultDependency.Inject()
}

func (o *dependency) Create() error {
	for _, b := range o.binds {
		o.beans[b.client] = nil
		o.beans[b.subject] = nil
	}

	for name := range o.beans {
		if factory, ok := o.factories[name]; ok {
			o.beans[name] = factory()
			fmt.Println("created : ", name)
		}
	}

	return nil
}

func (o *dependency) Inject() error {
	for _, iv := range o.binds {
		client := o.beans[iv.client]
		if client == nil {
			return fmt.Errorf("inject error - no object : %s", iv.client)
		}

		elem := reflect.ValueOf(client).Elem()
		field := elem.FieldByName(iv.field)
		ftype := field.Type()

		if !field.CanSet() {
			return fmt.Errorf("inject error - cannot set object %s, field %s", iv.client, iv.field)
		}

		subject := o.beans[iv.subject]
		if subject != nil {
			rs := reflect.ValueOf(subject)
			if rs.Type().ConvertibleTo(field.Type()) {
				field.Set(rs)
				continue
			}
		}

		ok, err := setWellKnownType(&field, ftype, iv.subject)
		if err != nil {
			return err
		}

		if ok {
			fmt.Println("injected : ", iv.client, " ", iv.field, " ", iv.subject)
			continue
		}

		if _, err = setPrimitiveType(&field, ftype, iv.subject); err != nil {
			return err
		}

		fmt.Println("injected : ", iv.client, " ", iv.field, " ", iv.subject)
	}
	return nil
}

func setWellKnownType(field *reflect.Value, ftype reflect.Type, subject string) (bool, error) {
	switch ftype.String() { // if filed is well-known type
	case "time.Time":
		t, err := time.Parse("2006-01-02T15:04:05Z", subject)
		if err != nil {
			return false, fmt.Errorf("inject error - not compatable type %s, field %s subject %s", ftype, field, subject)
		}
		field.Set(reflect.ValueOf(t))
		return true, nil
	case "time.Duration":
		dur, err := time.ParseDuration(subject)
		if err != nil {
			return false, fmt.Errorf("inject error - not compatable type %s, field %s subject %s", ftype, field, subject)
		}
		field.SetInt(int64(dur))
		return true, nil
	}
	return false, nil
}

func setPrimitiveType(field *reflect.Value, ftype reflect.Type, subject string) (bool, error) {
	switch ftype.Kind() {
	case reflect.String:
		field.SetString(subject)
	case reflect.Bool:
		v, err := strconv.ParseBool(subject)
		if err != nil {
			return false, fmt.Errorf("inject error - not compatable type %s, field %s subject %s", ftype, field, subject)
		}
		field.SetBool(v)
		return true, nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v, err := strconv.ParseInt(subject, 10, int(ftype.Size()*8))
		if err != nil {
			return false, fmt.Errorf("inject error - not compatable type %s, field %s subject %s", ftype, field, subject)
		}
		field.SetInt(v)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		v, err := strconv.ParseUint(subject, 10, int(ftype.Size()*8))
		if err != nil {
			return false, fmt.Errorf("inject error - not compatable type %s, field %s subject %s", ftype, field, subject)
		}
		field.SetUint(v)
	case reflect.Float32:
		v, err := strconv.ParseFloat(subject, 32)
		if err != nil {
			return false, fmt.Errorf("inject error - not compatable type %s, field %s subject %s", ftype, field, subject)
		}
		field.SetFloat(v)
	case reflect.Float64:
		v, err := strconv.ParseFloat(subject, 64)
		if err != nil {
			return false, fmt.Errorf("inject error - not compatable type %s, field %s subject %s", ftype, field, subject)
		}
		field.SetFloat(v)

	default:
		return false, fmt.Errorf("inject error - not supported type %s, field %s", ftype, field)
	}

	return true, nil
}

func New(name string) {
	defaultDependency.beans[name] = nil
}

func Set(client string, field string, subject string) error {
	if client == "" || field == "" || subject == "" {
		return fmt.Errorf("bad parameter(s) client : %s field : %s subject : %s", client, field, subject)
	}

	if _, ok := defaultDependency.factories[client]; !ok {
		return fmt.Errorf("no client : %s", client)
	}

	defaultDependency.binds = append(defaultDependency.binds,
		bind{
			client,
			field,
			subject,
		})

	return nil
}

func Get(name string) interface{} {
	return defaultDependency.beans[name]
}
