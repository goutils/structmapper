package structmapper

import "reflect"

func main() {

}

func getValueObject(v interface{}) reflect.Value {
	return reflect.New(reflect.TypeOf(v)).Elem()
}

func fromFieldValues(v reflect.Value, index int) (reflect.Value, reflect.StructField, string){
	return v.Field(index), v.Type().Field(index), v.Type().Field(index).Name
}

func AutoMap(from interface{}, to interface{}) (interface{}, error) {

	fromVal := reflect.ValueOf(from)
	toVal := getValueObject(to)

	for index := 0; index < fromVal.NumField(); index++ {

		fromField, fromFieldType, fromFieldTypeName := fromFieldValues(fromVal, index)
		toFieldType, toFieldExist := toVal.Type().FieldByName(fromFieldTypeName)
		
		if toFieldExist && fromFieldType.Type == toFieldType.Type {
				toVal.FieldByName(fromFieldTypeName).Set(fromField)
		}
	}
	return toVal.Interface(), nil
}
