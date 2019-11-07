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

			// Questo codice è identico al codice di quadrati_righe_alterne_1.go (esercizio 7)
			// in quel codice si distingueva tra righe di indice pari e di indice dispari.
			// In questo caso però, si devono alternare colonne e non righe, quindi l'indice da usare
			// è quello di colonna
			if colonna%2 == 0 {
				fmt.Print("* ")
			} else {
				fmt.Print("+ ")
			}
		}
		fmt.Println()
	}

}
