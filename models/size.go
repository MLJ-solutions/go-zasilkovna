package models

import (
	"gopkg.in/go-playground/validator.v9"
	"log"
)

type Size struct {
	Length int `xml:"length" validate:"required,max=120,min=10"`
	Width  int `xml:"width" validate:"required,max=120,min=7"`
	Height int `xml:"height" validate:"required,max=120,min=1"`
}

func NewSize(Length int, Width int, Height int) *Size {
	return &Size{
		Length: Length,
		Width:  Width,
		Height: Height,
	}
}

func ValidateSize(s Size) (isValidated bool, errorsArray []validator.FieldError) {
	v := validator.New()
	a := Size{
		Length: s.Length,
		Width:  s.Width,
		Height: s.Height,
	}

	isSumSizeOk := s.Length+s.Width+s.Height <= 150

	err := v.Struct(a)
	if err != nil { // If err contains errors, params are not validated
		isValidated = false
		for _, e := range err.(validator.ValidationErrors) {
			errorsArray = append(errorsArray, e)
		}
		return
	} else if !isSumSizeOk {
		isValidated = false
		log.Println("Sum of all 3 sizes is greater than 150!!!")
		return
	} else {
		isValidated = true
		return
	}
}
