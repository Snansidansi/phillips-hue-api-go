package hueapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/snansidansi/hueapi/models"
)

type RoomService struct {
	client *Client
}

func (s *RoomService) GetAllRooms() (*models.HueResponse[models.Room], error) {
	url := s.client.CreateURL("resource/room")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var hueResp models.HueResponse[models.Room]
	hueResp.StatusCode = resp.StatusCode

	if err := json.NewDecoder(resp.Body).Decode(&hueResp); err != nil {
		return &hueResp, fmt.Errorf("decoding failed (status %d): %w", resp.StatusCode, err)
	}

	return &hueResp, nil
}

func (s *RoomService) GetRoomByID(id string) (*models.HueResponse[models.Room], error) {
	url := s.client.CreateURL("resource/room/" + id)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := s.client.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var hueResp models.HueResponse[models.Room]
	hueResp.StatusCode = resp.StatusCode

	if err := json.NewDecoder(resp.Body).Decode(&hueResp); err != nil {
		return &hueResp, fmt.Errorf("decoding failed (status %d): %w", resp.StatusCode, err)
	}

	return &hueResp, nil
}

func (s *RoomService) UpdateRoom(id string, room models.RoomPut) (*models.HueActionResponse, error) {
	url := s.client.CreateURL("resource/room/" + id)

	jsonData, err := json.Marshal(room)
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
		return &hueResp, fmt.Errorf("decoding failed (status %d): %w", resp.StatusCode, err)
	}

	return &hueResp, nil
}

func (s *RoomService) CreateRoom(room models.RoomPost) (*models.HueActionResponse, error) {
	url := s.client.CreateURL("resource/room")

	jsonData, err := json.Marshal(room)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
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
		return &hueResp, fmt.Errorf("decoding failed (status %d): %w", resp.StatusCode, err)
	}

	return &hueResp, nil
}
