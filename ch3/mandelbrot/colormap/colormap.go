package colormap

import (
	"image/color"
	"math"
)

// Range min/max
type Range struct {
	Min float64
	Max float64
}

// Gradient list of colors and range
type Gradient struct {
	Colors color.Palette
	R      Range
}

const conv32to8 = math.MaxInt8 / math.MaxInt32

// FloatToColor converts an integer to a color in the gradient
func (g *Gradient) FloatToColor(n float64) color.Color {
	ccount := len(g.Colors)
	if ccount == 0 {
		return nil
	}
	if ccount == 1 {
		return g.Colors[0]
	}
	normalized := (n - g.R.Min) * float64(ccount-1) / (g.R.Max - g.R.Min)
	index := int(normalized)
	if index >= ccount-1 {
		return g.Colors[ccount-1]
	}
	ratio := normalized - math.Floor(normalized)
	r1, g1, b1, a1 := g.Colors[index].RGBA()
	r2, g2, b2, a2 := g.Colors[index+1].RGBA()

	return color.RGBA{combine(r1, r2, ratio), combine(g1, g2, ratio), combine(b1, b2, ratio), combine(a1, a2, ratio)}
}

func combine(first, second uint32, ratio float64) uint8 {
	return uint8((float64(uint8(first>>8)) * (1 - ratio)) + (float64(uint8(second>>8)) * ratio))
}
