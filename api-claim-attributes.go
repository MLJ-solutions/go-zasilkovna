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
func (c Client) CreatePacketClaim(claimAttributes models.ClaimAttributes) (*models.PacketIdDetail, error) {
	createPacketClaim := models.CreatePacketClaim{ApiPassword: c.credsProvider.ApiKey, ClaimAttributes: claimAttributes}
	requestBody, marshalErr := xml.Marshal(createPacketClaim)
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

	packetIdDetail := &models.PacketIdDetail{}

	unmarshalErr := xml.Unmarshal(body, packetIdDetail)
	log.Println(string(body))
	if unmarshalErr != nil {
		return nil, unmarshalErr
	}

	return packetIdDetail, nil
}

// ClaimAttributesValid Validates PacketAttributes.
// On success (the attributes are valid) returns <status>ok</status> `err.Status == ResponseStatusOk`.
// On error (the attributes are NOT valid) returns <status>fault</status> as Rfc7807Error `err.Status == ResponseStatusFault`.
func (c Client) ClaimAttributesValid(claimAttributes models.ClaimAttributes) error {
	claimAttributesValid := models.ClaimAttributesValid{ApiPassword: c.credsProvider.ApiKey, ClaimAttributes: claimAttributes}
	requestBody, marshalErr := xml.Marshal(claimAttributesValid)
	if marshalErr != nil {
		return marshalErr
	}

	resp, err := c.executeMethod(http.MethodPost, bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}

	fmt.Println(resp)
	body, err := io.ReadAll(resp.Body)
	defer closeResponse(resp)

	if err != nil {
		return err
	}

	validationErrors := &models.PacketAttributesFault{}

	unmarshalErr := xml.Unmarshal(body, validationErrors)
	fmt.Print(string(body))
	if unmarshalErr != nil {
		return unmarshalErr
	}

	return validationErrors.ToRfc7807Error(200)
}
