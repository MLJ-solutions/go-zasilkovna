package models

import (
	"bytes"
	"encoding/xml"
	"fmt"
	go_rfc7807 "github.com/MLJ-solutions/go-rfc7807"
)

type PacketAttributesFault struct {
	XMLName xml.Name `xml:"response"`
	Status  string   `xml:"status"`
	Fault   string   `xml:"fault"`
	String  string   `xml:"string"`
	Detail  Detail   `xml:"detail"`
}

type Detail struct {
	Attributes ErrorAttributes `xml:"attributes"`
}

type ErrorAttributes struct {
	Fault []Fault `xml:"fault"`
}

type Fault struct {
	Name  string `xml:"name"`
	Fault string `xml:"fault"`
}

func (e *PacketAttributesFault) ToRfc7807Error(code int) error {
	result := &go_rfc7807.Rfc7807Error{
		Code:  code,
		Type:  e.Fault,
		Title: e.String,
	}
	for _, detail := range e.Detail.Attributes.Fault {
		_ = result.PutParam(detail.Name, detail.Fault)
	}

	return result
}

func ToErrorResponse(err error) PacketAttributesFault {
	switch err := err.(type) {
	case PacketAttributesFault:
		return err
	default:
		return PacketAttributesFault{}
	}
}

// Error - Returns error string.
func (e PacketAttributesFault) Error() string {
	b := new(bytes.Buffer)
	for _, value := range e.Detail.Attributes.Fault {
		//goland:noinspection GoUnhandledErrorResult
		fmt.Fprintf(b, "%s=\"%s\"\n", value.Name, value.Fault)
	}

	return b.String()
}
