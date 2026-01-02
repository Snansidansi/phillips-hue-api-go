package models

type HueResponse[T any] struct {
	Errors []HueError `json:"errors"`
	Data   []T        `json:"data"`
}
