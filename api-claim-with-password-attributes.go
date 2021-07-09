package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"go-zasilkovna/models"
	"io"
	"net/http"
)

// CreatePacketClaimWithPassword Creates a claim assistant packet from ClaimWithPasswordAttributes.
// On success returns PacketDetail with information about the newly created packet.
func (c Client) CreatePacketClaimWithPassword(createPacketClaimWithPassword models.CreatePacketClaimWithPassword) (models.PacketDetail, error) {
	requestBody, marshalErr := xml.Marshal(createPacketClaimWithPassword)
	if marshalErr != nil {
		return models.PacketDetail{}, marshalErr
	}

	resp, err := c.executeMethod(http.MethodPost, bytes.NewBuffer(requestBody))
	if err != nil {
		return models.PacketDetail{}, err
	}

	fmt.Println(resp)
	body, err := io.ReadAll(resp.Body)
	defer closeResponse(resp)

	if err != nil {
		return models.PacketDetail{}, err
	}

	var packetDetail models.PacketDetail

	unmarshalErr := xml.Unmarshal(body, &packetDetail)
	fmt.Print(string(body))
	if unmarshalErr != nil {
		return models.PacketDetail{}, unmarshalErr
	}

	return packetDetail, nil
}
