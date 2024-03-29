package tempconv

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZero Kelvin = 0
	Freezing     Kelvin = 273.15
	Boiling      Kelvin = 373.15
)

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func KToC(k Kelvin) Celsius {
	return Celsius(k - 273.15)
}

func CToK(c Celsius) Kelvin {
	return Kelvin(c + 273.15)
}

func FtoK(f Fahrenheit) Kelvin {
	return CToK(FToC(f))
}

func KtoF(k Kelvin) Fahrenheit {
	return CToF(KToC(k))
}
