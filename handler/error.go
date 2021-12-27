package handler

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"runtime/debug"
	"strings"
)

// SendError given returns a fiber.error with the given code, message details and error
func SendError(code int, message, details string, err ...error) error {
	if len(err) > 0 {
		return fiber.NewError(
			code, strings.Join([]string{message, details, err[0].Error()}, ";"),
		)
	} else {
		return fiber.NewError(
			code, strings.Join([]string{message, details}, ";"),
		)
	}
}

func recovery() {
	if err := recover(); err != nil {
		stacktrace := string(debug.Stack())
		log.Panic(err, "\n", stacktrace)
	}
}
