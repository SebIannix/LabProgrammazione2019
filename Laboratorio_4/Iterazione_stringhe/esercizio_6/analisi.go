package main

import (
	"fmt"
)

func main() {

	// Contatori per ottenere il numero di lettere minuscole, maiuscole, consonanti e vocali
	maiuscole, minuscole := 0, 0
	consonanti, vocali := 0, 0

	var stringaInput string
	fmt.Scan(&stringaInput)

	// Per ogni carattere della stringa di input...
	for _, l := range stringaInput {

		// controllo se il carattere è una lettera dell'alfabeto inglese
		if (l >= 'A' && l <= 'Z') || (l >= 'a' && l <= 'z') {
			// controllo se la lettere è maiuscolo
			if l >= 'A' && l <= 'Z' {
				maiuscole++       // incremento contatore maiuscole
				l = l - 'A' + 'a' // trasformo la lettera in minuscolo
				// ( equivalente a l=unicode.ToLower(l) )
			} else {
				// se la lettera non è maiuscola, è per forza minuscola
				minuscole++ // incremento contatore minuscole
			}

			// controllo se la lettera è una vocale.
			// controllo solo le vocali minuscole perché ho trasformato in minuscolo la lettera in riga 24.
			// senza quella trasformazione avrei dovuto controllare anche se
			// l == 'A' || l == 'E' || l == 'I' || l == 'O' || l == 'U'
			if l == 'a' || l == 'e' || l == 'i' || l == 'o' || l == 'u' {
				vocali++
			} else {
				consonanti++
			}
		}

	}
	fmt.Println("Maiuscole:", maiuscole)
	fmt.Println("Minuscole:", minuscole)
	fmt.Println("Vocali:", vocali)
	fmt.Println("Consonanti:", consonanti)
}
