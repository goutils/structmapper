package structmapper

import "reflect"

func main() {

}

func fromFieldValues(v reflect.Value, index int) (reflect.Value, reflect.StructField, string) {
	return v.Field(index), v.Type().Field(index), v.Type().Field(index).Name
}

func AutoMap(from interface{}, to interface{}) error {

	fromVal := reflect.ValueOf(from)
	toVal := reflect.ValueOf(to).Elem()

	for index := 0; index < fromVal.NumField(); index++ {

		fromField, fromFieldType, fromFieldTypeName := fromFieldValues(fromVal, index)
		toFieldType, toFieldExist := toVal.Type().FieldByName(fromFieldTypeName)

		if toFieldExist && fromFieldType.Type == toFieldType.Type {
			toVal.FieldByName(fromFieldTypeName).Set(fromField)
		}
	}
	return nil
}
