package main

import "fmt"

func main() {

	NumeriAbbondanti(LeggiNumero())

}

func LeggiNumero() (n int) {
	fmt.Scan(&n)
	return
}

func SommaDivisoriPropri(n int) (somma int) {
	for i:=1; i<n; i++ {
		if n%i==0 {
			somma += i
		}
	}
	return
}

func ÈAbbondante(n int) bool {
	return SommaDivisoriPropri(n) > n
}

func NumeriAbbondanti(soglia int) {
	fmt.Print("Numeri abbondanti:")
	for i:=1; i<soglia; i++ {
		if ÈAbbondante(i) {
			fmt.Print(" ", i)
		}
	}
	fmt.Println()
}