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

			// in un quadrato, i caratteri che corrispondono alla diagonale hanno
			// riga e colonna uguali
			if riga == colonna {
				fmt.Print("o ")
			} else if riga < colonna {
				// quando l'indice di riga è inferiore all'indice di colonna ci troviamo
				// nel triangolo superiore del quadrato
				fmt.Print("+ ")
			} else {
				// questo ultimo ramo rappresenta il caso in cui l'indice di riga è superiore all'indice di colonna
				// e ci troviamo nel triangolo inferiore del quadrato
				fmt.Print("* ")
			}
		}
		fmt.Println()
	}

}
