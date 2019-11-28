package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	n := LeggiNumero()
	fmt.Println("Fattoriali:", Fattoriali(n))

}

func LeggiNumero() (n int) {
	n, _ = strconv.Atoi(os.Args[1])
	return
}

// Versione 2 della funzione Fattoriali.
// In questa versione il calcolo del fattoriale è integrato nella
// funzione Fattoriali.
// Questa versione è più performante perché usa il valore del fattoriale
// ad una iterazione per calcolare il valore del fattoriale all'iterazione successiva
func Fattoriali(n int) (listaFattoriali []int) {
    listaFattoriali = make([]int, n)
	fattoriale := 1
	for i:=1; i<=n; i++ {
		fattoriale *= i
		listaFattoriali[i-1] = fattoriale
	}
	return
}
