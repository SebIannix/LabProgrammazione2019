package main

import "fmt"

func main() {
	var numero int

	fmt.Print("Inserisci numero: ")
	fmt.Scan(&numero)

	fmt.Print("Divisori di ", numero, ": ")
	// Controllo se numero è multiplo di i per ogni valore compreso tra 1 e numero/2
	// Andare oltre numero/2 è inutile, perché nessun valore>numero/2 può essere divisore di numero
	for i := 1; i <= numero/2; i++ {
		// un numero è divisibile per i solo se il resto della divisione tra numero e i è 0
		if numero%i == 0 {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}
