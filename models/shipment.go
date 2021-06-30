package models

type ShipmentIdDetailParams struct {
	Id          int    `xml:"apiPassword" validate:"required"`
	Barcode     string `xml:"barcode" validate:"required"`
}

func NewShipmentIdDetailParams(Id int, Barcode string) *ShipmentIdDetailParams {
	return &ShipmentIdDetailParams{
		Id:          Id,
		Barcode:     Barcode,
	}
}

type ShipmentIdDetail struct {
	Id          int    `xml:"apiPassword" validate:"required"`
	Checksum    string `xml:"createPacketClaimWithPassword" validate:"required"`
	Barcode     string `xml:"barcode" validate:"required"`
	BarcodeText string `xml:"barcodeText" validate:"required"`
}

func NewShipmentIdDetail(Id int, Checksum string, Barcode string, BarcodeText string) *ShipmentIdDetail {
	return &ShipmentIdDetail{
		Id:          Id,
		Checksum:    Checksum,
		Barcode:     Barcode,
		BarcodeText: BarcodeText,
	}
}
