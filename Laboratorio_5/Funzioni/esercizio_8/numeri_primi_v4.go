// Stampa la sequenza di numeri primi inferiori ad una soglia data da standard input.
package main

import (
	"fmt"
	"math"
)

func main() {

	if soglia := LeggiNumero(); soglia > 0 {
		NumeriPrimi(soglia)
	} else {
		fmt.Println("La soglia inserita non è positiva")
	}

}

// LeggiNumero legge un numero da standard input e restituisce il suo valore.
// Se la lettura fallisce, restituisce il valore 0.
func LeggiNumero() (n int) {
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&n)
	return
}

// ÈPrimo calcola se un numero passato come argomento è primo.
// Restituisce true in caso affermativo, false altrimenti.
func ÈPrimo(n int) bool {
	for i:=2; i<=int(math.Sqrt(float64(n))); i++ {
		if n%i==0 {
			return false
		}
	}
	return true
}

// NumeriPrimi stampa a standard output la sequenza di numeri interi primi inferiori alla soglia passata come argomento.
func NumeriPrimi(n int) {
	fmt.Println("Numeri primi inferiori a", n)
	if n>2 {
		// in questa versione controlliamo solo i numeri dispari maggiori o uguali a 3.
		// infatti, a parte 2, nessun numero pari è primo e possiamo scartarli a priori.
		// Perciò, stampiamo 2 in output (l'unico numero pari primo), e poi proseguiamo il controllo solo sui numeri dispari
		fmt.Print(2, " ")
		for i := 3; i < n; i+=2 {
			if ÈPrimo(i) {
				fmt.Print(i, " ")
			}
		}
	}
	fmt.Println()
}