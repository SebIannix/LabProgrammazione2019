package main

import "fmt"
import "math"

func main() {

	var raggio float64
	fmt.Print("Raggio = ")
	fmt.Scan(&raggio)

	area := raggio * raggio * math.Pi
	circonferenza := raggio * 2 * math.Pi

	fmt.Println("Circonferenza =", circonferenza)
	fmt.Println("Area =", area)

}
