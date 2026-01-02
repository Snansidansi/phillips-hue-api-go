package hueapi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/snansidansi/hueapi/models"
)

type headerTransport struct {
	base   http.RoundTripper
	apiKey string
}

func (t *headerTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	newReq := req.Clone(req.Context())
	req.Header.Add("hue-application-key", t.apiKey)
	fmt.Printf("[HUE] %s %s\n", newReq.Method, newReq.URL.String())
	return t.base.RoundTrip(req)
}

type Client struct {
	IPAdress   string
	HTTPClient *http.Client

	Lights *LightService
}

// Uses http.DefaultCLient when httpClient is nil.
func NewClient(ipAdress, apiKey string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	httpClient.Transport = &headerTransport{
		base:   http.DefaultTransport,
		apiKey: apiKey,
	}

	c := &Client{
		IPAdress:   ipAdress,
		HTTPClient: httpClient,
	}

	c.Lights = &LightService{client: c}

	return c
}

// Uses http.DefaultCLient when httpClient is nil.
func DiscoverBridges(httpClient *http.Client) ([]models.Bridge, error) {
	const DiscoveryURL = "https://discovery.meethue.com/"

	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	resp, err := httpClient.Get(DiscoveryURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var foundBridges []models.Bridge
	if err := json.NewDecoder(resp.Body).Decode(&foundBridges); err != nil {
		return nil, err
	}

	return foundBridges, nil
}
