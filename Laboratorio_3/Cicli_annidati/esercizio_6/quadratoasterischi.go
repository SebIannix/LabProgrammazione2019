/*
Stampa un quadrato di asterischi

Dato un numero n letto a tastiera, stampa un quadrato n x n di * intervallati da spazi
*/
package main

import "fmt"

func main() {

	var dimensione int

	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&dimensione)

	// Il primo for serve per scorrere le righe del quadrato
	for riga := 0; riga < dimensione; riga++ {
		// Il secondo for per le colonne del quadrato
		for colonna := 0; colonna < dimensione; colonna++ {
			fmt.Print("* ")
		}

		// Al termine di ogni riga è necessario stampare un 'a capo'
		// altrimenti tutti gli asterischi sarebbero stampati su un'unica riga
		fmt.Println()
	}

}
