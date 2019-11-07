package main

import "fmt"
import "math"

func main() {

	// Dichiaro la variabile raggio come numero reale
	var raggio float64
	fmt.Print("Raggio = ")
	fmt.Scan(&raggio)

	// Calcolo area e circonferenza usando la costante Pi del package math
	area := raggio * raggio * math.Pi
	circonferenza := raggio * 2 * math.Pi

	fmt.Println("Circonferenza =", circonferenza)
	fmt.Println("Area =", area)

}
