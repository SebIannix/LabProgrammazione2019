package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {

	testo := LeggiTesto()
	testoTrasformato := Garibaldi(testo)

	fmt.Println("Risultato trasformazione:")
	fmt.Println(testoTrasformato)

}

func LeggiTesto() (testo string) {
	fmt.Println("Inserisci un testo su più righe (termina con riga vuota):")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		riga := scanner.Text()
		if riga != "" {
			testo += riga + "\n"
		} else {
			break
		}
	}
	return testo[:len(testo)-1]
}

func TrasformaCarattere(carattere rune) rune {
	// se il carattere è una vocale
	if strings.ContainsRune("aeioAEIO", carattere) {
		// se è una vocale maiucola ritorno una u minuscola
		// altrimenti una U maiuscola
		if unicode.IsLower(carattere) {
			return 'u'
		} else {
			return 'U'
		}
	} else {
		return carattere
	}
}

func Garibaldi(testo string) (testoTrasformato string) {
	for _, l := range testo {
		testoTrasformato += string(TrasformaCarattere(l))
	}
	return
}
