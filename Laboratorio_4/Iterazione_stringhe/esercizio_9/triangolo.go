package main

import "fmt"

func main() {
	var dimensione int
	fmt.Scan(&dimensione)

	// Versione 1:
	// uso due cicli annidati, uno per le righe e uno per le colonne

	if dimensione <= 0 {
		fmt.Println("Dimensione non sufficiente")
	} else {

		// un numero di righe pari alla dimensione letta in input
		for riga := 0; riga < dimensione; riga++ {
			// il numero di colonne dipende dalla riga in cui siamo:
			// alla prima riga dobbiamo stampare 1 *
			// alla seconda riga 2 *
			// alla terza riga 3 *
			// e così via
			// c'è quindi una relazione tra il numero di riga e il numero di asterischi
			// che dobbiamo stampare
			for colonna := 0; colonna <= riga; colonna++ {
				fmt.Print("*")
			}
			fmt.Println()
		}
	}

}
