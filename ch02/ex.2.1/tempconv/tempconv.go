package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
	ZeroK         Kelvin  = -273.15
)

func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string { return fmt.Sprintf("%g°K", k) }

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9)}
func CToK(c Celsius) Kelvin { return Kelvin(1000.0 * (c - AbsoluteZeroC)) }
func KToC(k Kelvin) Celsius { return Celsius(k/1000.0 - Kelvin(float64(AbsoluteZeroC)))}

/**
 * Feet To Meter
 */
type Feet float64
func FeetToMeter(f Feet) Meter { return Meter(f*0.3048) }
func (f Feet) String() string { return fmt.Sprintf("%0.2f Feet", f) }

/**
 * Meter to Feet
 */
type Meter float64
func MeterToFeet(m Meter) Feet { return Feet(m/0.3048) }
func (m Meter) String() string { return fmt.Sprintf("%0.2f Meter", m) }

/**
 * Pound To Kilogram
 */
type Pound float64
func PoundToKilogram(p Pound) Kilogram { return Kilogram(p * 0.453592) }
func (p Pound) String() string { return fmt.Sprintf("%0.2f Pound", p) }

/**
 * Kilogram to Pound
 */
type Kilogram float64
func KilogramToPound(k Kilogram) Pound { return Pound(k / 0.453592) }
func (k Kilogram) String() string { return fmt.Sprintf("%0.2f Pound", k) }