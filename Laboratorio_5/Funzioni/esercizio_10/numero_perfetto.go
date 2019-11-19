package main

import "fmt"

func main() {

	if n := LeggiNumero(); ÈPerfetto(n) {
		fmt.Println(n, "è perfetto")
	} else {
		fmt.Println(n, "non è perfetto")
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

func ÈPerfetto(n int) bool {
	return n !=0 && SommaDivisoriPropri(n) == n
}
