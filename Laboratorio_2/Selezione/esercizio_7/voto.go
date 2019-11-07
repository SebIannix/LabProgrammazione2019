package main

import "fmt"

func main() {
	var voto int
	fmt.Print("Inserisci il voto: ")
	fmt.Scan(&voto)

	// questa sequenza di if viene eseguita a cascata fino a trovare una condizione vera
	// se la condizione del primo if è falsa si passa al secondo if, e così via.
	// se la condizione del terzo if è vera, tutti gli if successivi non vengono presi in considerazione
	if voto < 0 || voto > 100 {
		fmt.Println("Errore")
	} else if voto < 60 {
		fmt.Println("Insufficiente")
	} else if voto < 70 {
		fmt.Println("Sufficiente")
	} else if voto < 80 {
		fmt.Println("Buono")
	} else if voto < 90 {
		fmt.Println("Distinto")
	} else {
		fmt.Println("Ottimo")
	}
}
