package models

type HueResponse[T any] struct {
	StatusCode int `json:"-"`

	Errors []struct {
		Description string `json:"description"`
	} `json:"errors"`

	Data []T `json:"data,omitempty"`
}

type HueActionResponse struct {
	StatusCode int `json:"-"`

	Errors []struct {
		Description string `json:"description"`
	} `json:"errors"`

	Data []struct {
		Rid   string `json:"rid"`
		Rtype string `json:"rtype"`
	} `json:"data,omitempty"`
}
