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
			// Un modo per risolvere questo problema è quello di considerare le due colonne di * e le due colonne di +
			// come se stessimo considerando 4 caratteri diversi:
			// colonna 1 -> stampa di un *
			// colonna 2 -> stampa di un *
			// colonna 3 -> stampa di un +
			// colonna 4 -> stampa di un +
			// Quindi l'esercizio si può risolvere come l'esercizio 10, ma considerando 4 casi invece di 2

			if colonna%4 == 0 {
				fmt.Print("* ")
			} else if colonna%4 == 1 {
				fmt.Print("* ")
			} else if colonna%4 == 2 {
				fmt.Print("+ ")
			} else if colonna%4 == 3 {
				fmt.Print("+ ")
			}
			// ovviamente il primo e il secondo if possono essere accorpati,
			// come possono essere accorpati il terzo e quarto if

		}
		fmt.Println()
	}

}
