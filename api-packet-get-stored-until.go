package main

import (
	"bytes"
	"encoding/xml"
	"go-zasilkovna/models"
	"io"
	"log"
	"net/http"
)

// PacketGetStoredUntil Fetches the date until which the packet specified by packetId is stored and waiting for pickup.
// If the packet is not yet ready for pickup or is already returning to sender null is returned.
func (c Client) PacketGetStoredUntil(packetId int) (models.ZasilkovnaDate, error) {
	packetGetStoredUntil := models.PacketGetStoredUntil{PacketId: packetId, ApiPassword: c.credsProvider.ApiKey}
	requestBody, marshalErr := xml.Marshal(packetGetStoredUntil)
	if marshalErr != nil {
		return models.ZasilkovnaDate{}, marshalErr
	}

	resp, err := c.executeMethod(http.MethodPost, bytes.NewBuffer(requestBody))
	if err != nil {
		return models.ZasilkovnaDate{}, err
	}

	log.Println(resp)
	body, err := io.ReadAll(resp.Body)
	defer closeResponse(resp)

	if err != nil {
		return models.ZasilkovnaDate{}, err
	}

	var zasilkovnaDate models.ZasilkovnaDate

	unmarshalErr := xml.Unmarshal(body, &zasilkovnaDate)
	log.Println(string(body))
	if unmarshalErr != nil {
		return models.ZasilkovnaDate{}, unmarshalErr
	}

	return zasilkovnaDate, nil
}
