package main

import "fmt"

func main() {

	var testo string
	fmt.Scan(&testo)

	fmt.Print("Testo maiuscolo: ")
	// Per ogni carattere del testo...
	for _, carattere := range testo {
		// se il carattere è una lettera minuscola dell'alfabeto inglese,
		// lo trasformo in una lettera maiuscola
		// (equivalente a carattere = unicode.ToUpper(carattere) )
		if carattere >= 'a' && carattere <= 'z' {
			carattere = 'A' + (carattere - 'a')
		}
		fmt.Printf("%c", carattere)
	}
	fmt.Println()

	fmt.Print("Testo minuscolo: ")
	// Per ogni carattere del testo...
	for _, carattere := range testo {
		// se il carattere è una lettera maiuscola dell'alfabeto inglese,
		// lo trasformo in una lettera minuscola
		// (equivalente a carattere = unicode.ToLower(carattere) )
		if carattere >= 'A' && carattere <= 'Z' {
			carattere = 'a' + (carattere - 'A')
		}
		fmt.Printf("%c", carattere)
	}
	fmt.Println()

}
