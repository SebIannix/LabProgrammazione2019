package main

import (
	"fmt"
	"math"
)

func main() {

	var cifreDecimali int
	var valore float64

	fmt.Print("Inserisci il valore: ")
	fmt.Scan(&valore)
	fmt.Print("Inserisci il numero di cifre dopo la virgola: ")
	fmt.Scan(&cifreDecimali)

	potenza := math.Pow(10, float64(cifreDecimali))
	valoreTroncato := math.Trunc(valore*potenza) / potenza
	fmt.Println("Valore troncato =", valoreTroncato)

	valoreArrotondato := math.Round(valore*potenza) / potenza
	fmt.Println("Valore arrotondato =", valoreArrotondato)

}
