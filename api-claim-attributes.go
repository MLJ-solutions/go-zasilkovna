package go_zasilkovna

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"github.com/MLJ-solutions/go-zasilkovna/models"
	"io"
	"log"
	"net/http"
)

// CreatePacketClaim Creates a claim assistant packet from ClaimAttributes.
// On success returns PacketIdDetail with information about the newly created packet.
func (c Client) CreatePacketClaim(claimAttributes models.ClaimAttributes) (models.PacketIdDetail, error) {
	createPacketClaim := models.CreatePacketClaim{ApiPassword: c.credsProvider.ApiKey, ClaimAttributes: claimAttributes}
	requestBody, marshalErr := xml.Marshal(createPacketClaim)
	if marshalErr != nil {
		return models.PacketIdDetail{}, marshalErr
	}

	resp, err := c.executeMethod(http.MethodPost, bytes.NewBuffer(requestBody))
	if err != nil {
		return models.PacketIdDetail{}, err
	}

	log.Println(resp)
	body, err := io.ReadAll(resp.Body)
	defer closeResponse(resp)

	if err != nil {
		return models.PacketIdDetail{}, err
	}

	var packetIdDetail models.PacketIdDetail

	unmarshalErr := xml.Unmarshal(body, &packetIdDetail)
	log.Println(string(body))
	if unmarshalErr != nil {
		return models.PacketIdDetail{}, unmarshalErr
	}

	return packetIdDetail, nil
}

// ClaimAttributesValid Validates PacketAttributes.
// On success (the attributes are valid) returns <status>ok</status>.
func (c Client) ClaimAttributesValid(claimAttributes models.ClaimAttributes) (models.ErrorResponse, error) {
	claimAttributesValid := models.ClaimAttributesValid{ApiPassword: c.credsProvider.ApiKey, ClaimAttributes: claimAttributes}
	requestBody, marshalErr := xml.Marshal(claimAttributesValid)
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
	fmt.Print(string(body))
	if unmarshalErr != nil {
		return models.ErrorResponse{}, unmarshalErr
	}

	return validationErrors, nil
}
