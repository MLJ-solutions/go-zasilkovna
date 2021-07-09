package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"go-zasilkovna/models"
	"io"
	"net/http"
)

// PacketTracking Fetches the whole tracking history of the packet specified by packetId.
// On success returns StatusRecords struct
func (c Client) PacketTracking(packetId int) (models.StatusRecords, error) {
	packetTracking := models.PacketTracking{PacketId: packetId, ApiPassword: c.credsProvider.ApiKey}
	requestBody, marshalErr := xml.Marshal(packetTracking)
	if marshalErr != nil {
		return models.StatusRecords{}, marshalErr
	}

	resp, err := c.executeMethod(http.MethodPost, bytes.NewBuffer(requestBody))
	if err != nil {
		return models.StatusRecords{}, err
	}

	fmt.Println(resp)
	body, err := io.ReadAll(resp.Body)
	defer closeResponse(resp)

	if err != nil {
		return models.StatusRecords{}, err
	}

	var statusRecords models.StatusRecords

	unmarshalErr := xml.Unmarshal(body, &statusRecords)
	fmt.Print(string(body))
	if unmarshalErr != nil {
		return models.StatusRecords{}, unmarshalErr
	}

	return statusRecords, nil
}
