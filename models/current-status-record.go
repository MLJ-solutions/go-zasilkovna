package models

import (
	"encoding/xml"
	"gopkg.in/go-playground/validator.v9"
	"time"
)

type PacketStatus struct {
	XMLName     xml.Name `xml:"packetStatus"`
	ApiPassword string   `xml:"apiPassword" validate:"required"`
	PacketId    int      `xml:"packetId" validate:"required"`
}

type CurrentStatusRecord struct {
	DateTime             ZasilkovnaDateTime `xml:"dateTime" validate:"required"`
	StatusCode           int                `xml:"statusCode" validate:"required"`
	CodeText             string             `xml:"codeText" validate:"required"`
	StatusText           string             `xml:"statusText" validate:"required"`
	BranchId             int                `xml:"branchId" validate:"required"`
	DestinationBranchId  int                `xml:"destinationBranchId" validate:"required"`
	ExternalTrackingCode string             `xml:"externalTrackingCode,omitempty"`
	IsReturning          *bool              `xml:"isReturning" validate:"required"` // *bool bcs of required
	StoredUntil          ZasilkovnaDate     `xml:"storedUntil" validate:"required"`
	CarrierId            int                `xml:"carrierId,omitempty"`
	CarrierName          string             `xml:"carrierName,omitempty"`
}

func NewPacketStatus(apiPassword string, packetId int) *PacketStatus {
	return &PacketStatus{ApiPassword: apiPassword, PacketId: packetId}
}

func NewCurrentStatusRecordRequired(DateTime ZasilkovnaDateTime, StatusCode int, CodeText string, StatusText string,
	BranchId int, DestinationBranchId int, IsReturning *bool, StoredUntil ZasilkovnaDate) *CurrentStatusRecord {
	return &CurrentStatusRecord{
		DateTime:            DateTime,
		StatusCode:          StatusCode,
		CodeText:            CodeText,
		StatusText:          StatusText,
		BranchId:            BranchId,
		DestinationBranchId: DestinationBranchId,
		IsReturning:         IsReturning,
		StoredUntil:         StoredUntil,
	}
}

func NewCurrentStatusRecord(DateTime ZasilkovnaDateTime, StatusCode int, CodeText string, StatusText string,
	BranchId int, DestinationBranchId int, ExternalTrackingCode string, IsReturning *bool, StoredUntil ZasilkovnaDate,
	CarrierId int, CarrierName string) *CurrentStatusRecord {
	return &CurrentStatusRecord{
		DateTime:             DateTime,
		StatusCode:           StatusCode,
		CodeText:             CodeText,
		StatusText:           StatusText,
		BranchId:             BranchId,
		DestinationBranchId:  DestinationBranchId,
		ExternalTrackingCode: ExternalTrackingCode,
		IsReturning:          IsReturning,
		StoredUntil:          StoredUntil,
		CarrierId:            CarrierId,
		CarrierName:          CarrierName,
	}
}

func ValidateCurrentStatusRecord() (isValidated bool, errorsArray []validator.FieldError) {
	v := validator.New()
	f := false
	a := CurrentStatusRecord{
		DateTime:            ZasilkovnaDateTime(time.Now()),
		StatusCode:          111,
		CodeText:            "Some code text",
		StatusText:          "Some status text",
		BranchId:            222,
		DestinationBranchId: 333,
		IsReturning:         &f, // Cannot set *bool type directly, because is pointer to memory
		StoredUntil:         ZasilkovnaDate(time.Now()),
	}
	err := v.Struct(a)
	if err != nil { // If err contains errors, params are not validated
		isValidated = false
		for _, e := range err.(validator.ValidationErrors) {
			errorsArray = append(errorsArray, e)
		}
		return
	} else {
		isValidated = true
		return
	}
}
