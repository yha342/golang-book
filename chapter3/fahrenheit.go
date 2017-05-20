package main

import "fmt"

func main () {
	fmt.Print("Enter Temperatur in Fahrenheit: ")
	var inputFahrenheit float64
	fmt.Scanf("%f", &inputFahrenheit)

	outputCelsius := (inputFahrenheit - 32) * 5/9

	fmt.Println (outputCelsius)

}
