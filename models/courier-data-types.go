package models

type CourierNumbers struct {
	CourierNumber string `json:"courier_number_id"`
}

type CourierBarcodes struct {
	CourierBarcode string `json:"courier_barcode_id"`
}

type CourierTrackingNumbers struct {
	CourierTrackingNumber string `json:"courier_tracking_number"`
}
