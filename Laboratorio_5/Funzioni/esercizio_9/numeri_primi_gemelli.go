package main

import (
	"fmt"
	"math"
)

func main() {

	if soglia := LeggiNumero(); soglia > 0 {
		NumeriPrimiGemelli(soglia)
	} else {
		fmt.Println("La soglia inserita non è positiva")
	}

}

func LeggiNumero() (n int) {
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&n)
	return
}

func ÈPrimo(n int) bool {
	for i:=2; i<=int(math.Sqrt(float64(n))); i++ {
		if n%i==0 {
			return false
		}
	}
	return true
}

func NumeriPrimiGemelli(n int) {
	fmt.Println("Numeri primi gemelli inferiori a", n)
	for i:=2; i<n-2; i++ {
		if ÈPrimo(i) && ÈPrimo(i+2) {
			fmt.Print("(", i, ",", i+2, ") ")
		}
	}
	fmt.Println()
}