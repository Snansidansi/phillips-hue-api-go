package hueapi

import "fmt"

type HueError struct {
	Type        int    `json:"type"`
	Description string `json:"description"`
}

func (e *HueError) Error() string {
	return fmt.Sprintf("Hue Error #%d: %s", e.Type, e.Description)
}

const (
	ErrLinkButtonNotPressed = 101
)
