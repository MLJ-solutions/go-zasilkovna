package models

import (
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)

type CourierInfoItem struct {
	CourierId              int                      `xml:"courier_id" validate:"required"`
	CourierName            string                   `xml:"courier_name" validate:"required"`
	CourierNumbers         []CourierNumbers         `xml:"courier_numbers,omitempty"`
	CourierBarcodes        []CourierBarcodes        `xml:"courier_barcodes,omitempty"`
	CourierTrackingNumbers []CourierTrackingNumbers `xml:"courier_tracking_numbers,omitempty"`
	CourierTrackingUrls    []CourierTrackingUrl     `xml:"courier_tracking_urls,omitempty"`
}

func ValidateCourierInfoItem() (isValidated bool, errorsArray []validator.FieldError) {
	fmt.Println("---CourierInfoItem---")

	v := validator.New()
	a := CourierInfoItem{
		CourierId:   10,
		CourierName: "Some courier's name",
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
