package main

import (
	"fmt"
	"strconv"
)

func main() {

  // ho diviso il programma in 4 funzioni principali:
  // - due per leggere i comandi
  // - una per stampare gli spostamenti totali
  // - una per stampare i comandi in ordine inverso

	comandi := LeggiComandi(LeggiNumeroComandi())

	fmt.Println()

	StampaSpostamenti(comandi)
	StampaComandiInversi(comandi)
}

// La funzione legge il numero di comandi che dovranno essere inseriti
func LeggiNumeroComandi() (n int) {
	fmt.Println("Inserisci il numero di comandi:")
	fmt.Scan(&n)
	return
}

// Legge i comandi e li memorizza in modo sequenziale in una stringa
func LeggiComandi(numeroComandi int) (comandi string) {
	fmt.Println("Inserisci i comandi:")
	for i:=0; i<numeroComandi; i++ {
    // Per leggere direzione e numero di passi dichiaro due variabili string:
    // potrei dichiarare passi come int, ma poi lo dovrei comunque convertire in string
		var direzione, passi string
		fmt.Scan(&direzione, &passi)
		comandi += direzione + passi
	}
	return
}

// Stampa gli spostamenti totali in ogni direzione.
// Riceve in input la stringa di comandi.
func StampaSpostamenti(comandi string) {
  // contatori dei passi totali
	var nord, ovest, est, sud int

  // dato che ogni comando occupa due byte della stringa (direzione e passi)
  // l'indice viene incrementato di due ad ogni iterazione
	for i:=0; i<len(comandi); i+=2 {
		direzione := comandi[i] // direzione
		passi, _ := strconv.Atoi(string(comandi[i+1])) // passi
		switch direzione {
		case 'N': nord += passi
		case 'O': ovest += passi
		case 'E': est += passi
		case 'S': sud += passi
		}
	}

	fmt.Println("Movimenti totali:")
	if nord>0 {
		fmt.Println("N", nord)
	}
	if sud>0 {
		fmt.Println("S", sud)
	}
	if est>0 {
		fmt.Println("E", est)
	}
	if ovest>0 {
		fmt.Println("O", ovest)
	}
}

// Stampa i comandi per riportare il robot alla posizione originale
func StampaComandiInversi(comandi string) {
	fmt.Println("Comandi inversi:")
  // scorro la stringa di comandi al contrario,
  // dall'ultimo comando al primo
	for i:=len(comandi)-2; i>=0; i-=2 {
		direzione := comandi[i] // direzione
		passi := string(comandi[i+1]) // passi
		switch direzione { // cambio il valore della direzione con la direzione opposta
		case 'N': direzione = 'S'
		case 'O': direzione = 'E'
		case 'E': direzione = 'O'
		case 'S': direzione = 'N'
		}
		fmt.Print(string(direzione), " ", passi)
		if i>0 {
			fmt.Print(", ")
		}
	}
	fmt.Println()
}
