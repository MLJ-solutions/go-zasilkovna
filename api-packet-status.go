package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"go-zasilkovna/models"
	"io"
	"net/http"
)

// PacketStatus fetches information about the current status of the packet specified by packetId.
// On success returns CurrentStatusRecord.
func (c Client) PacketStatus(packetId int) (models.CurrentStatusRecord, error) {
	packetStatus := models.PacketStatus{PacketId: packetId, ApiPassword: c.credsProvider.ApiKey}
	requestBody, marshalErr := xml.Marshal(packetStatus)
	if marshalErr != nil {
		return models.CurrentStatusRecord{}, marshalErr
	}

	fmt.Println(string(requestBody))

	resp, err := c.executeMethod(http.MethodPost, bytes.NewBuffer(requestBody))
	if err != nil {
		return models.CurrentStatusRecord{}, err
	}

	fmt.Println(resp)
	body, err := io.ReadAll(resp.Body)
	defer closeResponse(resp)

	if err != nil {
		return models.CurrentStatusRecord{}, err
	}

	var currentStatusRecord models.CurrentStatusRecord

	unmarshalErr := xml.Unmarshal(body, &currentStatusRecord)
	fmt.Print(string(body))
	if unmarshalErr != nil {
		return models.CurrentStatusRecord{}, unmarshalErr
	}

	return currentStatusRecord, nil
}
