package main

import "fmt"

func main() {

	// la dimensione del tronco è data e non è un input
	// ma uso delle costanti per rendere più leggibile e manutenibile il codice
	const larghezzaTronco = 3
	const altezzaTronco = 3

	var dimensione int
	fmt.Scan(&dimensione)

	// calcolo quanto deve essere largo l'albero
	larghezzaAlbero := 2*dimensione - 1

	// se la chioma è più larga del tronco, allora il tronco dovrà essere traslato verso destra
	// e calcolo di quanti caratteri ' ' (spazio) deve essere traslato
	var traslazioneTronco int
	if larghezzaAlbero > larghezzaTronco {
		traslazioneTronco = (larghezzaAlbero - larghezzaTronco) / 2
	}

	// se il tronco è più largo della chioma, allora la chioma dovrà essere traslata verso destra
	var traslazioneAlbero int
	if larghezzaTronco > larghezzaAlbero {
		traslazioneAlbero = (larghezzaTronco - larghezzaAlbero) / 2
	}

	// Stampa chioma dell'albero
	for riga := 0; riga < dimensione; riga++ {
		// per disegnare la chioma devo spostarmi di un numero di spazi che dipende da:
		// - la riga della chioma
		// - la traslazione della chioma (se più piccola del tronco)
		for colonna := 0; colonna < dimensione-1-riga+traslazioneAlbero; colonna++ {
			fmt.Print(" ")
		}

		for colonna := 0; colonna < 2*riga+1; colonna++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	// Stampa tronco
	for i := 0; i < altezzaTronco; i++ {
		// per disegnare il tronco devo prima spostarmi di un numero di spazi che dipende
		// dalla traslazione del tronco (se di larghezza inferiore alla chioma)
		for j := 0; j < traslazioneTronco; j++ {
			fmt.Print(" ")
		}

		for j := 0; j < larghezzaTronco; j++ {
			fmt.Print("+")
		}
		fmt.Println()
	}

}
