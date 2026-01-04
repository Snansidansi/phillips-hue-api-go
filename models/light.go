package models

import (
	"time"
)

type Light struct {
	ID        string `json:"id"`
	IDV1      string `json:"id_v1,omitempty"`
	Owner     Owner  `json:"owner"`
	Type      string `json:"type"`
	ServiceID int    `json:"service_id"`

	Metadata Metadata `json:"metadata"`

	On       On     `json:"on"`
	Mode     string `json:"mode"`
	Identify *struct {
		Action string `json:"action,omitempty"`
	} `json:"identify,omitempty"`

	Dimming               *Dimming               `json:"dimming,omitempty"`
	DimmingDelta          *DimmingDelta          `json:"dimming_delta,omitempty"`
	ColorTemperature      *ColorTemperature      `json:"color_temperature,omitempty"`
	ColorTemperatureDelta *ColorTemperatureDelta `json:"color_temperature_delta,omitempty"`
	Color                 *Color                 `json:"color,omitempty"`
	Dynamics              *Dynamics              `json:"dynamics,omitempty"`
	Alert                 *Alert                 `json:"alert,omitempty"`
	Signaling             *Signaling             `json:"signaling,omitempty"`
	Gradient              *Gradient              `json:"gradient,omitempty"`
	Effects               *Effects               `json:"effects,omitempty"`
	EffectsV2             *EffectsV2             `json:"effects_v2,omitempty"`
	TimedEffects          *TimedEffects          `json:"timed_effects,omitempty"`
	PowerUp               *PowerUp               `json:"powerup,omitempty"`
}

type Owner struct {
	RID   string `json:"rid"`
	RType string `json:"rtype"`
}

type Metadata struct {
	Name        string       `json:"name"`
	Archetype   string       `json:"archetype"`
	Function    string       `json:"function"`
	ProductData *ProductData `json:"product_data,omitempty"`
}

type ProductData struct {
	Name      string `json:"name,omitempty"`
	Archetype string `json:"archetype,omitempty"`
	Function  string `json:"function,omitempty"`
}

type On struct {
	On bool `json:"on"`
}

type Dimming struct {
	Brightness  float64  `json:"brightness"`
	MinDimLevel *float64 `json:"min_dim_level,omitempty"`
}

type DimmingDelta struct {
	Action          string  `json:"action"`
	BrightnessDelta float64 `json:"brightness_delta,omitempty"`
}

type ColorTemperature struct {
	Mirek       *int         `json:"mirek"`
	MirekValid  bool         `json:"mirek_valid"`
	MirekSchema *MirekSchema `json:"mirek_schema,omitempty"`
}

type MirekSchema struct {
	MirekMinimum int `json:"mirek_minimum"`
	MirekMaximum int `json:"mirek_maximum"`
}

type ColorTemperatureDelta struct {
	Action     string `json:"action"`
	MirekDelta int    `json:"mirek_delta,omitempty"`
}

type Color struct {
	XY        XY     `json:"xy"`
	Gamut     *Gamut `json:"gamut,omitempty"`
	GamutType string `json:"gamut_type,omitempty"`
}

type XY struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type Gamut struct {
	Red   XY `json:"red"`
	Green XY `json:"green"`
	Blue  XY `json:"blue"`
}

type Dynamics struct {
	Status       string   `json:"status"`
	StatusValues []string `json:"status_values"`
	Speed        float64  `json:"speed"`
	SpeedValid   bool     `json:"speed_valid"`
}

type Alert struct {
	ActionValues []string `json:"action_values"`
}

type Signaling struct {
	SignalValues []string `json:"signal_values"`
	Status       *struct {
		Signal       string     `json:"signal"`
		EstimatedEnd *time.Time `json:"estimated_end,omitempty"`
	} `json:"status,omitempty"`
	Colors []ColorFeatureBasic `json:"colors,omitempty"`
}

type ColorFeatureBasic struct {
	XY XY `json:"xy"`
}

