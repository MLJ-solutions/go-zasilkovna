package models

import (
	"fmt"
	"gopkg.in/go-playground/validator.v9"
	"time"
)

type CurrentStatusRecord struct {
	DateTime             ZasilkovnaDateTime `xml:"date_time" validate:"required"`
	StatusCode           int                `xml:"status_code" validate:"required"`
	CodeText             string             `xml:"code_text" validate:"required"`
	StatusText           string             `xml:"status_text" validate:"required"`
	BranchId             int                `xml:"branch_id" validate:"required"`
	DestinationBranchId  int                `xml:"destination_branch_id" validate:"required"`
	ExternalTrackingCode string             `xml:"external_tracking_code,omitempty"`
	IsReturning          *bool              `xml:"is_returning" validate:"required"` // *bool bcs of required
	StoredUntil          ZasilkovnaDate     `xml:"stored_until" validate:"required"`
	CarrierId            int                `xml:"carrier_id,omitempty"`
	CarrierName          string             `xml:"carrier_name,omitempty"`
}

func ValidateCurrentStatusRecord() (isValidated bool, errorsArray []validator.FieldError) {
	fmt.Println("---CurrentStatusRecord---")

	v := validator.New()
	f := false
	a := CurrentStatusRecord{
		DateTime:            ZasilkovnaDateTime(time.Now()),
		StatusCode:          111,
		CodeText:            "Some code text",
		StatusText:          "Some status text",
		BranchId:            222,
		DestinationBranchId: 333,
		IsReturning:         &f, // Cannot set *bool type directly, because is pointer to memory
		StoredUntil:         ZasilkovnaDate(time.Now()),
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
