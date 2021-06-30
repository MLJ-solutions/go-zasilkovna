package models

import (
	"encoding/xml"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)

type CreatePacketClaim struct {
	XMLName         xml.Name        `xml:"createPacketClaim"`
	ApiPassword     string          `xml:"apiPassword" validate:"required"`
	ClaimAttributes ClaimAttributes `xml:"claimAttributes" validate:"required"`
}

type ClaimAttributes struct {
	Id               int     `xml:"id,omitempty"`
	Number           string  `xml:"number" validate:"required"`
	Email            string  `xml:"email" validate:"required,email"`
	Phone            string  `xml:"phone" validate:"required"`
	Value            float32 `xml:"value" validate:"required"`
	Currency         string  `xml:"currency,omitempty"`
	Eshop            string  `xml:"eshop" validate:"required"` // TODO - Required when using more senders
	SendLabelToEmail bool    `xml:"SendLabelToEmail,omitempty"`
}

func NewCreatePacketClaim(apiPassword string, c ClaimAttributes) *CreatePacketClaim {
	return &CreatePacketClaim{ApiPassword: apiPassword, ClaimAttributes: c}
}

func NewClaimAttributesRequired(Id int, Number string, Email string, Phone string, Value float32,
	Eshop string) *ClaimAttributes {
	return &ClaimAttributes{
		Id:               Id,
		Number:           Number,
		Email:            Email,
		Phone:            Phone,
		Value:            Value,
		Eshop:            Eshop,
		SendLabelToEmail: true, // We always want to send label to recipient via email
	}
}

func NewClaimAttributes(Id int, Number string, Email string, Phone string, Value float32, Currency string,
	Eshop string) *ClaimAttributes {
	return &ClaimAttributes{
		Id:               Id,
		Number:           Number,
		Email:            Email,
		Phone:            Phone,
		Value:            Value,
		Currency:         Currency,
		Eshop:            Eshop,
		SendLabelToEmail: true, // We always want to send email to customer
	}
}

func ValidateClaimAttributes() (isValidated bool, errorsArray []validator.FieldError) {
	fmt.Println("---ClaimAttributes---")

	v := validator.New()
	a := ClaimAttributes{
		Id:     10,
		Number: "number",
		Email:  "a@b.cz",
		Value:  10.0,
		Phone:  "123456789",
		Eshop:  "furybeans.cz",
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