type Gradient struct {
	Points        []GradientPoint `json:"points"`
	Mode          string          `json:"mode"`
	PointsCapable int             `json:"points_capable"`
	ModeValues    []string        `json:"mode_values"`
	PixelCount    int             `json:"pixel_count,omitempty"`
}

type GradientPoint struct {
	Color ColorFeatureBasic `json:"color"`
}

type Effects struct {
	Status       string   `json:"status"`
	StatusValues []string `json:"status_values"`
	EffectValues []string `json:"effect_values"`
}

type EffectsV2 struct {
	Status       *EffectStatus `json:"status,omitempty"`
	EffectValues []string      `json:"effect_values"`
	Action       *EffectAction `json:"action,omitempty"`
}

type EffectStatus struct {
	Effect string `json:"effect"`
}

type EffectAction struct {
	Effect     string            `json:"effect"`
	Parameters *EffectParameters `json:"parameters,omitempty"`
}

type EffectParameters struct {
	Color            *ColorFeatureBasic `json:"color,omitempty"`
	ColorTemperature *ColorTemperature  `json:"color_temperature,omitempty"`
	Speed            float64            `json:"speed,omitempty"`
}

type TimedEffects struct {
	Status       string   `json:"status"`
	StatusValues []string `json:"status_values"`
	EffectValues []string `json:"effect_values"`
	Duration     *int     `json:"duration,omitempty"`
}

type PowerUp struct {
	Preset     string `json:"preset"`
	Configured bool   `json:"configured"`
	On         *struct {
		Mode string `json:"mode"`
		On   *On    `json:"on,omitempty"`
	} `json:"on,omitempty"`
	Dimming *struct {
		Mode    string   `json:"mode"`
		Dimming *Dimming `json:"dimming,omitempty"`
	} `json:"dimming,omitempty"`
	Color *struct {
		Mode             string            `json:"mode"`
		ColorTemperature *ColorTemperature `json:"color_temperature,omitempty"`
		Color            *Color            `json:"color,omitempty"`
	} `json:"color,omitempty"`
	ContentConfiguration *struct {
		Orientation *struct {
			Status       string `json:"status"`
			Configurable bool   `json:"configurable"`
			Orientation  string `json:"orientation"`
		} `json:"orientation,omitempty"`
		Order *struct {
			Status       string `json:"status"`
			Configurable bool   `json:"configurable"`
			Order        string `json:"order"`
		} `json:"order,omitempty"`
	} `json:"content_configuration,omitempty"`
}

type LightPut struct {
	Type                  *string                   `json:"type,omitempty"`
	Metadata              *MetadataPut              `json:"metadata,omitempty"`
	Function              *string                   `json:"function,omitempty"`
	Identify              *IdentifyPut              `json:"identify,omitempty"`
	On                    *OnPut                    `json:"on,omitempty"`
	Dimming               *DimmingPut               `json:"dimming,omitempty"`
	DimmingDelta          *DimmingDeltaPut          `json:"dimming_delta,omitempty"`
	ColorTemperature      *ColorTemperaturePut      `json:"color_temperature,omitempty"`
	ColorTemperatureDelta *ColorTemperatureDeltaPut `json:"color_temperature_delta,omitempty"`
	Color                 *ColorPut                 `json:"color,omitempty"`
	Dynamics              *DynamicsPut              `json:"dynamics,omitempty"`
	Alert                 *AlertPut                 `json:"alert,omitempty"`
	Signaling             *SignalingPut             `json:"signaling,omitempty"`
	Gradient              *GradientPut              `json:"gradient,omitempty"`
	Effects               *EffectsPut               `json:"effects,omitempty"` // Deprecated
	EffectsV2             *EffectsV2Put             `json:"effects_v2,omitempty"`
	TimedEffects          *TimedEffectsPut          `json:"timed_effects,omitempty"`
	PowerUp               *PowerUpPut               `json:"powerup,omitempty"`
	ContentConfiguration  *ContentConfigurationPut  `json:"content_configuration,omitempty"`
}

