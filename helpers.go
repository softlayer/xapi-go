package xenserver

import (
	"github.com/kolo/xmlrpc"
	"log"
	"reflect"
	"strconv"
)

func unMarshallXmlRPC(in xmlrpc.Struct, out interface{}) {
	ov := reflect.Indirect(reflect.ValueOf(out))
	iv := reflect.ValueOf(in["Value"])

	if ov.Kind() == reflect.Struct {
		setStruct(in["Value"].(xmlrpc.Struct), ov)
		return
	}

	switch in["Value"].(type) {
	case string:
		ov.Set(iv)
		return
	case []interface{}:
		l := iv.Len()
		ov.Set(reflect.MakeSlice(ov.Type(), l, l))
		for i, a := range in["Value"].([]interface{}) {
			if a == nil {
				continue
			}

			if ov.Index(i).Kind() == reflect.Struct {
				setStruct(a.(xmlrpc.Struct), ov.Index(i))
			} else {
				ov.Index(i).Set(reflect.ValueOf(a))
			}

		}
		return
	case xmlrpc.Struct:
		setStruct(in, ov)
	}
}

func setStruct(in xmlrpc.Struct, str reflect.Value) {
	for k, v := range in {
		field := str.FieldByName(UF(k))
		if !field.CanSet() || !field.IsValid() || v == nil {
			continue
		}

		switch field.Kind() {
		case reflect.Int:
			setInt(v, field)
		case reflect.Float64:
			setFloat64(v, field)
		case reflect.Bool:
			field.SetBool(v.(bool))
		case reflect.String:
			setString(v, field)
		case reflect.Map:
			field.Set(reflect.MakeMap(field.Type()))
			for a, b := range v.(xmlrpc.Struct) {
				if a == "" || b == nil {
					continue
				}
				field.SetMapIndex(reflect.ValueOf(a), reflect.ValueOf(b))
			}
		case reflect.Slice:
			l := reflect.ValueOf(v).Len()
			field.Set(reflect.MakeSlice(field.Type(), l, l))
			for i, a := range v.([]interface{}) {
				if a == nil {
					continue
				}
				field.Index(i).Set(reflect.ValueOf(a))
			}
		default:
			field.Set(reflect.ValueOf(v))
		}
	}

	return
}

func setString(in interface{}, out reflect.Value) {
	switch in.(type) {
	case string:
		out.SetString(in.(string))
	default:
		log.Printf("TODO: String %#v", in)
	}
}

func setFloat64(in interface{}, out reflect.Value) {
	var f float64
	switch in.(type) {
	case string:
		f, _ = strconv.ParseFloat(in.(string), 0)
	default:
		f = in.(float64)
	}
	out.SetFloat(f)
}

func setInt(in interface{}, out reflect.Value) {
	var f int64
	switch in.(type) {
	case string:
		f, _ = strconv.ParseInt(in.(string), 0, 0)
	case int:
		f = int64(in.(int))
	case uint:
		f = int64(in.(uint))
	case int32:
		f = int64(in.(int32))
	case int64:
		f = in.(int64)
	}
	out.SetInt(f)
}
