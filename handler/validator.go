package handler

import (
	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

//InitValidator Initialise un nouveau validateur de demande
func InitValidator() {
	Validate = validator.New()
}
