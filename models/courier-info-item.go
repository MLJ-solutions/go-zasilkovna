package models

type CourierInfoItem struct {
	CourierId              int                      `json:"courier_id"`
	CourierName            string                   `json:"courier_name"`
	CourierNumbers         []CourierNumbers         `json:"courier_numbers,omitempty"`
	CourierBarcodes        []CourierBarcodes        `json:"courier_barcodes,omitempty"`
	CourierTrackingNumbers []CourierTrackingNumbers `json:"courier_tracking_numbers,omitempty"`
	CourierTrackingUrls    []CourierTrackingUrl     `json:"courier_tracking_urls,omitempty"`
}
