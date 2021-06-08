package models

type CurrentStatusRecord struct {
	DateTime             ZasilkovnaDateTime `json:"date_time"`
	StatusCode           int                `json:"status_code"`
	CodeText             string             `json:"code_text"`
	StatusText           string             `json:"status_text"`
	BranchId             int                `json:"branch_id"`
	DestinationBranchId  int                `json:"destination_branch_id"`
	ExternalTrackingCode string             `json:"external_tracking_code,omitempty"`
	IsReturning          bool               `json:"is_returning"`
	StoredUntil          ZasilkovnaDate     `json:"stored_until"`
	CarrierId            int                `json:"carrier_id,omitempty"`
	CarrierName          string             `json:"carrier_name,omitempty"`
}
