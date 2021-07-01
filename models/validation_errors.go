package models

import "encoding/xml"

type ErrorResponse struct {
	XMLName xml.Name `xml:"response"`
	Status  string   `xml:"status"`
	Fault   string   `xml:"fault"`
	Detail  Detail   `xml:"detail"`
}

type Detail struct {
	XMLName    xml.Name        `xml:"detail"`
	Attributes ErrorAttributes `xml:"attributes"`
}

type ErrorAttributes struct {
	XMLName xml.Name `xml:"attributes"`
	Fault   []Fault  `xml:"fault"`
}

type Fault struct {
	XMLName xml.Name `xml:"fault"`
	Name    string   `xml:"name"`
	Fault   string   `xml:"fault"`
}
