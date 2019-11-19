package main

import (
	"bufio"
	"fmt"
	"os"
//	"strings" // per versione alternativa
)

func main() {

	testo := LeggiTesto()
	testoTrasformato := TraduciTestoInFarfallino(testo)

	fmt.Println("Risultato:")
	fmt.Println(testoTrasformato)

}

func LeggiTesto() (testo string) {
	fmt.Println("Inserisci un testo su più righe (termina con CTRL D):")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		testo += scanner.Text() + "\n"
	}
	return
}

func TraduciCarattereInFarfallino(carattere rune) (carattereTrasformato string) {
	carattereTrasformato = string(carattere)
	
	if carattere=='a' || carattere=='e' || carattere=='i' || carattere=='o' || carattere=='u' {
		carattereTrasformato += "f" + carattereTrasformato
	}
	if carattere=='A' || carattere=='E' || carattere=='I' || carattere=='O' || carattere=='U' {
		carattereTrasformato += "F" + carattereTrasformato
	}
	/*
	VERSIONE ALTERNATIVA: controllo se il carattere è presente
	in una stringa contenente solo vocali
	if strings.ContainsRune("aeiou", carattere) {
		carattereTrasformato += "f" + carattereTrasformato
	}
	if strings.ContainsRune("AEIOU", carattere) {
		carattereTrasformato += "F" + carattereTrasformato
	}
	*/
	return
}

func TraduciTestoInFarfallino(testo string) (testoTrasformato string) {
	for _, carattere := range testo {
		testoTrasformato += TraduciCarattereInFarfallino(carattere)
	}
	return
}
