package models

import (
	"gopkg.in/go-playground/validator.v9"
)

type AttributeCollection struct {
	Attribute []Attribute `xml:"attribute"`
}

func NewAttributeCollection(Attribute []Attribute) *AttributeCollection {
	return &AttributeCollection{
		Attribute: Attribute,
	}
}

type Attribute struct {
	Key   string `xml:"key" validate:"required"`
	Value string `xml:"value" validate:"required"`
}

func NewAttribute(Key string, Value string) *Attribute {
	return &Attribute{
		Key:   Key,
		Value: Value,
	}
}

func ValidateAttribute(key string, value string) (isValidated bool, errorsArray []validator.FieldError) {
	v := validator.New()
	a := Attribute{
		Key:   key,
		Value: value,
	}
	err := v.Struct(a)
	if err != nil { // If err contains errors, params are not validated
		isValidated = false
		for _, e := range err.(validator.ValidationErrors) {
			errorsArray = append(errorsArray, e)
		}
		return
	} else {
		isValidated = true
		return
	}
}
