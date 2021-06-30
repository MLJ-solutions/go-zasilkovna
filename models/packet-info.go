package models

import (
	"encoding/xml"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)

type PacketInfoResponse struct {
	XMLName xml.Name         `xml:"response"`
	Status  string           `xml:"status"`
	Result  PacketInfoResult `xml:"result"`
}

func NewPacketInfoResponse(Status string, Result PacketInfoResult) *PacketInfoResponse {
	return &PacketInfoResponse{
		Status: Status,
		Result: Result,
	}
}

type PacketInfoResult struct {
	XMLName             xml.Name    `xml:"result"`
	BranchId            int         `xml:"branchId" validate:"required"`
	InvoicedWeightGrams int         `xml:"invoicedWeightGrams"`
	CourierInfo         CourierInfo `xml:"courierInfo"`
}

func NewPacketInfoResult(BranchId int, InvoicedWeightGrams int, CourierInfo CourierInfo) *PacketInfoResult {
	return &PacketInfoResult{
		BranchId:            BranchId,
		InvoicedWeightGrams: InvoicedWeightGrams,
		CourierInfo:         CourierInfo,
	}
}

type PacketInfo struct {
	XMLName     xml.Name `xml:"packetInfo"`
	ApiPassword string   `xml:"apiPassword" validate:"required"`
	PacketId    int      `xml:"packetId" validate:"required"`
}

func NewPacketInfo(ApiPassword string, PacketId int) *PacketInfo {
	return &PacketInfo{
		ApiPassword: ApiPassword,
		PacketId:    PacketId,
	}
}

func ValidatePacketInfo() (isValidated bool, errorsArray []validator.FieldError) {
	fmt.Println("---CurrentStatusRecord---")

	v := validator.New()
	a := PacketInfo{
		ApiPassword: "password",
		PacketId:    111,
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
