package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Dimensione non sufficiente")
	} else {
		dimensione := 2*n + 1

		// l'approccio utilizzato è molto simile a quello riportato nel sorgente triangolo_v2.go
		// dell'esercizio 9:
		// creo due stringhe, una di spazi e una di asterischi, e ad ogni iterazione
		// stampo una fetta della stringa di spazi e una fetta della stringa di asterischi
		stringaSpazi := ""
		stringaAsterischi := ""

		for i := 0; i < dimensione; i++ {
			stringaSpazi += " "
			stringaAsterischi += "*"
		}

		// parte superiore del rombo:
		// ad ogni iterazione diminuisce il numero di spazi stampati e
		// aumenta il numero di asterischi stampati
		for i := 0; i < dimensione/2; i++ {
			fmt.Print(stringaSpazi[:dimensione/2-i], stringaAsterischi[:2*i+1], "\n")
		}

		// riga centrale del rombo: solo asterischi
		fmt.Println(stringaAsterischi)

		// parte inferiore del rombo:
		// ad ogni iterazione aumenta il numero di spazi stampati e
		// diminuisce il numero di asterischi stampati
		for i := dimensione/2 - 1; i >= 0; i-- {
			fmt.Print(stringaSpazi[:dimensione/2-i], stringaAsterischi[:2*i+1], "\n")
		}
	}

}
