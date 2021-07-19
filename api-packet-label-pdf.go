package main

import (
	"bytes"
	"encoding/binary"
	"encoding/xml"
	"go-zasilkovna/models"
	"io"
	"log"
	"net/http"
)

// PacketLabelPdf returns binary result in base64 encoding.
// Fetches a label for packet specified by packetId in format specified by format on a position specified
// by offset. The position is calculated left to right starting at top left corner of the document.
func (c Client) PacketLabelPdf(packetIds models.Ids, format string, offset int) (binary.ByteOrder, error) {
	packetLabelPdf := models.PacketsLabelsPdf{ApiPassword: c.credsProvider.ApiKey, PacketIds: packetIds, Format: format, Offset: offset}
	requestBody, marshalErr := xml.Marshal(packetLabelPdf)
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

	var binaryResult binary.ByteOrder

	unmarshalErr := xml.Unmarshal(body, &binaryResult)
	log.Println(string(body))
	if unmarshalErr != nil {
		return nil, unmarshalErr
	}

	return binaryResult, nil
}
