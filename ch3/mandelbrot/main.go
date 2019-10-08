package main

import (
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/benjohns1/go-programming-language-exercises/ch3/mandelbrot/colormap"
	"github.com/benjohns1/go-programming-language-exercises/ch3/mandelbrot/mandelbrot"
)

var (
	pink   = color.RGBA{0xFF, 0xAA, 0xAA, 0xFF}
	red    = color.RGBA{0xFF, 0x00, 0x00, 0xFF}
	orange = color.RGBA{0xFF, 0xA5, 0xFF, 0xFF}
	yellow = color.RGBA{0xFF, 0xFF, 0x00, 0xFF}
	green  = color.RGBA{0x00, 0x90, 0x00, 0xFF}
	blue   = color.RGBA{0x00, 0x00, 0xFF, 0xFF}
	indigo = color.RGBA{0x4b, 0x00, 0x82, 0xFF}
	violet = color.RGBA{0xEE, 0x88, 0xEE, 0xFF}
	gray   = color.RGBA{0x7F, 0x7F, 0x7F, 0xFF}
	white  = color.White
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		ydiff, xdiff           = ymax - ymin, xmax - xmin
		width, height          = 1024, 1024
		supersample            = 3
		subwidth, subheight    = width * supersample, height * supersample
		iterations             = 100
		multiplier             = 64
		maxColorRange          = float64(iterations*multiplier) / 4
	)
	g := colormap.Gradient{Colors: []color.Color{white, blue, indigo, violet, gray, white, pink, red, orange, yellow, green}, R: colormap.Range{Min: 0, Max: maxColorRange}}
	m := mandelbrot.Mandelbrot{Iterations: iterations}

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < subheight; py += supersample {
		var ys []float64
		for i := 0; i < supersample; i++ {
			ys = append(ys, float64(py+i)/subheight*ydiff+ymin)
		}
		for px := 0; px < subwidth; px += supersample {
			var allVals uint32
			allIn := true
			for i := 0; i < supersample; i++ {
				x := float64(px+i)/subwidth*xdiff + xmin
				for _, y := range ys {
					v := m.Factorial(x, y)
					allIn = allIn && v.In
					allVals += uint32(v.N)
				}
			}
			var c color.Color
			if allIn {
				c = color.Black
			} else {
				floatVal := float64(allVals) / float64(supersample*supersample)
				c = g.FloatToColor(float64(floatVal * multiplier))
			}
			img.Set(px/supersample, py/supersample, c)
		}
	}
	png.Encode(os.Stdout, img)
}