type IdentifyPut struct {
	Action   *string `json:"action,omitempty"`
	Duration *int64  `json:"duration,omitempty"`
}

type OnPut struct {
	On *bool `json:"on,omitempty"`
}

type DimmingPut struct {
	Brightness *float64 `json:"brightness,omitempty"`
}

type DimmingDeltaPut struct {
	Action          *string  `json:"action,omitempty"`
	BrightnessDelta *float64 `json:"brightness_delta,omitempty"`
}

type ColorTemperaturePut struct {
	Mirek *int `json:"mirek,omitempty"`
}

type ColorTemperatureDeltaPut struct {
	Action     *string `json:"action,omitempty"`
	MirekDelta *int    `json:"mirek_delta,omitempty"`
}

type ColorPut struct {
	XY *XYPut `json:"xy,omitempty"`
}

type XYPut struct {
	X *float64 `json:"x,omitempty"`
	Y *float64 `json:"y,omitempty"`
}

type DynamicsPut struct {
	Duration *int64   `json:"duration,omitempty"`
	Speed    *float64 `json:"speed,omitempty"`
}

type AlertPut struct {
	Action *string `json:"action,omitempty"`
}

type SignalingPut struct {
	Signal   *string                `json:"signal,omitempty"`
	Duration *int64                 `json:"duration,omitempty"`
	Colors   []ColorFeatureBasicPut `json:"colors,omitempty"`
}

type ColorFeatureBasicPut struct {
	XY *XYPut `json:"xy,omitempty"`
}

type GradientPut struct {
	Points []GradientPointPut `json:"points,omitempty"`
	Mode   *string            `json:"mode,omitempty"`
}

type GradientPointPut struct {
	Color *ColorFeatureBasicPut `json:"color,omitempty"`
}

type EffectsPut struct {
	Effect *string `json:"effect,omitempty"`
}

type EffectsV2Put struct {
	Action *EffectActionPut `json:"action,omitempty"`
}

type EffectActionPut struct {
	Effect     *string              `json:"effect,omitempty"`
	Parameters *EffectParametersPut `json:"parameters,omitempty"`
}

type EffectParametersPut struct {
	Color            *ColorFeatureBasicPut `json:"color,omitempty"`
	ColorTemperature *ColorTemperaturePut  `json:"color_temperature,omitempty"`
	Speed            *float64              `json:"speed,omitempty"`
}

type TimedEffectsPut struct {
	Effect   *string `json:"effect,omitempty"`
	Duration *int64  `json:"duration,omitempty"`
}

type PowerUpPut struct {
	Preset               *string                  `json:"preset,omitempty"`
	On                   *PowerUpOnPut            `json:"on,omitempty"`
	Dimming              *PowerUpDimmingPut       `json:"dimming,omitempty"`
	Color                *PowerUpColorPut         `json:"color,omitempty"`
	ContentConfiguration *ContentConfigurationPut `json:"content_configuration,omitempty"`
}

type PowerUpOnPut struct {
	Mode *string `json:"mode,omitempty"`
	On   *OnPut  `json:"on,omitempty"`
}

type PowerUpDimmingPut struct {
	Mode    *string     `json:"mode,omitempty"`
	Dimming *DimmingPut `json:"dimming,omitempty"`
}

type PowerUpColorPut struct {
	Mode             *string              `json:"mode,omitempty"`
	ColorTemperature *ColorTemperaturePut `json:"color_temperature,omitempty"`
	Color            *ColorPut            `json:"color,omitempty"`
}

type ContentConfigurationPut struct {
	Orientation *OrientationPut `json:"orientation,omitempty"`
	Order       *OrderPut       `json:"order,omitempty"`
}

type OrientationPut struct {
	Orientation *string `json:"orientation,omitempty"`
}

type OrderPut struct {
	Order *string `json:"order,omitempty"`
}
