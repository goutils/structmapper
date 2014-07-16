package structmapper

import (
	"log"
	"reflect"
)

func main() {
}

func fromFieldValues(v reflect.Value, index int) (reflect.Value, reflect.Kind, string) {
	return v.Field(index), v.Field(index).Kind(), v.Type().Field(index).Name
}

func AutoMap(from interface{}, to interface{}) error {
	fromVal := reflect.ValueOf(from)
	toVal := reflect.ValueOf(to).Elem()

	for index := 0; index < fromVal.NumField(); index++ {

		fromField, fromFieldKind, fromFieldTypeName := fromFieldValues(fromVal, index)
		toFieldType, toFieldExist := toVal.Type().FieldByName(fromFieldTypeName)

		if toFieldExist {
			switch fromFieldKind {
			case reflect.Struct:
				toStructField := reflect.New(toFieldType.Type)
				AutoMap(fromField, toStructField.Interface())
			case reflect.Slice:
				if reflect.ValueOf(fromField.Type()).Interface() == toFieldType.Type {
					log.Println("copying value", fromFieldTypeName)
					toVal.FieldByName(fromFieldTypeName).Set(fromField)
				} else {
					// toSliceField := reflect.New(toFieldType.Type)
					s := reflect.ValueOf(fromField.Interface())

					for i := 0; i < s.Len(); i++ {
						log.Println(s.Index(i).Type())

					}
				}
			default:
				toVal.FieldByName(fromFieldTypeName).Set(fromField)
			}
		}
	}

	return nil
}
