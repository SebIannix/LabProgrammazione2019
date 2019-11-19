package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	testo := LeggiTesto()
	testoSpaziato := Spazia(testo)

	fmt.Println("Risultato:")
	fmt.Println(testoSpaziato)
}

func LeggiTesto() (testo string) {
	fmt.Println("Inserisci un testo su più righe (termina con CTRL D):")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		testo += scanner.Text() + "\n"
	}
	return
}

func Spazia(testo string) (testoTrasformato string) {
	caratterePrecedente := ' '
	// aggiungo uno spazio solo quando sia il carattere precedente che 
	// quello attuale non sono spazi
	for _, carattere := range testo {
		if !unicode.IsSpace(carattere) && !unicode.IsSpace(caratterePrecedente) {
			testoTrasformato += " " + string(carattere)
		} else {
			testoTrasformato += string(carattere)
		}
		caratterePrecedente = carattere
	}
	return
}
