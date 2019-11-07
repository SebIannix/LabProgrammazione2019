package main

import "fmt"

func main() {

	var a, b int

	fmt.Scan(&a, &b)

	var somma int // accumulatore per ottenere la somma di una sequenza di numeri
	// questo ciclo for permette di assegnare alla variabile i tutti i valori all'interno
	// dell'intervallo (a,b), estremi esclusi
	for i := a + 1; i < b; i++ {
		// devo sommare solo i valori pari
		// per testare se un numero è pari posso controllare il risultato del resto della divisione
		// tra il numero stesso e 2: se il resto è 0, allora è pari. Altrimenti è dispari
		if i%2 == 0 {
			somma += i
		}
	}

	fmt.Println("Somma =", somma)
}
