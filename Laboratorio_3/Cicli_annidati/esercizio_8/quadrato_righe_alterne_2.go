/*
Stampa un quadrato di *, + e o

Dato un numero n letto a tastiera, stampa un quadrato composto da n righe di *, + e o alternate
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

			// Essendoci 3 possibili casi (*, + e o)
			// basta controllare il resto della divisione tra l'indice e 3
			// per sapere quale dei 3 caratteri dobbiamo stampare
			switch riga % 3 {
			case 0:
				fmt.Print("* ")
			case 1:
				fmt.Print("+ ")
			case 2:
				fmt.Print("o ")
			}
		}
		fmt.Println()
	}

}
