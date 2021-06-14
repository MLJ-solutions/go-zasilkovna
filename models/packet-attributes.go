package models

import (
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)

type PacketAttributes struct {
	Id                  int                   `xml:"id,omitempty"`
	Number              string                `xml:"number" validate:"required"`
	Name                string                `xml:"name" validate:"required"`
	Surname             string                `xml:"surname" validate:"required"`
	Company             string                `xml:"company,omitempty"`
	Email               string                `xml:"email" validate:"required"`
	Phone               string                `xml:"phone" validate:"required"`
	AddressId           int                   `xml:"address_id" validate:"required"`
	Currency            string                `xml:"currency,omitempty"`
	Cod                 float32               `xml:"cod,omitempty"`
	Value               float32               `xml:"value" validate:"required"`
	Weight              float32               `xml:"weight" validate:"required,max=10"`
	DeliverOn           ZasilkovnaDate        `xml:"deliver_on,omitempty"`
	Eshop               string                `xml:"eshop" validate:"required"`
	AdultContent        bool                  `xml:"adult_content,omitempty"`
	Note                string                `xml:"note,omitempty"`
	Street              string                `xml:"street" validate:"required"`
	HouseNumber         string                `xml:"house_number" validate:"required"`
	City                string                `xml:"city" validate:"required"`
	Province            string                `xml:"province,omitempty"`
	Zip                 string                `xml:"zip" validate:"required"`
	CarrierService      string                `xml:"carrier_service,omitempty"`
	CustomerBarcode     string                `xml:"customer_barcode,omitempty"`
	CarrierPickupPoint  string                `xml:"carrier_pickup_point" validate:"required"`
	CustomsDeclaration  []ItemCollection      `xml:"customs_declaration" validate:"required"`
	Size                Size                  `xml:"size" validate:"required"`
	AttributeCollection []AttributeCollection `xml:"attribute_collection" validate:"required"`
	Items               []ItemCollection      `xml:"items" validate:"required"`
}

func ValidatePacketAttributes() (isValidated bool, errorsArray []validator.FieldError) {
	fmt.Println("---PacketAttributes---")

	v := validator.New()

	size := Size{
		100, 30, 20,
	}
	sizeIsValidated, sizeErr := ValidateSize(size)
	fmt.Println(sizeIsValidated)
	fmt.Println(sizeErr)

	attribute := Attribute{
		Key:   "Some Key",
		Value: "Some Value",
	}
	attributeIsValidated, attributeErr := ValidateAttribute(attribute.Key, attribute.Value)
	fmt.Println(attributeIsValidated)
	fmt.Println(attributeErr)

	a := PacketAttributes{
		Number:             "Some number",
		Name:               "Some packet's name",
		Surname:            "Some packet's surname",
		Email:              "Some packet's email",
		Phone:              "Some packet's phone",
		AddressId:          10,
		Value:              20,
		Weight:             10,
		Eshop:              "Some e-shop's name",
		Street:             "Some street",
		HouseNumber:        "Some packet's house number",
		City:               "Some city",
		Zip:                "Some packet's zip",
		CarrierPickupPoint: "Some carrier's pickup point",
		CustomsDeclaration: []ItemCollection{
			{Item: Item{
				Attribute: attribute,
			}},
		},
		Size: size,
		AttributeCollection: []AttributeCollection{
			{Attribute: attribute},
		},
		Items: []ItemCollection{
			{Item: Item{
				Attribute: attribute,
			}},
		},
	}
	err := v.Struct(a)
	if err != nil { // If err contains errors, params are not validated
		isValidated = false
		for _, e := range err.(validator.ValidationErrors) {
			errorsArray = append(errorsArray, e)
		}
		return
	} else if !sizeIsValidated || !attributeIsValidated {
		isValidated = false
		return
	} else {
		isValidated = true
		return
	}
}
