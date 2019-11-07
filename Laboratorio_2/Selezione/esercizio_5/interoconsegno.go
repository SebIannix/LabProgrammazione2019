package main

import "fmt"

func main() {

	var numeroInput int

	fmt.Print("Inserisci numero: ")
	fmt.Scan(&numeroInput)

	// se il numero è positivo deve essere aggiunto il segno +
	if numeroInput > 0 {
		fmt.Print("+", numeroInput, "\n")
	} else {
		// nel caso di un numero negativo, il segno non deve essere aggiunto in quanto viene già stampato
		// nel caso l'input sia invece 0, non serve il segno
		fmt.Println(numeroInput)
	}

}
