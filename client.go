package hueapi

import (
	"encoding/json"
	"net/http"

	"github.com/snansidansi/hueapi/models"
)

type Client struct {
	IPAdress   string
	APIKey     string
	HTTPClient *http.Client

	Lights *LightService
}

// Uses http.DefaultCLient when httpClient is nil.
func NewClient(ipAdress, apiKey string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	c := &Client{
		IPAdress:   ipAdress,
		APIKey:     apiKey,
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
