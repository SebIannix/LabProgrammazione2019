/**
Convertitore

Permette di scegliere tra varie conversioni di grandezze numeriche
*/
package main

import "fmt"

func main() {
	var valore float64
	var scelta int = 0

	fmt.Print(
		"Scegli la conversione:\n" +
			"1) secondi -> ore\n" +
			"2) secondi -> minuti\n" +
			"3) minuti -> ore\n" +
			"4) minuti -> secondi\n" +
			"5) ore -> secondi\n" +
			"6) ore -> minuti\n" +
			"7) minuti -> giorni e ore\n" +
			"8) minuti -> anni e giorni\n: ")
	fmt.Scan(&scelta)

	if scelta >= 1 && scelta <= 8 {
		fmt.Print("Inserisci il valore da convertire: ")
		fmt.Scan(&valore)

		// in base alla scelta effettuata, si calcola la conversione
		if scelta == 1 {
			ore := valore / (60 * 60)
			fmt.Println(valore, "secondi corrispondono a", ore, "ore")
		} else if scelta == 2 {
			minuti := valore / 60
			fmt.Println(valore, "secondi corrispondono a", minuti, "minuti")
		} else if scelta == 3 {
			fmt.Println(valore, "minuti corrispondono a", valore/60, "ore")
		} else if scelta == 4 {
			fmt.Println(valore, "minuti corrispondono a", valore*60, "secondi")
		} else if scelta == 5 {
			fmt.Println(valore, "ore corrispondono a", valore*3600, "secondi")
		} else if scelta == 6 {
			fmt.Println(valore, "ore corrispondono a", valore*60, "minuti")
		} else if scelta == 7 {
			giorni := int(valore) / (60 * 24)
			ore := valore/60 - float64(24*giorni)
			fmt.Println(valore, "minuti corrispondono a", giorni, "giorni e", ore, "ore")
		} else if scelta == 8 {
			anni := int(valore) / (60 * 24 * 365)
			giorni := valore/(60*24) - float64(365*anni)
			fmt.Println(valore, "minuti corrispondono a", anni, "anni e", giorni, "giorni")
		}
	} else {
		fmt.Println("Scelta errata")
	}
}
