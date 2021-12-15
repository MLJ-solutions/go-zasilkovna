package go_zasilkovna

import (
	"bytes"
	"encoding/xml"
	"github.com/MLJ-solutions/go-zasilkovna/models"
	"io"
	"log"
	"net/http"
)

// CreatePacket Creates packet from PacketAttributes.
// On success returns PacketIdDetail with information about the newly created packet.
func (c Client) CreatePacket(packetAttributes models.PacketAttributes) (*models.CreatePacketResponse, error) {
	createPacket := models.CreatePacket{ApiPassword: c.credsProvider.ApiKey, PacketAttributes: packetAttributes}
	requestBody, marshalErr := xml.Marshal(createPacket)
	if marshalErr != nil {
		return nil, marshalErr
	}

	resp, err := c.executeMethod(http.MethodPost, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}

	log.Println(resp)
	body, err := io.ReadAll(resp.Body)
	defer closeResponse(resp)

	if err != nil {
		return nil, err
	}

	packetIdDetail := &models.CreatePacketResponse{}

	unmarshalErr := xml.Unmarshal(body, packetIdDetail)
	log.Println(string(body))
	if unmarshalErr != nil {
		return nil, unmarshalErr
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

	log.Println(resp)
	body, err := io.ReadAll(resp.Body)
	defer closeResponse(resp)

	if err != nil {
		return models.ErrorResponse{}, err
	}

	var validationErrors models.ErrorResponse

	unmarshalErr := xml.Unmarshal(body, &validationErrors)
	log.Println(string(body))
	if unmarshalErr != nil {
		return models.ErrorResponse{}, unmarshalErr
	}

	return validationErrors, nil
}
