package models

import "encoding/xml"

type BarcodePng struct {
	XMLName     xml.Name `xml:"barcodePng"`
	ApiPassword string   `xml:"apiPassword" validate:"required"`
	Barcode     string   `xml:"barcode" validate:"required"`
}

func NewBarcodePng(ApiPassword string, Barcode string) *BarcodePng {
	return &BarcodePng{ApiPassword: ApiPassword, Barcode: Barcode}
}
