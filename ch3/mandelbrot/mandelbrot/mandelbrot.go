package mandelbrot

import "math/cmplx"

// Value describes a value at a single point
type Value struct {
	In bool
	N  uint16
}

// Mandelbrot is a config object for running the Factorial function to return values in the mandelbrot set
type Mandelbrot struct {
	Iterations uint16
}

// Factorial returns the Value at a location
func (m *Mandelbrot) Factorial(z complex128) Value {
	var v complex128
	for n := uint16(0); n < m.Iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return Value{N: n}
		}
	}
	return Value{In: true}
}
