package main

import "fmt"

func main() {

	var numeroInput int

	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&numeroInput)

	// Per controllare se un numero è multiplo di 10 basta usare l'operatore %
	// L'operatore % da come risultato il resto della divisione tra interi.
	// Ad esempio, 17%10 da come risultato 7.
	// Se il risultato di numeroInput%10 è 0, allora significa che numeroInput è multiplo di 10.
	// L'operatore % funziona ovviamente con qualsiasi coppia di numeri interi, ad esempio:
	// 7%2 da il resto della divisione intera tra 7 e 2, ovvero 1
	if numeroInput%10 == 0 {
		fmt.Println(numeroInput, "è multiplo di 10")
	} else {
		fmt.Println(numeroInput, "non è multiplo di 10")
	}

}
