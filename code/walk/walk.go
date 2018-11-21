package walk

import (
	"reflect"
)

func getValue(v interface{}) reflect.Value {
	val := reflect.ValueOf(v)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	return val
}

// Walk walks through a struct x and calls fn for all string fields.
func Walk(x interface{}, fn func(string)) {
	val := getValue(x)

	var fieldCount int
	var getField func(int) reflect.Value

	switch val.Kind() {
	case reflect.Struct:
		fieldCount = val.NumField()
		getField = val.Field
	case reflect.Slice:
		fieldCount = val.Len()
		getField = val.Index
	case reflect.String:
		fn(val.String())
	}

	for i := 0; i < fieldCount; i++ {
		field := getField(i)

		Walk(field.Interface(), fn)
	}
}
