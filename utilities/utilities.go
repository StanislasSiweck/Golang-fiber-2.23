package utilities

import (
	"golang.org/x/crypto/bcrypt"
	"reflect"
)

// IsStruct vérifier si l'objet est une structure (pas un pointeur)
func IsStruct(object interface{}) bool {
	return reflect.ValueOf(object).Type().Kind() == reflect.Struct
}

/*infos
HashString:
Hashes a given string based on the bcrypt algorithm with bcrypt.DefaultCost
*/
func HashString(s string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hash)
}

/*
	CheckPassword:
	Given a hash and a string compares if they are the same
*/
func CheckPassword(hash []byte, s string) bool {
	err := bcrypt.CompareHashAndPassword(hash, []byte(s))
	if err != nil {
		return false
	}
	return true
}
