package models

import "encoding/xml"

type CourierNumbers struct {
	XMLName       xml.Name `xml:"courierNumbers"`
	CourierNumber string   `xml:"courierNumber" validate:"required"`
}

type CourierBarcodes struct {
	XMLName        xml.Name `xml:"courierBarcodes"`
	CourierBarcode string   `xml:"courierBarcode" validate:"required"`
}

type CourierTrackingNumbers struct {
	XMLName               xml.Name `xml:"courierTrackingNumbers"`
	CourierTrackingNumber string   `xml:"courierTrackingNumber" validate:"required"`
}
