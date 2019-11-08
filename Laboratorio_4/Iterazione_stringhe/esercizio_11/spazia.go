package main

import "fmt"

func main() {

	var testo string
	fmt.Print("Inserisci una stringa di testo: ")
	fmt.Scan(&testo)

	// questo ciclo for stampa ogni carattere del testo inserito con l'aggiunta di uno spazio
	for _, c := range testo {
		fmt.Printf("%c ", c)
	}
	fmt.Println()

}
