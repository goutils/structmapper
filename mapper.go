package main

import "reflect"

func GetValueObject(v interface{}) reflect.Value {
	return reflect.New(reflect.TypeOf(v)).Elem()
}

func fromFieldValues(v reflect.Value, index int) (reflect.Value, reflect.StructField, string){
	return v.Field(index), v.Type().Field(index), v.Type().Field(index).Name
}

func AutoMap(from interface{}, to interface{}) (interface{}, error) {

	fromVal := reflect.ValueOf(from)
	toVal := GetValueObject(to)

	for index := 0; index < fromVal.NumField(); index++ {

		fromField, fromFieldType, fromFieldTypeName := fromFieldValues(fromVal, index)

		toField := toVal.FieldByName(fromFieldTypeName)
		toFieldType, toFieldExist := toVal.Type().FieldByName(fromFieldTypeName)
		if toFieldExist && fromFieldType.Type == toFieldType.Type {
			if fromField.Type().Kind().String() == "struct" && toField.Type().Kind().String() == "struct" {
				field, err := AutoMap(fromField.Interface(), toField.Interface())
				if err != nil {
					return nil, err
				}
				toVal.FieldByName(fromFieldTypeName).Set(reflect.ValueOf(field))
			}else {
				toVal.FieldByName(fromFieldTypeName).Set(fromField)
			}
		}
	}
	return toVal.Interface(), nil
}
