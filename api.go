package go_zasilkovna

import (
	"encoding/xml"
	"github.com/MLJ-solutions/go-zasilkovna/models"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// List of success status.
var successStatus = []int{
	http.StatusOK,
	http.StatusCreated,
	http.StatusNoContent,
	http.StatusPartialContent,
}

func (c Client) executeMethod(method string, body io.Reader) (res *http.Response, err error) {
	// create request
	req, err := http.NewRequest(method, BasicUrl, body)
	req.Header.Set("Content-Type", "text/xml")
	if err != nil {
		errRequest := models.ToErrorResponse(err)
		return nil, errRequest
	}

	res, err = c.do(req)
	if err != nil {
		return nil, err
	}

	// For any known successful http status, return quickly.
	for _, httpStatus := range successStatus {
		if httpStatus == res.StatusCode {
			return res, nil
		}
	}

	//log.Println(res)
	all, err := ioutil.ReadAll(res.Body)
	closeResponse(res)
	if err != nil {
		return nil, err
	}

	var apiError models.PacketAttributesFault
	unmarshalErr := xml.Unmarshal(all, &apiError)
	if unmarshalErr != nil {
		log.Panic(string(all))
		return nil, unmarshalErr
	}

	log.Println(apiError)

	return nil, apiError
}

func (c Client) constructUrl() (*url.URL, error) {
	return url.Parse(c.EndpointURL().String())
}
