package go_zasilkovna

import (
	"bytes"
	"encoding/binary"
	"encoding/xml"
	"github.com/MLJ-solutions/go-zasilkovna/models"
	"io"
	"log"
	"net/http"
)

// BarcodePng returns binary result in base64 encoding.
// The barcode is created with Code 128 symbology.
// NOTE: The method does not validate the barcode in any way. If you wish to create a barcode to use it on
// your labels, it is important that you use packetId prefixed by the letter Z e.g. Z1234567890.
func (c Client) BarcodePng(barcode string) (binary.ByteOrder, error) {
	barcodePng := models.BarcodePng{Barcode: barcode, ApiPassword: c.credsProvider.ApiKey}
	requestBody, marshalErr := xml.Marshal(barcodePng)
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
