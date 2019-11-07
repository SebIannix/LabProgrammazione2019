package main

import (
	"fmt"
)

func main() {
	numeriLetti := 0
	mediaAritmetica := 0.0

	fmt.Print("Inserisci una sequenza di numeri (interrompi con numero<=0): ")
	// questo ciclo è senza condizione e può terminare solamente con l'esecuzione dell'istruzione break
	for {
		var x float64
		fmt.Scan(&x)

		if x > 0 {
			numeriLetti++
			mediaAritmetica += x
		} else {
			// se il numero letto non è positivo (<=0), allora esco dal ciclo for
			break
		}
	}

	mediaAritmetica /= float64(numeriLetti)

	fmt.Println("Media aritmetica:", mediaAritmetica)

}
