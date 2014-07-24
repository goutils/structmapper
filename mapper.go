package structmapper

import (
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

		if toFieldExist && fromFieldKind == toFieldType.Type.Kind() {
			switch fromFieldKind {
			case reflect.Struct:
					toStructField := reflect.New(toFieldType.Type)
					AutoMap(fromField.Interface(), toStructField.Interface())
					toVal.FieldByName(fromFieldTypeName).Set(toStructField.Elem())
			case reflect.Slice:
				if reflect.ValueOf(fromField.Type()).Interface() == toFieldType.Type {
					toVal.FieldByName(fromFieldTypeName).Set(fromField)
				} else {
					toFieldElemType := toVal.FieldByName(fromFieldTypeName).Type().Elem()
					if fromField.Type().Elem().Kind() == toFieldElemType.Kind() {

						toSliceField := reflect.New(toFieldType.Type).Elem()
						for i := 0; i < fromField.Len(); i++ {
							fromFieldElem := fromField.Index(i)
							toFieldElem := reflect.New(toFieldElemType)
							AutoMap(fromFieldElem.Interface(), toFieldElem.Interface())
							toSliceField = reflect.Append(toSliceField, toFieldElem.Elem())
						}

						toVal.FieldByName(fromFieldTypeName).Set(toSliceField)
					}
				}
			default:
				toVal.FieldByName(fromFieldTypeName).Set(fromField)
			}
		}

	}

	return nil
}
