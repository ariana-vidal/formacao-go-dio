package main

func main() { // Declaração de variáveis
	var celsius float64
	var fahrenheit float64
	var kelvin float64

	// Conversão de Celsius para Fahrenheit
	celsius = 25.0
	fahrenheit = (celsius * 9 / 5) + 32
	println(celsius, "°C é igual a", fahrenheit, "°F")

	// Conversão de Fahrenheit para Celsius
	fahrenheit = 77.0
	celsius = (fahrenheit - 32) * 5 / 9
	println(fahrenheit, "°F é igual a", celsius, "°C")

	// Conversão de Celsius para Kelvin
	celsius = 25.0
	kelvin = celsius + 273.15
	println(celsius, "°C é igual a", kelvin, "K")

	// Conversão de Kelvin para Celsius
	kelvin = 298.15
	celsius = kelvin - 273.15
	println(kelvin, "K é igual a", celsius, "°C")

}
