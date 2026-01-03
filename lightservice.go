package hueapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/snansidansi/hueapi/models"
)

type LightService struct {
	client *Client
}

func (l *LightService) GetAllLights() (*models.HueResponse[models.Light], error) {
	url := l.client.CreateURL("resource/light")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := l.client.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var hueResp models.HueResponse[models.Light]
	hueResp.StatusCode = resp.StatusCode

	if err := json.NewDecoder(resp.Body).Decode(&hueResp); err != nil {
		return &hueResp, fmt.Errorf("decoding failed (status %d): %w", resp.StatusCode, err)
	}

	return &hueResp, nil
}

func (l *LightService) GetLightByID(id string) (*models.HueResponse[models.Light], error) {
	url := l.client.CreateURL("resource/light/" + id)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := l.client.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var hueResp models.HueResponse[models.Light]
	hueResp.StatusCode = resp.StatusCode

	if err := json.NewDecoder(resp.Body).Decode(&hueResp); err != nil {
		return &hueResp, fmt.Errorf("decoding failed (status %d): %w", resp.StatusCode, err)
	}

	return &hueResp, nil
}

func (s *LightService) SetLightState(id string, update models.LightPut) (*models.HueActionResponse, error) {
	url := s.client.CreateURL("resource/light/" + id)

	jsonData, err := json.Marshal(update)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := s.client.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var hueResp models.HueActionResponse
	hueResp.StatusCode = resp.StatusCode

	if err := json.NewDecoder(resp.Body).Decode(&hueResp); err != nil {
		return &hueResp, fmt.Errorf("response decode failed (status %d): %w", resp.StatusCode, err)
	}

	return &hueResp, nil
}
