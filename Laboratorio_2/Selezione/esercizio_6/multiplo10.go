package main

import "fmt"

func main() {

	var numero int

	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&numero)

	if numero%10 == 0 {
		fmt.Println(numero, "è multiplo di 10")
	} else {
		fmt.Println(numero, "non è multiplo di 10")
	}

}
