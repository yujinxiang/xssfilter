package xssfilter

import (
	"errors"
	"html"

	"reflect"
)

func XssFilter(data interface{}) (err error) {

	value := reflect.ValueOf(data)

	switch value.Kind() {
	case reflect.Ptr:
		PtrXssFilter(value.Interface())
	case reflect.Map:
		MapXssFilter(value.Interface())
	case reflect.Slice:
		SliceXssFilter(value.Interface())
	default:
		return errors.New("not support the type XssFilter")
	}

	return

}

func PtrXssFilter(data interface{}) {

	dvalue := reflect.ValueOf(data)

	if dvalue.Kind() != reflect.Ptr {
		panic("data must be ptr")
	}

	value := dvalue.Elem()
	dealNoMapType(value)

	return

}

func SliceXssFilter(data interface{}) {

	dvalue := reflect.ValueOf(data)
	if dvalue.Kind() != reflect.Slice {
		panic("data must be slice")
	}

	for i := 0; i < dvalue.Len(); i++ {
		value := dvalue.Index(i)
		dealNoMapType(value)
	}
}

func MapXssFilter(data interface{}) {

	dvalue := reflect.ValueOf(data)
	if dvalue.Kind() != reflect.Map {
		panic("data must be map")
	}

	for _, key := range dvalue.MapKeys() {
		value := dvalue.MapIndex(key)
		switch value.Kind() {
		case reflect.String:
			valueStr := html.EscapeString(value.Interface().(string))
			dvalue.SetMapIndex(key, reflect.ValueOf(valueStr))
		case reflect.Interface:
			newValue := reflect.ValueOf(value.Interface())
			if newValue.Kind() == reflect.String {
				valueStr := html.EscapeString(newValue.Interface().(string))
				dvalue.SetMapIndex(key, reflect.ValueOf(valueStr))
			} else {
				choiceXssFilterMethod(reflect.ValueOf(value.Interface()))
			}
		default:
			choiceXssFilterMethod(value)
		}
	}
}

func structXssFilter(data interface{}) {

	dvalue := reflect.ValueOf(data)
	elemValue := dvalue.Elem()

	if elemValue.Kind() != reflect.Struct {
		panic("data must be struct")
	}

	elemType := elemValue.Type()

	for i := 0; i < elemType.NumField(); i++ {
		value := elemValue.Field(i)
		dealNoMapType(value)
	}

	return

}

func arrayXssFilter(data interface{}) {

	dvalue := reflect.ValueOf(data).Elem()
	if dvalue.Kind() != reflect.Array {
		panic("data must be array")
	}

	for i := 0; i < dvalue.Len(); i++ {
		value := dvalue.Index(i)
		dealNoMapType(value)
	}
}

func dealNoMapType(value reflect.Value) {
	if value.Kind() == reflect.String {
		valueStr := html.EscapeString(value.Interface().(string))
		value.SetString(valueStr)
	} else if value.Kind() == reflect.Interface {
		newValue := reflect.ValueOf(value.Interface())
		if newValue.Kind() == reflect.String {
			valueStr := html.EscapeString(newValue.Interface().(string))
			value.Set(reflect.ValueOf(valueStr))
		} else {
			choiceXssFilterMethod(reflect.ValueOf(value.Interface()))
		}

	} else {
		choiceXssFilterMethod(value)
	}
}

func choiceXssFilterMethod(value reflect.Value) {
	switch value.Kind() {
	case reflect.Ptr:
		PtrXssFilter(value.Interface())
	case reflect.Struct:
		newValue := value.Addr()
		structXssFilter(newValue.Interface())
	case reflect.Map:
		MapXssFilter(value.Interface())
	case reflect.Slice:
		SliceXssFilter(value.Interface())
	case reflect.Array:
		newValue := value.Addr()
		arrayXssFilter(newValue.Interface())
	}
}
