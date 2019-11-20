package main

import (
	"fmt"
)

const Chiave = 5
const NumeroLettere = 26

func main() {
	s := "Guarda, è Zorro!"

	for _, carattere := range s {

		switch {
		case carattere >= 'A' && carattere <= 'Z':
			carattere = 'A' + ((carattere - 'A' + Chiave) % NumeroLettere)
		case carattere >= 'a' && carattere <= 'z':
			carattere = 'a' + ((carattere - 'a' + Chiave) % NumeroLettere)
		}

		fmt.Printf("%c",carattere)

	}

	fmt.Println()

	fmt.Println(s)
}
