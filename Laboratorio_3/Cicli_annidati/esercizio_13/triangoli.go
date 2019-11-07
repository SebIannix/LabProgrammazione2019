package main

import (
	"fmt"
)

func main() {

	var dimensione int
	fmt.Scan(&dimensione)

	// triangolo superiore
	// possiamo immaginare il triangolo superiore come un quadrato in cui la diagonale, il bordo superiore
	// e il bordo di destra sono stampati con un carattere *, mentre tutto il resto
	// è stampato con un carattere ' ' (spazio)
	for riga := 0; riga < dimensione; riga++ {
		for colonna := 0; colonna < dimensione; colonna++ {
			if riga == 0 || colonna == dimensione-1 || riga == colonna {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

	// triangolo inferiore
	// possiamo immaginare il triangolo inferiore come un quadrato in cui la diagonale, il bordo inferiore
	// e il bordo di sinistra sono stampati con un carattere *, mentre tutto il resto
	// è stampato con un carattere ' ' (spazio)
	// il bordo inferiore però deve essere anche traslato verso destra: possiamo traslare una figura stampando
	// degli spazi aggiuntivi in ogni riga
	for riga := 0; riga < dimensione; riga++ {
		// stampo gli spazi per traslare il triangolo verso destra
		for colonna := 0; colonna < dimensione-1; colonna++ {
			fmt.Print(" ")
		}

		// stampo una riga del triangolo inferiore
		for j := 0; j < dimensione; j++ {
			if riga == dimensione-1 || j == 0 || riga == j {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
