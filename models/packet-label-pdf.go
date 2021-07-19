package models

import "encoding/xml"

type PacketsLabelsPdf struct {
	XMLName     xml.Name `xml:"packetsLabelsPdf"`
	ApiPassword string   `xml:"apiPassword" validate:"required"`
	PacketIds   Ids      `xml:"packetIds" validate:"required"`
	Format      string   `xml:"format" validate:"required"`
	Offset      int      `xml:"offset" validate:"required"`
}

type Ids struct {
	PacketIds []int `xml:"id" validate:"required"`
}

func NewPacketsLabelsPdf(ApiPassword string, PacketIds Ids, Format string, Offset int) *PacketsLabelsPdf {
	return &PacketsLabelsPdf{
		ApiPassword: ApiPassword,
		PacketIds:   PacketIds,
		Format:      Format,
		Offset:      Offset,
	}
}

type PacketLabelPdf struct {
	XMLName     xml.Name `xml:"packetLabelPdf"`
	ApiPassword string   `xml:"apiPassword" validate:"required"`
	PacketId    int      `xml:"packetId" validate:"required"`
	Format      string   `xml:"format" validate:"required"`
	Offset      int      `xml:"offset" validate:"required"`
}

func NewPacketLabelPdf(ApiPassword string, PacketId int, Format string, Offset int) *PacketLabelPdf {
	return &PacketLabelPdf{
		ApiPassword: ApiPassword,
		PacketId:    PacketId,
		Format:      Format,
		Offset:      Offset,
	}
}
