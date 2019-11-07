/*
Pari o dispari

Dato un numero intero n, stampa se n è pari o dispari
*/
package main

import "fmt"

func main() {
	var n int
	fmt.Println("Inserisci un numero: ")
	fmt.Scanln(&n)

	// Per sapere se un numero è pari o fispari basta controllare il resto della sua divisione intera con il numero 2:
	// Se il resto è 0, allora il numero è pari, altrimenti è dispari.
	// Per calcolare il resto della divisione tra interi si usa l'operatore %
	if n%2 == 0 {
		fmt.Println(n, "è pari")
	} else {
		fmt.Println(n, "è dispari")
	}
}
