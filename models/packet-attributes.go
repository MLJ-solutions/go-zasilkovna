package models

import (
	"encoding/xml"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)

type CreatePacket struct {
	XMLName          xml.Name         `xml:"createPacket"`
	ApiPassword      string           `xml:"apiPassword" validate:"required"`
	PacketAttributes PacketAttributes `xml:"packetAttributes" validate:"required"`
}

type PacketAttributes struct {
	Id                  int                   `xml:"id,omitempty"`
	Number              string                `xml:"number" validate:"required"`
	Name                string                `xml:"name" validate:"required"`
	Surname             string                `xml:"surname" validate:"required"`
	Company             string                `xml:"company,omitempty"`
	Email               string                `xml:"email" validate:"required"`
	Phone               string                `xml:"phone" validate:"required"`
	AddressId           int                   `xml:"addressId" validate:"required"`
	Currency            string                `xml:"currency,omitempty"`
	Cod                 float32               `xml:"cod,omitempty"`
	Value               float32               `xml:"value" validate:"required"`
	Weight              float32               `xml:"weight" validate:"required,max=10"`
	DeliverOn           ZasilkovnaDate        `xml:"deliverOn,omitempty"`
	Eshop               string                `xml:"eshop" validate:"required"`
	AdultContent        bool                  `xml:"adultContent,omitempty"`
	Note                string                `xml:"note,omitempty"`
	Street              string                `xml:"street" validate:"required"`
	HouseNumber         string                `xml:"houseNumber" validate:"required"`
	City                string                `xml:"city" validate:"required"`
	Province            string                `xml:"province,omitempty"`
	Zip                 string                `xml:"zip" validate:"required"`
	CarrierService      string                `xml:"carrierService,omitempty"`
	CustomerBarcode     string                `xml:"customerBarcode,omitempty"`
	CarrierPickupPoint  string                `xml:"carrierPickupPoint" validate:"required"`
	CustomsDeclaration  []ItemCollection      `xml:"customsDeclaration" validate:"required"`
	Size                Size                  `xml:"size" validate:"required"`
	AttributeCollection []AttributeCollection `xml:"attributeCollection" validate:"required"`
	Items               []ItemCollection      `xml:"items" validate:"required"`
}

func NewCreatePacket(ApiPassword string, PacketAttributes PacketAttributes) *CreatePacket {
	return &CreatePacket{ApiPassword: ApiPassword, PacketAttributes: PacketAttributes}
}

func NewPacketAttributesRequired(Number string, Name string, Surname string, Email string, Phone string, AddressId int,
	Value float32, Weight float32, Eshop string, Street string, HouseNumber string, City string,
	Zip string, CarrierPickupPoint string,
	CustomsDeclaration []ItemCollection, Size Size, AttributeCollection []AttributeCollection,
	Items []ItemCollection) *PacketAttributes {
	return &PacketAttributes{
		Number:              Number,
		Name:                Name,
		Surname:             Surname,
		Email:               Email,
		Phone:               Phone,
		AddressId:           AddressId,
		Value:               Value,
		Weight:              Weight,
		Eshop:               Eshop,
		Street:              Street,
		HouseNumber:         HouseNumber,
		City:                City,
		Zip:                 Zip,
		CarrierPickupPoint:  CarrierPickupPoint,
		CustomsDeclaration:  CustomsDeclaration,
		Size:                Size,
		AttributeCollection: AttributeCollection,
		Items:               Items,
	}
}

func NewPacketAttributes(Id int, Number string, Name string, Surname string,
	Company string, Email string, Phone string, AddressId int, Currency string, Cod float32,
	Value float32, Weight float32, DeliverOn ZasilkovnaDate, Eshop string,
	AdultContent bool, Note string, Street string, HouseNumber string, City string, Province string,
	Zip string, CarrierService string, CustomerBarcode string, CarrierPickupPoint string,
	CustomsDeclaration []ItemCollection, Size Size, AttributeCollection []AttributeCollection,
	Items []ItemCollection) *PacketAttributes {
	return &PacketAttributes{
		Id:                  Id,
		Number:              Number,
		Name:                Name,
		Surname:             Surname,
		Company:             Company,
		Email:               Email,
		Phone:               Phone,
		AddressId:           AddressId,
		Currency:            Currency,
		Cod:                 Cod,
		Value:               Value,
		Weight:              Weight,
		DeliverOn:           DeliverOn,
		Eshop:               Eshop,
		AdultContent:        AdultContent,
		Note:                Note,
		Street:              Street,
		HouseNumber:         HouseNumber,
		City:                City,
		Province:            Province,
		Zip:                 Zip,
		CarrierService:      CarrierService,
		CustomerBarcode:     CustomerBarcode,
		CarrierPickupPoint:  CarrierPickupPoint,
		CustomsDeclaration:  CustomsDeclaration,
		Size:                Size,
		AttributeCollection: AttributeCollection,
		Items:               Items,
	}
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
