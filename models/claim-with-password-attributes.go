package models

import (
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)

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
