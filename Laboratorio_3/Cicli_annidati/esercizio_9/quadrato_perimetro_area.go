/*
Stampa un quadrato di asterischi

Dato un numero n letto a tastiera, stampa un quadrato n x n con * sui bordi e + al centro
*/
package main

import "fmt"

func main() {

	var dimensione int

	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&dimensione)

	for riga := 0; riga < dimensione; riga++ {
		for colonna := 0; colonna < dimensione; colonna++ {

			// Il programma deve stampare un asterisco solo sui bordi del quadrato.
			// Per il bordo superiore e quello di sinistra basta controllare semplicemente
			// se gli indici hanno valore uguale a 0 (la prima riga o la prima colonna)
			// Per il bordo inferiore e quello di destra, si controlla in modo analogo
			// che l'indice sia uguale alla dimensione del quadrato -1 (ovvero l'ultima riga o l'ultima colonna)
			if riga == 0 || colonna == 0 || riga == dimensione-1 || colonna == dimensione-1 {
				fmt.Print("* ")
			} else {
				fmt.Print("+ ")
			}
		}
		fmt.Println()
	}

}
