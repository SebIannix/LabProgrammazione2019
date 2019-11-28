package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	n := LeggiNumero()
	fmt.Println("Fattoriali:", Fattoriali(n))

}

// Il testo dell'esercizio chiede di leggere un singolo valore
// da riga di comando
func LeggiNumero() (n int) {
	n, _ = strconv.Atoi(os.Args[1])
	return
}

func Fattoriale(n int) int {
	fattoriale := 1
	for i:=1; i<=n; i++ {
		fattoriale *= i
	}
	return fattoriale
}

// Versione 1 della funzione Fattoriali.
// La funzione crea una slice di dimension n, 
// in modo da memorizzare
// in listaFattoriali[0] = 1!
// in listaFattoriali[1] = 2!
// in listaFattoriali[2] = 3!
// ...
func Fattoriali(n int) (listaFattoriali []int) {
	listaFattoriali = make([]int, n)
	for i:=1; i<=n; i++ {
        // il calcolo di ogni fattoriali viene demandato 
        // ad una funzione apposita
		listaFattoriali[i-1] = Fattoriale(i)
	}
	return listaFattoriali
}
