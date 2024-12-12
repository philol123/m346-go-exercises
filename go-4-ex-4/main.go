package main

import (
	"fmt"
)

func convertCelsiusToFahrenheit(celsius float64) float64 {
	return (celsius * 9 / 5) + 32
}

func convertFahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}

type Celsius float64

func (c Celsius) ConvertToFahrenheit() float64 {
	return (float64(c) * 9 / 5) + 32
}

type Fahrenheit float64

func (f Fahrenheit) ConvertToCelsius() float64 {
	return (float64(f) - 32) * 5 / 9
}

func main() {
	tempC1 := 23.0
	fmt.Printf("23°C in Fahrenheit: %.2f\n", convertCelsiusToFahrenheit(tempC1)) // Output: 73.40

	tempC2 := 0.0
	fmt.Printf("0°C in Fahrenheit: %.2f\n", convertCelsiusToFahrenheit(tempC2)) // Output: 32.00

	tempC3 := -10.0
	fmt.Printf("-10°C in Fahrenheit: %.2f\n", convertCelsiusToFahrenheit(tempC3)) // Output: 14.00

	tempF1 := 73.40
	fmt.Printf("73.40°F in Celsius: %.2f\n", convertFahrenheitToCelsius(tempF1)) // Output: 23.00

	tempF2 := 32.00
	fmt.Printf("32°F in Celsius: %.2f\n", convertFahrenheitToCelsius(tempF2)) // Output: 0.00

	tempF3 := 14.00
	fmt.Printf("14°F in Celsius: %.2f\n", convertFahrenheitToCelsius(tempF3)) // Output: -10.00

	var cozy Celsius = 23.0
	fmt.Printf("23°C to Fahrenheit using method: %.2f\n", cozy.ConvertToFahrenheit()) // Output: 73.40

	var cold Fahrenheit = 15.3
	fmt.Printf("15.3°F to Celsius using method: %.2f\n", cold.ConvertToCelsius()) // Output: -9.67
}

// Ich finde folgende notation übersichtlicher und einfacher zu lesen:
// fmt.Println(convertFahrenheitToCelsius(convertCelsiusToFahrenheit(23.0)))
