/*
Stampa un quadrato di * e +

Dato un numero n letto a tastiera, stampa un quadrato composto da n righe di * e + alternate
*/
package main

import (
	"fmt"
)

func main() {

	var dimensione int
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&dimensione)

	for riga := 0; riga < dimensione; riga++ {
		for colonna := 0; colonna < dimensione; colonna++ {

			// Possiamo distinguere i due casi '*' e '+' semplicemente controllando se l'indice di riga
			// è multiplo di 2.
			// In questo caso, con righe di indice pari stamperemo *, mentre con righe di indice dispari dei +
			if riga%2 == 0 {
				fmt.Print("* ")
			} else {
				fmt.Print("+ ")
			}
		}
		fmt.Println()
	}

}
