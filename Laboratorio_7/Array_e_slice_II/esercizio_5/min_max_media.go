package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {

	numeri := LeggiNumeri()
	fmt.Println("Minimo:", Minimo(numeri))
	fmt.Println("Massimo:", Massimo(numeri))
	fmt.Println("Media:", Media(numeri))

}

// Converto i valori passati da riga di comando in una slice di interi.
// ATTENZIONE: in questa versione non viene fatto un controllo dell'input
func LeggiNumeri() (numeri []int) {
    // la slice numeri viene creata di dimensione uguale alla slice
    // os.Args ma scartando il primo valore, che è il nome dell'eseguibile
    numeri = make([]int, len(os.Args)-1)
	for i, v := range os.Args[1:] {
	    numeri[i], _ = strconv.Atoi(v)
	}
	return
}

// Calcola il valore minimo in una slicedi interi.
func Minimo(numeri []int) (min int) {
	min = math.MaxInt64
	for _, n := range numeri {
		if min > n {
			min = n
		}
	}
	return
}

// Calcola il valore massimo in una slice di interi.
func Massimo(numeri []int) (max int) {
	max = math.MinInt64
	for _, n := range numeri {
		if max < n {
			max = n
		}
	}
	return
}

// Calcola la media artimetica di una slice di interi.
func Media(numeri []int) (media float64) {
	for _, n := range numeri {
        // la variabile media viene usata inizialmente come
        // accumulatore per memorizzare la somma
        // dei valori interi della slice
		media += float64(n)
	}
	media /= float64(len(numeri))
	return
}
