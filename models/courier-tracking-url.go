package models

import (
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)

type CourierTrackingUrl struct {
	Lang string `xml:"lang" validate:"required"`
	Url  string `xml:"url" validate:"required"`
}

func ValidateCourierTrackingUrl() (isValidated bool, errorsArray []validator.FieldError) {
	fmt.Println("---CourierTrackingUrl---")

	v := validator.New()
	a := CourierTrackingUrl{
		Lang: "English",
		Url:  "www.url.com",
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
