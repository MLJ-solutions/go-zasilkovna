package go_zasilkovna

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

const BasicUrl = "https://www.zasilkovna.cz/api/rest/"

type Client struct {
	// Parsed accountEndpoint url provided by the user.
	endpointURL *url.URL

	// Holds various credential providers.
	credsProvider *Credentials
	httpClient    *http.Client
}

type Options struct {
	Transport http.RoundTripper
	Creds     *Credentials
}

type Credentials struct {
	sync.Mutex

	ApiKey string
}

// New - instantiate minio client with options
func New(opts *Options) (*Client, error) {
	endpoint := BasicUrl
	if opts == nil {
		return nil, errors.New("no options provided")
	}
	clnt, err := privateNew(endpoint, opts)
	if err != nil {
		return nil, err
	}

	return clnt, nil
}

func NewCredentials(key string) *Credentials {
	return &Credentials{
		ApiKey: key,
	}
}

// EndpointURL returns the URL of the zasilkovna.
func (c *Client) EndpointURL() *url.URL {
	endpoint := *c.endpointURL // copy to prevent callers from modifying internal state
	return &endpoint
}

// getEndpointURL - construct a new accountEndpoint.
func getEndpointURL(endpoint string) (*url.URL, error) {
	endpointURLStr := fmt.Sprintf("%s/", strings.Trim(endpoint, "/"))
	endpointURL, err := url.Parse(endpointURLStr)
	if err != nil {
		return nil, err
	}

	// Validate incoming accountEndpoint URL.
	return endpointURL, nil
}

// Redirect requests by re signing the request.
func (c *Client) redirectHeaders(req *http.Request, via []*http.Request) error {
	if len(via) >= 5 {
		return errors.New("stopped after 5 redirects")
	}
	if len(via) == 0 {
		return nil
	}

	*c.endpointURL = *req.URL

	return nil
}

func privateNew(endpoint string, opts *Options) (*Client, error) {
	// construct accountEndpoint.
	endpointURL, err := getEndpointURL(endpoint)
	if err != nil {
		return nil, err
	}

	// instantiate new Client.
	clnt := new(Client)

	// Save the credentials.
	clnt.credsProvider = opts.Creds
	// Save accountEndpoint URL, user agent for future uses.
	clnt.endpointURL = endpointURL

	transport := opts.Transport
	if transport == nil {
		transport, err = DefaultTransport()
		if err != nil {
			return nil, err
		}
	}

	// Instantiate http client and bucket location cache.
	clnt.httpClient = &http.Client{
		Transport:     transport,
		CheckRedirect: clnt.redirectHeaders,
	}

	// Return.
	return clnt, nil
}

func (c Client) do(req *http.Request) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		// Handle this specifically for now until future Golang versions fix this issue properly.
		if urlErr, ok := err.(*url.Error); ok {
			if strings.Contains(urlErr.Err.Error(), "EOF") {
				return nil, &url.Error{
					Op:  urlErr.Op,
					URL: urlErr.URL,
					Err: errors.New("Connection closed by foreign host " + urlErr.URL + ". Retry again."),
				}
			}
		}
		return nil, err
	}

	// Response cannot be non-nil, report error if thats the case.
	if resp == nil {
		return nil, errors.New("Response is empty.")
	}

	return resp, nil
}
