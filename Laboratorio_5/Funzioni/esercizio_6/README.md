# Testo invertito

Scrivere un programma che legga da **standard input** un testo formato da un numero variabile di righe, terminando la lettura quando viene inserita da **standard input** una riga vuota (`""`), e lo ristampi dall’ultimo carattere al primo.

Scrivere un programma che: 
* legga da **standard input** un testo su più righe;
* termini la lettura quando viene inserita da **standard input** una riga vuota (`""`);
* ristampi il testo letto (riga vuota esclusa) dall’ultimo carattere al primo.

Il programma deve essere organizzato nelle seguenti funzioni:
* una funzione `LeggiTesto() string` che legge da **standard input** un testo su più righe e terminato da una riga vuota (`""`), restituendo un valore (di tipo) `string` (una stringa) in cui è memorizzato il testo letto (riga vuota esclusa);
* una funzione `InvertiStringa(s string) string` che riceve in input un valore (di tipo) `string` (una stringa) nel parametro `s` e restituisce un valore (di tipo) `string` (una stringa) in cui il primo carattere è l'ultimo che definisce `s`, il secondo carattere è il penultimo che definisce `s`, ...  e l'ultimo carattere è il primo che definisce `s`;
* una funzione `main()` che richiama la funzione `LeggiTesto()` per leggere il testo da **standard input**, inverte il testo letto (inverte la stringa in cui è memorizzato il testo letto) richiamando opportunamente la funzione `InvertiStringa()` e stampa il testo invertito (la stringa invertita).

*Suggerimento:* Per leggere da **standard input** un testo formato da un numero variabile di righe (dove ogni riga è una stringa definita da una sequenza di caratteri arbitrari e terminata da '\n') terminando la lettura quando viene inserita da **standard input** una riga vuota (`""`), è possibile utilizzare il seguente programma.

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	stringa := ""
	fmt.Println("Inserisci una o più righe di testo:")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {

		stringaLocale := scanner.Text()
		if stringaLocale == "" {
			fmt.Print("stringaLocale == \"\": lettura terminata.\n")
			break
		} else {
			fmt.Print("stringaLocale == ",stringaLocale,"\n")
			stringa += stringaLocale + "\n"
		}

	}
	fmt.Println("Righe di testo lette: ")
	fmt.Print(stringa)
}
```

##### Esempio d'esecuzione:

```text
$ go run stringainvertita.go
Inserisci un testo su più righe (termina con riga vuota):
Testo di prova
disposto su due righe

Testo invertito:
ehgir eud us otsopsid
avorp id otseT
```