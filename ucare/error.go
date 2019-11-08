package ucare

import (
	"errors"
	"fmt"
)

// API response errors
var (
	ErrInvalidAuthCreds = errors.New("Incorrect authentication credentials")
	ErrAuthForbidden    = errors.New("Simple authentication over HTTP is " +
		"forbidden. Please, use HTTPS or signed requests instead")
	ErrInvalidVersion = errors.New("This feature is not support. " +
		"Try to change the version (refer to " +
		"https://uploadcare.com/api-refs/rest-api/v0.6.0/ for " +
		"more information on what methods belongs to what version).")
	ErrFileTooLarge = errors.New("Direct uploads only support " +
		"files smaller than 100MB")
)

type respErr struct {
	Details string `json:"detail"`
}

// Error implements error interface
func (e respErr) Error() string {
	return e.Details
}

type authErr struct{ respErr }

type throttleErr struct {
	RetryAfter int
}

func (e throttleErr) Error() string {
	if e.RetryAfter == 0 {
		return "Request was throttled."
	}
	return fmt.Sprintf(
		"Request was throttled. Expected available in %d second",
		e.RetryAfter,
	)
}

type reqValidationErr struct{ respErr }

func (e reqValidationErr) Error() string {
	return fmt.Sprintf("Request parameters validation error: %s", e.Details)
}

type reqForbiddenErr struct{ respErr }
