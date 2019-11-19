package main

import "fmt"

func main() {

	soglia := LeggiNumero()
	if soglia>0 {
		NumeriAmichevoli(soglia)
	} else {
		fmt.Println("La soglia inserita non è positiva")
	}

}

func LeggiNumero() (n int) {
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&n)
	return
}

func SommaDivisoriPropri(n int) (somma int){
	for i:=1; i<n; i++ {
		if n%i==0 {
			somma += i
		}
	}
	return
}

func SonoAmichevoli(n, m int) bool {
	return SommaDivisoriPropri(n) == m && SommaDivisoriPropri(m) == n
}

func NumeriAmichevoli(soglia int) {
	fmt.Println("Numeri amichevoli a", soglia)

	for i:=1; i<soglia; i++ {
		for j := i+1; j < soglia; j++ {
			if SonoAmichevoli(i, j) {
				fmt.Print("(", i, ",", j, ") ")
			}
		}
	}
	fmt.Println()
}