package builders

import (
	"github.com/snansidansi/hueapi/models"
	"github.com/snansidansi/hueapi/util"
)

type LightBuilder struct {
	put models.LightPut
}

func NewLightBuilder() *LightBuilder {
	return &LightBuilder{
		put: models.LightPut{},
	}
}

func (b *LightBuilder) On() *LightBuilder {
	state := true
	if b.put.On == nil {
		b.put.On = &models.OnPut{}
	}
	b.put.On.On = &state
	return b
}

func (b *LightBuilder) Off() *LightBuilder {
	state := false
	if b.put.On == nil {
		b.put.On = &models.OnPut{}
	}
	b.put.On.On = &state
	return b
}

func (b *LightBuilder) SetOnOff(state bool) *LightBuilder {
	if b.put.On == nil {
		b.put.On = &models.OnPut{}
	}
	b.put.On.On = &state
	return b
}

func (b *LightBuilder) Brightness(percent float64) *LightBuilder {
	if b.put.Dimming == nil {
		b.put.Dimming = &models.DimmingPut{}
	}
	b.put.Dimming.Brightness = &percent
	return b
}

func (b *LightBuilder) ColorXY(x, y float64) *LightBuilder {
	if b.put.Color == nil {
		b.put.Color = &models.ColorPut{XY: &models.XYPut{}}
	}
	b.put.Color.XY.X = &x
	b.put.Color.XY.Y = &y
	return b
}

func (b *LightBuilder) ColorRGB(r, g, b_int int) *LightBuilder {
	x, y := util.RGBToXY(r, g, b_int)
	return b.ColorXY(x, y)
}

func (b *LightBuilder) Temperature(mirek int) *LightBuilder {
	if b.put.ColorTemperature == nil {
		b.put.ColorTemperature = &models.ColorTemperaturePut{}
	}
	b.put.ColorTemperature.Mirek = &mirek
	return b
}

func (b *LightBuilder) Duration(ms int64) *LightBuilder {
	if b.put.Dynamics == nil {
		b.put.Dynamics = &models.DynamicsPut{}
	}
	b.put.Dynamics.Duration = &ms
	return b
}

func (b *LightBuilder) Build() models.LightPut {
	return b.put
}
