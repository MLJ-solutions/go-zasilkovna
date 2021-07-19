package models

import "encoding/xml"

type PacketTracking struct {
	XMLName     xml.Name `xml:"packetTracking"`
	ApiPassword string   `xml:"apiPassword" validate:"required"`
	PacketId    int      `xml:"packetId"`
}

type StatusRecords struct {
	StatusRecord []StatusRecord `xml:"statusRecords"`
}

type StatusRecord struct {
	XMLName              xml.Name           `xml:"statusRecord"`
	DateTime             ZasilkovnaDateTime `xml:"dateTime" validate:"required"`
	StatusCode           int                `xml:"statusCode" validate:"required"`
	CodeText             string             `xml:"codeText" validate:"required"`
	StatusText           string             `xml:"statusText" validate:"required"`
	BranchId             int                `xml:"branchId" validate:"required"`
	DestinationBranchId  int                `xml:"destinationBranchId" validate:"required"`
	ExternalTrackingCode string             `xml:"externalTrackingCode,omitempty"`
}

func NewStatusRecord(DateTime ZasilkovnaDateTime, StatusCode int, CodeText string, StatusText string,
	BranchId int, DestinationBranchId int, ExternalTrackingCode string) *StatusRecord {
	return &StatusRecord{
		DateTime:             DateTime,
		StatusCode:           StatusCode,
		CodeText:             CodeText,
		StatusText:           StatusText,
		BranchId:             BranchId,
		DestinationBranchId:  DestinationBranchId,
		ExternalTrackingCode: ExternalTrackingCode,
	}
}

func NewStatusRecordRequired(DateTime ZasilkovnaDateTime, StatusCode int, CodeText string, StatusText string,
	BranchId int, DestinationBranchId int) *StatusRecord {
	return &StatusRecord{
		DateTime:            DateTime,
		StatusCode:          StatusCode,
		CodeText:            CodeText,
		StatusText:          StatusText,
		BranchId:            BranchId,
		DestinationBranchId: DestinationBranchId,
	}
}
