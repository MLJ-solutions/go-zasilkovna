package models

type CourierNumbers struct {
	CourierNumber string `xml:"courier_number_id" validate:"required"`
}

type CourierBarcodes struct {
	CourierBarcode string `xml:"courier_barcode_id" validate:"required"`
}

type CourierTrackingNumbers struct {
	CourierTrackingNumber string `xml:"courier_tracking_number" validate:"required"`
}
