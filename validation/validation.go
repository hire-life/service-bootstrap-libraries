package validation

import (
	"github.com/go-playground/validator/v10"
)

type Validatable interface {
	Validate() error
}

type Error struct {
	Field   string
	Rule    string
	Message string
}

func Validate[I any](i *I) error {
	return validator.New().Struct(i)
}
