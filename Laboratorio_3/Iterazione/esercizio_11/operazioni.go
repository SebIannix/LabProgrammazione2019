/*
Operazioni numeriche

Effettua delle operazioni tra due numeri interi inseriti da tastiera.
Le operazioni sono:
- individuazione del minimo e del massimo
- somma, sottrazione, prodotto e divisione
- elevamento a potenza
- media
*/
package main

import (
	"fmt"
	"math"
)

func main() {

	var x, y int

	fmt.Println("Inserisci due numeri interi: ")
	fmt.Scan(&x, &y)

	// inizializzo una variabile max per trovare il maggiore tra i due valori.
	// max prende il valore di x, e se tale valore è minore del valore di y,
	// allora assegno a max il valore di y, trovando così il valore massimo tra x e y
	max := x
	if max < y {
		max = y
	}

	min := x
	if min > y {
		min = y
	}

	fmt.Println("Maggiore:", max)
	fmt.Println("Minore:", min)

	fmt.Println("Somma:", x+y)
	fmt.Println("Differenza:", max-min)
	fmt.Println("Prodotto:", x*y)
	fmt.Println("Divisione:", float64(x)/float64(y))

	fmt.Println("Valore medio:", (x+y)/2)

	// La potenza "x elevato alla y" si può calcolare come il prodotto di x per sé stesso y volte
	// si parte da un accumulatore opportunamente inizializzato (la variabile potenza)
	// e ad ogni iterazione si effettua il prodotto della potenza per x
	potenza := 1 // quando si tratta di un prodotto, l'accumulatore si inizializza ad 1
	for i := 0; i < y; i++ {
		potenza *= x
		// supponendo che x=3 e y=5
		// alla prima iterazione (i=0) avremo potenza = 1 * 3
		// alla seconda iterazione (i=1) avremo potenza = 3 * 3
		// alla terza iterazione (i=2) avremo potenza = 9 * 3
		// e così via per 5 iterazioni totali

	}
	fmt.Println("Potenza (ciclo for):", potenza)
	fmt.Println("Potenza (math.Pow):", int(math.Pow(float64(x), float64(y))))
}
