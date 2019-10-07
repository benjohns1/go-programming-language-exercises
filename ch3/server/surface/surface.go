package surface

import (
	"fmt"
	"io"
	"math"
)

const angle = math.Pi / 6

var sin30, cos30 = math.Sin(angle), math.Cos(angle)
var min, max = math.MaxFloat64, 0.0

type polygon struct {
	points string
	z      float64
}

// SvgSurface writes the surface data in SVG format
func SvgSurface(w io.Writer, width, height, cells int, xyrange float64) {
	var (
		xyscale float64 = float64(width) / 2.0 / xyrange
		zscale  float64 = float64(height) * 0.4
	)
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' style='stroke: black; fill: white; stroke-width: 0.7' width='%d' height='%d'>", width, height)
	polys := []polygon{}
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, az, err := corner(i+1, j, height, width, cells, xyrange, xyscale, zscale)
			if err != nil {
				continue
			}
			bx, by, bz, err := corner(i, j, height, width, cells, xyrange, xyscale, zscale)
			if err != nil {
				continue
			}
			cx, cy, cz, err := corner(i, j+1, height, width, cells, xyrange, xyscale, zscale)
			if err != nil {
				continue
			}
			dx, dy, dz, err := corner(i+1, j+1, height, width, cells, xyrange, xyscale, zscale)
			if err != nil {
				continue
			}
			var z float64
			z, min, max = avgMinMax(min, max, az, bz, cz, dz)
			polys = append(polys, polygon{points: fmt.Sprintf("%g, %g %g, %g %g, %g %g, %g", ax, ay, bx, by, cx, cy, dx, dy), z: z})
		}
	}
	for _, poly := range polys {
		color := getColor(poly.z, min, max)
		fmt.Fprintf(w, "<polygon points='%s' style='fill: #%06x;'/>", poly.points, color)
	}

	fmt.Fprintf(w, "</svg>")
}

func avgMinMax(min float64, max float64, zs ...float64) (z, newMin, newMax float64) {

	var count int
	var total float64
	for _, z := range zs {
		if z > max {
			max = z
		}
		if z < min {
			min = z
		}
		count++
		total += z
	}
	return (total / float64(count)), min, max
}

func getColor(z, min, max float64) (color int) {
	lerped := (z - min) / (max - min)
	blue := int(lerped * 256)
	red := int((1 - lerped) * 256)
	return (red * 0x10000) + blue
}

func corner(i, j, height, width, cells int, xyrange, xyscale, zscale float64) (float64, float64, float64, error) {
	x := xyrange * (float64(i)/float64(cells) - 0.5)
	y := xyrange * (float64(j)/float64(cells) - 0.5)

	z := f(x, y)
	if math.IsNaN(z) {
		return 0, 0, 0, fmt.Errorf("invalid value")
	}
	if math.IsInf(z, 0) {
		return 0, 0, 0, fmt.Errorf("infinity")
	}

	sx := float64(width)/2.0 + (x-y)*cos30*xyscale
	sy := float64(height)/2.0 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z, nil
}

func f(x, y float64) float64 {
	// r := math.Hypot(x, y)
	r := math.Sqrt(x*x + y*y)
	return math.Sin(r) / r
}
