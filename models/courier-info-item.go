package models

import (
	"encoding/xml"
	"gopkg.in/go-playground/validator.v9"
)

type CourierInfo struct {
	XMLName         xml.Name          `xml:"courierInfo"`
	CourierInfoItem []CourierInfoItem `xml:"courierInfoItem"`
}

type CourierInfoItem struct {
	XMLName                xml.Name                 `xml:"courierInfoItem"`
	CourierId              int                      `xml:"courierId" validate:"required"`
	CourierName            string                   `xml:"courierName" validate:"required"`
	CourierNumbers         []CourierNumbers         `xml:"courierNumbers,omitempty"`
	CourierBarcodes        []CourierBarcodes        `xml:"courierBarcodes,omitempty"`
	CourierTrackingNumbers []CourierTrackingNumbers `xml:"courierTrackingNumbers,omitempty"`
	CourierTrackingUrls    []CourierTrackingUrls    `xml:"courierTrackingUrls,omitempty"`
}

func NewCourierInfoItemRequired(CourierId int, CourierName string) *CourierInfoItem {
	return &CourierInfoItem{
		CourierId:   CourierId,
		CourierName: CourierName,
	}
}

func NewCourierInfoItem(CourierId int, CourierName string, CourierNumbers []CourierNumbers,
	CourierBarcodes []CourierBarcodes, CourierTrackingNumbers []CourierTrackingNumbers,
	CourierTrackingUrls []CourierTrackingUrls) *CourierInfoItem {
	return &CourierInfoItem{
		CourierId:              CourierId,
		CourierName:            CourierName,
		CourierNumbers:         CourierNumbers,
		CourierBarcodes:        CourierBarcodes,
		CourierTrackingNumbers: CourierTrackingNumbers,
		CourierTrackingUrls:    CourierTrackingUrls,
	}
}

func ValidateCourierInfoItem() (isValidated bool, errorsArray []validator.FieldError) {
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
