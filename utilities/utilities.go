package utilities

import "reflect"

// IsStruct vérifier si l'objet est une structure (pas un pointeur)
func IsStruct(object interface{}) bool {
	return reflect.ValueOf(object).Type().Kind() == reflect.Struct
}
