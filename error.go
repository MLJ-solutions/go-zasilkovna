package go_zasilkovna

import (
	"encoding/xml"
	"log"
)

type ToRfc7807Error interface {
	ToRfc7807Error(code int) error
}

func unmarshalToErrorRfc7807(body []byte, rError ToRfc7807Error) error {
	unmarshalErr := xml.Unmarshal(body, rError)
	log.Println(string(body))
	if unmarshalErr != nil {
		return unmarshalErr
	}

	return rError.ToRfc7807Error(400)
}
