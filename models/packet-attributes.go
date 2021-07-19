package models

import (
	"encoding/xml"
	"gopkg.in/go-playground/validator.v9"
	"log"
)

type CreatePacket struct {
	XMLName          xml.Name         `xml:"createPacket"`
	ApiPassword      string           `xml:"apiPassword" validate:"required"`
	PacketAttributes PacketAttributes `xml:"packetAttributes" validate:"required"`
}

type PacketAttributes struct {
	Id                  int                 `xml:"id,omitempty"`
	Number              string              `xml:"number" validate:"required"`
	Name                string              `xml:"name" validate:"required"`
	Surname             string              `xml:"surname" validate:"required"`
	Company             string              `xml:"company,omitempty"`
	Email               string              `xml:"email" validate:"required"` //Required email or phone
	Phone               string              `xml:"phone" validate:"required"` //Required email or phone
	AddressId           int                 `xml:"addressId" validate:"required"`
	Currency            string              `xml:"currency,omitempty"`
	Cod                 float32             `xml:"cod,omitempty"`
	Value               float32             `xml:"value" validate:"required"`
	Weight              float32             `xml:"weight" validate:"max=10"`
	DeliverOn           ZasilkovnaDate      `xml:"deliverOn,omitempty"`
	Eshop               string              `xml:"eshop" validate:"required"`
	AdultContent        bool                `xml:"adultContent,omitempty"`
	Note                string              `xml:"note,omitempty"`
	Street              string              `xml:"street"`
	HouseNumber         string              `xml:"houseNumber"`
	City                string              `xml:"city"`
	Province            string              `xml:"province,omitempty"`
	Zip                 string              `xml:"zip"`
	CarrierService      string              `xml:"carrierService,omitempty"`
	CustomerBarcode     string              `xml:"customerBarcode,omitempty"`
	CarrierPickupPoint  string              `xml:"carrierPickupPoint"`  // Required for some carriers
	CustomsDeclaration  ItemCollection      `xml:"customsDeclaration"`  // Required on address delivery outside EU
	Size                Size                `xml:"size"`                // Required for some carriers
	AttributeCollection AttributeCollection `xml:"attributeCollection"` // Required for some carriers
	Items               ItemCollection      `xml:"items"`               // Required for some carriers
}

type PacketIdDetail struct {
	XMLName     xml.Name `xml:"packetIdDetail"`
	Id          int      `xml:"id"`
	Barcode     string   `xml:"barcode"`
	BarcodeText string   `xml:"barcodeText"`
}

type PacketDetail struct {
	XMLName     xml.Name `xml:"packetIdDetail"`
	Id          int      `xml:"id"`
	Barcode     string   `xml:"barcode"`
	BarcodeText string   `xml:"barcodeText"`
	Password    string   `xml:"password"`
}

type PacketAttributesValid struct {
	XMLName          xml.Name         `xml:"packetAttributesValid"`
	ApiPassword      string           `xml:"apiPassword" validate:"required"`
	PacketAttributes PacketAttributes `xml:"packetAttributes" validate:"required"`
}

func NewPacketAttributesValid(ApiPassword string, PacketAttributes PacketAttributes) *PacketAttributesValid {
	return &PacketAttributesValid{
		ApiPassword:      ApiPassword,
		PacketAttributes: PacketAttributes,
	}
}

func NewPacketIdDetail(Id int, Barcode string, BarcodeText string) *PacketIdDetail {
	return &PacketIdDetail{
		Id:          Id,
		Barcode:     Barcode,
		BarcodeText: BarcodeText,
	}
}

func NewCreatePacket(ApiPassword string, PacketAttributes PacketAttributes) *CreatePacket {
	return &CreatePacket{ApiPassword: ApiPassword, PacketAttributes: PacketAttributes}
}

func NewPacketAttributesRequired(Number string, Name string, Surname string, Email string, Phone string, AddressId int,
	Value float32, Weight float32, Eshop string, Street string, HouseNumber string, City string,
	Zip string, CarrierPickupPoint string,
	CustomsDeclaration ItemCollection, Size Size, AttributeCollection AttributeCollection,
	Items ItemCollection) *PacketAttributes {
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
	CustomsDeclaration ItemCollection, Size Size, AttributeCollection AttributeCollection,
	Items ItemCollection) *PacketAttributes {
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
	v := validator.New()

	size := Size{
		100, 30, 20,
	}
	sizeIsValidated, sizeErr := ValidateSize(size)
	log.Println(sizeIsValidated)
	log.Println(sizeErr)

	attribute := Attribute{
		Key:   "Some Key",
		Value: "Some Value",
	}
	attributeIsValidated, attributeErr := ValidateAttribute(attribute.Key, attribute.Value)
	log.Println(attributeIsValidated)
	log.Println(attributeErr)

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
