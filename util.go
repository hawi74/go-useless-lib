package lib

import (
	"github.com/dannyvankooten/vat"
)

func Foobar(n string) bool {
	valid, err := vat.ValidateNumber("NL123456789B01")
	if err != nil {
		return false
	}
	return valid
}
