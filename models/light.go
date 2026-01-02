package models

type Light struct {
	ID       string `json:"id"`
	Metadata struct {
		Name string `json:"name"`
	} `json:"metadata"`

	On struct {
		On bool `json:"on"`
	} `json:"on"`

	Dimming *struct {
		Brightness float64 `json:"brightness"`
	} `json:"dimming"`

	Color *struct {
		XY struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"xy"`
	} `json:"color"`

	ColorTemperature *struct {
		Mirek int `json:"mirek"`
	} `json:"color_temperature"`
}
