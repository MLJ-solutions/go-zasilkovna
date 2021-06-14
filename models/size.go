package models

import (
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)

type Size struct {
	Length int `xml:"length" validate:"required,max=120,min=10"`
	Width  int `xml:"width" validate:"required,max=120,min=7"`
	Height int `xml:"height" validate:"required,max=120,min=1"`
}

func ValidateSize(s Size) (isValidated bool, errorsArray []validator.FieldError) {
	fmt.Println("---Size---")

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
		fmt.Println("Sum of all 3 sizes is greater than 150!!!")
		return
	} else {
		isValidated = true
		return
	}
}
