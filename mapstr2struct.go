package map2struct

import (
	"reflect"
	"strconv"
	"time"
)

// item, the struct interface u want to bind
// m, the map[string]string u want to decode
func DecodeSs(item interface{}, m map[string]string) error {
	reflectType := reflect.TypeOf(item)

	reflectValue := reflect.Indirect(reflect.ValueOf(item))

	if reflectType.Kind() == reflect.Ptr {
		reflectType = reflectType.Elem()
	}

	for i, n := 0, reflectType.NumField(); i < n; i++ {
		tag := reflectType.Field(i).Tag.Get("json")

		if tag == "" || tag == "-" {
			tag = reflectType.Field(i).Name
		}

		switch reflectType.Field(i).Type.Kind() {
		// string to int64
		case reflect.Int64:
			if value, ok := m[tag]; ok {
				if reflectValue.CanSet() {
					v, err := strconv.ParseInt(value, 10, 64)
					if err != nil {
						return err
					}
					reflectValue.FieldByName(reflectType.Field(i).Name).SetInt(v)
				}
			}
		// string to int
		case reflect.Int:
			if value, ok := m[tag]; ok {
				if reflectValue.CanSet() {
					v, err := strconv.Atoi(value)
					if err != nil {
						return err
					}
					reflectValue.FieldByName(reflectType.Field(i).Name).SetInt(int64(v))
				}
			}
		// string to string
		case reflect.String:
			if value, ok := m[tag]; ok {
				if reflectValue.CanSet() {
					reflectValue.FieldByName(reflectType.Field(i).Name).SetString(value)
				}
			}
		// string to time
		case reflect.Struct:
			if _, ok := reflectValue.FieldByName(reflectType.Field(i).Name).Interface().(time.Time); ok {
				if value, ok := m[tag]; ok {
					moment, err := time.Parse(time.RFC3339, value)
					if err != nil {
						return err
					}
					reflectValue.FieldByName(reflectType.Field(i).Name).Set(reflect.ValueOf(moment))
				}
			}
		}
	}

	return nil
}
