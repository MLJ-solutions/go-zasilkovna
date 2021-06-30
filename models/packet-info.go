package models

import (
	"encoding/xml"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)

type Response struct {
	XMLName xml.Name `xml:"response"`
	Status  string   `xml:"status"`
	Result  Result   `xml:"result"`
}

type Result struct {
	XMLName             xml.Name    `xml:"result"`
	BranchId            int         `xml:"branchId" validate:"required"`
	InvoicedWeightGrams int         `xml:"invoicedWeightGrams"`
	CourierInfo         CourierInfo `xml:"courierInfo"`
}

func NewResult(BranchId int, InvoicedWeightGrams int, CourierInfo CourierInfo) *Result {
	return &Result{
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
