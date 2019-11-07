package main

import "fmt"

func main() {

	var numero int

	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&numero)

	if numero%3 == 0 && numero%5 == 0 {
		// Se il numero è sia multiplo di 3 che di 5
		fmt.Println("Fizz Buzz")
	} else if numero%3 == 0 {
		// Se il numero è multiplo solo di 3 (si arriva in questo ramo SOLO SE l'espressione dell'if in riga 12 era falsa)
		fmt.Println("Fizz")
	} else if numero%5 == 0 {
		// Se il numero è multiplo solo di 5 (si arriva in questo ramo SOLO SE le espressioni degli if in riga 12 e 15 erano false)
		fmt.Println("Buzz")
	} else {
		// si arriva in questo ramo solo se il numero non è multiplo di 3 e non è multiplo di 5
		// questo ramo else è superfluo e può essere rimosso in quanto il programma non deve stampare nulla
	}

}
