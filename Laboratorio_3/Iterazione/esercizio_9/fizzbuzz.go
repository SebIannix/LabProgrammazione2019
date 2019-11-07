package main

import "fmt"

func main() {

	var n int

	fmt.Scan(&n)

	// Questo programma riprende il codice dell'esercizio 8 in Laboratorio_2/Selezione
	// e applica il test ad ogni valore compreso tra 1 e n
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Print("FizzBuzz ")
		} else if i%3 == 0 {
			fmt.Print("Fizz ")
		} else if i%5 == 0 {
			fmt.Print("Buzz ")
		} else {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()

}
