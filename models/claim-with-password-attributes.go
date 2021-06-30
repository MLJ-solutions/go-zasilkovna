package models

import (
	"encoding/xml"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)

type CreatePacketClaimWithPassword struct {
	XMLName                       xml.Name                    `xml:"createPacketClaimWithPassword"`
	ApiPassword                   string                      `xml:"apiPassword" validate:"required"`
	CreatePacketClaimWithPassword ClaimWithPasswordAttributes `xml:"createPacketClaimWithPassword" validate:"required"`
}

type ClaimWithPasswordAttributes struct {
	Id                  int     `xml:"id" validate:"required"`
	Number              string  `xml:"number" validate:"required"`
	Email               string  `xml:"email" validate:"required,email"`
	Phone               string  `xml:"phone" validate:"required"`
	Value               float32 `xml:"value" validate:"required"`
	Currency            string  `xml:"currency,omitempty"`
	Eshop               string  `xml:"eshop" validate:"required"`
	ConsignCountry      string  `xml:"consign_country" validate:"required"`
	SendEmailToCustomer bool    `xml:"send_email_to_customer,omitempty"`
}

func NewCreatePacketClaimWithPassword(apiPassword string, c ClaimWithPasswordAttributes) *CreatePacketClaimWithPassword {
	return &CreatePacketClaimWithPassword{ApiPassword: apiPassword, CreatePacketClaimWithPassword: c}
}

func NewClaimWithPasswordAttributesRequired(Id int, Number string, Email string, Phone string, Value float32,
	Eshop string, ConsignCountry string) *ClaimWithPasswordAttributes {
	return &ClaimWithPasswordAttributes{
		Id:                  Id,
		Number:              Number,
		Email:               Email,
		Phone:               Phone,
		Value:               Value,
		Eshop:               Eshop,
		ConsignCountry:      ConsignCountry,
		SendEmailToCustomer: true, // We always want to send email to customer
	}
}

func NewClaimWithPasswordAttributes(Id int, Number string, Email string, Phone string, Value float32, Currency string,
	Eshop string, ConsignCountry string) *ClaimWithPasswordAttributes {
	return &ClaimWithPasswordAttributes{
		Id:                  Id,
		Number:              Number,
		Email:               Email,
		Phone:               Phone,
		Value:               Value,
		Currency:            Currency,
		Eshop:               Eshop,
		ConsignCountry:      ConsignCountry,
		SendEmailToCustomer: true, // We always want to send email to customer
	}
}

func ValidateClaimWithPasswordAttributes() (isValidated bool, errorsArray []validator.FieldError) {
	fmt.Println("---ClaimWithPasswordAttributes---")

	v := validator.New()
	a := ClaimWithPasswordAttributes{
		Id:             10,
		Number:         "number",
		Email:          "a@b.cz",
		Value:          10.0,
		Phone:          "123456789",
		Eshop:          "furybeans.cz",
		ConsignCountry: "consigncountry",
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
