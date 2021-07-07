package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"go-zasilkovna/models"
	"io"
	"net/http"
)

// CreatePacket Creates packet from PacketAttributes.
// On success returns PacketIdDetail with information about the newly created packet.
func (c Client) CreatePacket(packetAttributes models.PacketAttributes) (models.PacketIdDetail, error) {
	createPacket := models.CreatePacket{ApiPassword: c.credsProvider.ApiKey, PacketAttributes: packetAttributes}
	requestBody, marshalErr := xml.Marshal(createPacket)
	if marshalErr != nil {
		return models.PacketIdDetail{}, marshalErr
	}

	resp, err := c.executeMethod(http.MethodPost, bytes.NewBuffer(requestBody))
	if err != nil {
		return models.PacketIdDetail{}, err
	}

	fmt.Println(resp)
	body, err := io.ReadAll(resp.Body)
	defer closeResponse(resp)

	if err != nil {
		return models.PacketIdDetail{}, err
	}

	var packetIdDetail models.PacketIdDetail

	unmarshalErr := xml.Unmarshal(body, &packetIdDetail)
	fmt.Print(string(body))
	if unmarshalErr != nil {
		return models.PacketIdDetail{}, unmarshalErr
	}

	return packetIdDetail, nil
}

// PacketAttributesValid Validates PacketAttributes.
// On success (the attributes are valid) returns <status>ok</status>.
func (c Client) PacketAttributesValid(packetAttributes models.PacketAttributes) (models.ErrorResponse, error) {
	packetAttributesValid := models.PacketAttributesValid{ApiPassword: c.credsProvider.ApiKey, PacketAttributes: packetAttributes}
	requestBody, marshalErr := xml.Marshal(packetAttributesValid)
	if marshalErr != nil {
		return models.ErrorResponse{}, marshalErr
	}

	resp, err := c.executeMethod(http.MethodPost, bytes.NewBuffer(requestBody))
	if err != nil {
		return models.ErrorResponse{}, err
	}

	fmt.Println(resp)
	body, err := io.ReadAll(resp.Body)
	defer closeResponse(resp)

	if err != nil {
		return models.ErrorResponse{}, err
	}

	var validationErrors models.ErrorResponse

	unmarshalErr := xml.Unmarshal(body, &validationErrors)
	fmt.Print("BODY")
	fmt.Print(string(body))
	if unmarshalErr != nil {
		return models.ErrorResponse{}, unmarshalErr
	}

	return validationErrors, nil
}
