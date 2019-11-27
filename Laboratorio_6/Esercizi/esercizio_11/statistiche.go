package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	numeroParole, lunghezza := StatisticheParole(LeggiTesto())

	fmt.Println("Statistiche:")
	fmt.Println("Numero parole:", numeroParole)
	fmt.Println("Lunghezza media:", float64(numeroParole)/float64(lunghezza))
}

func LeggiTesto() (testo string) {
	fmt.Println("Inserisci un testo su più righe (termina con Ctrl+D):")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		testo += scanner.Text() + "\n"
	}
	return
}

func StatisticheParole(testo string) (numeroParole int, lunghezzaTotale int) {
	ultimoCarattereLettera := false
	for _, carattere := range testo {

		if unicode.IsLetter(carattere) {
			lunghezzaTotale++
			if !ultimoCarattereLettera {
				numeroParole++
			}
			ultimoCarattereLettera = true
		} else {
			ultimoCarattereLettera = false
		}

	}
	return
}