/*
Il programma stampa una sequenza di numeri

Dato un numero n inserito da tastiera, il programma stampa tutti i numeri compresi tra 1 e n (estremi inclusi).
La sequenza è prodotta su un'unica linea di output, con i numeri separati da uno spazio.
*/
package main

import (
	"fmt"
)

func main() {

	var n int

	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&n)

	for i := 1; i < n; { // for i := 1; i < n { : manca la terza parte del ciclo for, che può essere anche vuota
		fmt.Print(i)
		i++
	}

	fmt.Println()

}
