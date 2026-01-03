package hueapi

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/snansidansi/hueapi/models"
)

type headerTransport struct {
	base    http.RoundTripper
	apiKey  string
	logging bool
}

func (t *headerTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	newReq := req.Clone(req.Context())
	newReq.Header.Add("hue-application-key", t.apiKey)

	if t.logging {
		fmt.Printf("[HUE] %s %s\n", newReq.Method, newReq.URL.String())
	}

	return t.base.RoundTrip(newReq)
}

type Client struct {
	Bridge     models.Bridge
	HTTPClient *http.Client

	Lights   *LightService
	Rooms    *RoomService
	Register *RegisterService
}

// Uses http.DefaultCLient when httpClient is nil.
func NewClient(bridge models.Bridge, apiKey string, httpClient *http.Client, logging bool) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}

	customTransport := http.DefaultTransport.(*http.Transport).Clone()
	customTransport.TLSClientConfig = &tls.Config{
		InsecureSkipVerify: true,
	}

	httpClient.Transport = &headerTransport{
		base:    customTransport,
		apiKey:  apiKey,
		logging: logging,
	}

	c := &Client{
		Bridge:     bridge,
		HTTPClient: httpClient,
	}

	c.Lights = &LightService{client: c}
	c.Rooms = &RoomService{client: c}
	c.Register = &RegisterService{client: c}

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

func (c *Client) CreateURL(specificURL string) string {
	return fmt.Sprintf("https://%s/clip/v2/%s", c.Bridge.IPAdress, specificURL)
}
