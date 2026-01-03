package util

import (
	"math"
)

func XYToRGB(x, y float64, brightness float64) (r, g, b int) {
	Y := brightness / 100.0
	if y == 0 {
		y = 0.000001
	}

	X := (Y / y) * x
	Z := (Y / y) * (1.0 - x - y)

	rL := X*3.2406 - Y*1.5372 - Z*0.4986
	gL := -X*0.9689 + Y*1.8758 + Z*0.0415
	bL := X*0.0557 - Y*0.2040 + Z*1.0570

	applyGamma := func(v float64) float64 {
		if v <= 0.0031308 {
			return 12.92 * v
		}
		return 1.055*math.Pow(v, (1.0/2.4)) - 0.055
	}

	rS := applyGamma(rL)
	gS := applyGamma(gL)
	bS := applyGamma(bL)

	clamp := func(v float64) int {
		res := int(v * 255)
		if res < 0 {
			return 0
		}
		if res > 255 {
			return 255
		}
		return res
	}

	return clamp(rS), clamp(gS), clamp(bS)
}

func RGBToXY(r, g, b int) (x, y float64) {
	rN := float64(r) / 255.0
	gN := float64(g) / 255.0
	bN := float64(b) / 255.0

	gamma := func(v float64) float64 {
		if v > 0.04045 {
			return math.Pow((v+0.055)/1.055, 2.4)
		}
		return v / 12.92
	}

	rL := gamma(rN)
	gL := gamma(gN)
	bL := gamma(bN)

	X := rL*0.4124 + gL*0.3576 + bL*0.1805
	Y := rL*0.2126 + gL*0.7152 + bL*0.0722
	Z := rL*0.0193 + gL*0.1192 + bL*0.9505

	sum := X + Y + Z
	if sum == 0 {
		return 0, 0
	}

	return X / sum, Y / sum
}
