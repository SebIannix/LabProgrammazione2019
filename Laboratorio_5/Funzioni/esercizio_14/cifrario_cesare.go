package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {

	chiave := LeggiNumero()
	testo := LeggiTesto()

	testoCifrato := CifraTesto(testo, chiave)

	fmt.Println("Testo cifrato:")
	fmt.Println(testoCifrato)
}

func LeggiTesto() (testo string) {
	fmt.Println("Inserisci un testo su più righe (termina con CTRL D):")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		testo += scanner.Text() + "\n"
	}
	return
}

func LeggiNumero() (n int) {
	fmt.Print("Inserisci un numero: ")
	fmt.Scan(&n)
	return
}

func CifraCarattere(carattere rune, chiave int) rune {

	// la cifratura di Cesare deve essere circolare,
	// e deve ricominciare dalla A nel caso si vada oltre la Z
	// Le lettere dell'alfabeto inglese sono 26, ma se non lo ricordassi
	// potrei calcolare questo numero come differenza tra l'ultima lettere
	// e la prima lettera dell'alfabeto (dato che sono codici numerici)
	const rangeLettere = 'Z'-'A'+1

	// se la chiave è superiore a 26 ne faccio il modulo
	// così da portarla ad avere valori tra 0 e 25:
	// infatti, usando una chiave con valore 27 si ottiene lo stesso risultato 
	// che si otterrebbe con una chiave di valore 1
	chiavePositiva := rune(chiave%rangeLettere)
	// nel caso la chiave fosse negativa, la trasformo in positiva:
	// infatti, una chiave di valore -1 da lo stesso risultato
	// di una chiave di valore 25
	if chiavePositiva < 0 {
		chiavePositiva = rangeLettere + chiavePositiva
	}

	// se il carattere in input è una lettera dell'alfabeto inglese
	if (carattere>='A' && carattere<='Z') || carattere>='a' && carattere<='z' {
		
		offset := 'A'
		if unicode.IsLower(carattere) {
			offset = 'a'
		}

		// dal valore del carattere tolgo il valore della lettera 'a' (maiuscola o minuscola)
		// ottenendo così la posizione occupata dal carattere nell'alfabeto inglese 
		// (da 0 a 25: se carattere fosse uguale a 'B', facendo 'B'-'A' otterrei 1).
		// Al valore ottenuto aggiungo il valore della chiave e uso il modulo per 
		// far rientrare il valore all'interno dell'intervallo delle lettere dell'alfabeto.
		// Infine, aggiungo il valore della lettera 'a' (maiuscola o minuscola) per
		// ottenere di nuovo il valore di una lettera secondo codifica ASCII
		return (((carattere - offset) + chiavePositiva) % rangeLettere) + (offset)
	} else {
		return carattere
	}
}

func CifraTesto(testo string, chiave int) (testoCifrato string) {
	for _, carattere := range testo {
		testoCifrato += string(CifraCarattere(carattere, chiave))
	}
	return
}


