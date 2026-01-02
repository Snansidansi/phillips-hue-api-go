package hueapi

import (
	"encoding/json"

	"github.com/snansidansi/hueapi/models"
)

type LightService struct {
	client *Client
}

func (l *LightService) GetAllLights() ([]models.Light, []models.HueError, error) {
	url := l.client.CreateURL("resource/light")
	resp, err := l.client.HTTPClient.Get(url)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	var respData models.HueResponse[models.Light]
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return nil, nil, err
	}

	return respData.Data, respData.Errors, nil
}

func (l *LightService) GetLightByID(id string) (*models.Light, *models.HueError, error) {
	url := l.client.CreateURL("resource/light/" + id)
	resp, err := l.client.HTTPClient.Get(url)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	var respData models.HueResponse[models.Light]
	if err := json.NewDecoder(resp.Body).Decode(&respData); err != nil {
		return nil, nil, err
	}

	if len(respData.Data) > 0 {
		return &respData.Data[0], nil, nil
	}
	return nil, &respData.Errors[0], nil
}
