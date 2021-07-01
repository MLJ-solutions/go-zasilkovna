package models

import "encoding/xml"

type PacketGetStoredUntil struct {
	XMLName     xml.Name `xml:"packetGetStoredUntil"`
	ApiPassword string   `xml:"apiPassword" validate:"required"`
	PacketId    int      `xml:"packetId" validate:"required"`
}

func NewPacketGetStoredUntil(apiPassword string, packetId int) *PacketGetStoredUntil {
	return &PacketGetStoredUntil{ApiPassword: apiPassword, PacketId: packetId}
}
