package hueapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/snansidansi/hueapi/models"
	"github.com/snansidansi/hueapi/util"
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

func (s *LightService) On(id string) (*models.HueActionResponse, error) {
	return s.SetOnOff(id, true)
}

func (s *LightService) Off(id string) (*models.HueActionResponse, error) {
	return s.SetOnOff(id, false)
}

func (s *LightService) SetOnOff(id string, on bool) (*models.HueActionResponse, error) {
	update := models.LightPut{
		On: &models.OnPut{
			On: &on,
		},
	}
	return s.SetLightState(id, update)
}

func (s *LightService) Rename(id string, name string) (*models.HueActionResponse, error) {
	update := models.LightPut{
		Metadata: &models.MetadataPut{
			Name: &name,
		},
	}
	return s.SetLightState(id, update)
}

func (s *LightService) SetBrightness(id string, brightness float64) (*models.HueActionResponse, error) {
	update := models.LightPut{
		Dimming: &models.DimmingPut{
			Brightness: &brightness,
		},
	}
	return s.SetLightState(id, update)
}

func (s *LightService) SetColor(id string, r, g, b int) (*models.HueActionResponse, error) {
	x, y := util.RGBToXY(r, g, b)
	return s.SetColorXY(id, x, y)
}

func (s *LightService) SetColorXY(id string, x, y float64) (*models.HueActionResponse, error) {
	update := models.LightPut{
		Color: &models.ColorPut{
			XY: &models.XYPut{
				X: &x,
				Y: &y,
			},
		},
	}
	return s.SetLightState(id, update)
}

func (s *LightService) SetTemperature(id string, mirek int) (*models.HueActionResponse, error) {
	update := models.LightPut{
		ColorTemperature: &models.ColorTemperaturePut{
			Mirek: &mirek,
		},
	}
	return s.SetLightState(id, update)
}

func (s *LightService) Identify(id string, durationMs int64) (*models.HueActionResponse, error) {
	action := "identify"
	update := models.LightPut{
		Identify: &models.IdentifyPut{
			Action:   &action,
			Duration: &durationMs,
		},
	}
	return s.SetLightState(id, update)
}
