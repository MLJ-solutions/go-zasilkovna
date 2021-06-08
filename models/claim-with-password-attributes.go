package models

type ClaimWithPasswordAttributes struct {
	Id                  int     `json:"id" validate:"required"`
	Number              string  `json:"number" validate:"required"`
	Email               string  `json:"email" validate:"required,email"`
	Phone               string  `json:"phone" validate:"required"`
	Value               float32 `json:"value" validate:"required"`
	Currency            string  `json:"currency,omitempty"`
	Eshop               string  `json:"eshop" validate:"required"`
	ConsignCountry      string  `json:"consign_country" validate:"required"`
	SendEmailToCustomer bool    `json:"send_email_to_customer,omitempty"`
}
