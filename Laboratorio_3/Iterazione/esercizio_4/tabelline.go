/*
Calcola la tabellina di un numero

Letto un numero n da tastiera, il package calcola la sua tabellina
*/
package main

import "fmt"

func main() {
	var numero int

	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&numero)

	// questo ciclo for permette di ripetere le istruzioni contenute per 10 volte
	for i := 1; i <= 10; i++ {
		// supponendo che numero=5
		// alla prima iterazione i=1 e risultato=1*5
		// alla seconda iterazione i=2 e risultato=2*5
		// alla terza iterazione i=3 e risultato=3*5
		// e così fino a i=11: in questo caso la condizione del for è falsa e il ciclo for termina
		risultato := i * numero
		fmt.Println(i, "x", numero, "=", risultato)
	}
}
