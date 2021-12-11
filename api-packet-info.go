package go_zasilkovna

import (
	"bytes"
	"encoding/xml"
	"github.com/MLJ-solutions/go-zasilkovna/models"
	"io"
	"net/http"
)

// PacketInfo Returns additional information about packet and its consignment to an external courier, if there is one.
// On success it returns PacketInfoResult.
func (c Client) PacketInfo(packetId int) (models.PacketInfoResponse, error) {
	packetInfo := models.PacketInfo{PacketId: packetId, ApiPassword: c.credsProvider.ApiKey}
	requestBody, marshalErr := xml.Marshal(packetInfo)
	if marshalErr != nil {
		return models.PacketInfoResponse{}, marshalErr
	}
	resp, err := c.executeMethod(http.MethodPost, bytes.NewBuffer(requestBody))
	if err != nil {
		return models.PacketInfoResponse{}, err
	}

	body, err := io.ReadAll(resp.Body)
	defer closeResponse(resp)

	if err != nil {
		return models.PacketInfoResponse{}, err
	}

	var response models.PacketInfoResponse

	unmarshalErr := xml.Unmarshal(body, &response)
	if unmarshalErr != nil {
		return models.PacketInfoResponse{}, unmarshalErr
	}

	return response, nil
}
