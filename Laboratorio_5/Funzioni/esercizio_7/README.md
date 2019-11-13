# Linguaggio farfallino

Nel linguaggio farfallino ciascuna vocale non accentata (`vocale`) viene sostituita da una sequenza di tre caratteri `vocale-f-vocale`. Per esempio, la vocale `a` viene sostituita dalla sequenza `afa`, la vocale `e` dalla sequenza `efe` e così via. Se una vocale è maiuscola, anche la sequenza di tre caratteri che sostituisce la vocale deve essere definita da caratteri maiuscoli (ad esempio, la vocale `A` viene sostituita dalla sequenza `AFA`).

Scrivere un programma che: 
* legga da **standard input** un testo su più righe (alcune delle quali possono essere delle righe vuote (`""`));
* termini la lettura quando, premendo la combinazione di tasti `Ctrl+D`, viene inserito da **standard input** l'indicatore End-Of-File (EOF);
* ristampi il testo letto dopo averlo tradotto in linguaggio farfallino.

Il programma deve essere organizzato nelle seguenti funzioni:
* una funzione `LeggiTesto() string` che legge da **standard input** un testo su più righe (alcune delle quali possono essere delle righe vuote (`""`)) e terminato dall'indicatore EOF, restituendo un valore `string` in cui è memorizzato il testo letto;

* una funzione `TraduciCarattereInFarfallino(c rune) string` che riceve in input un valore `rune` nel parametro `c` e restituisce un valore `string` che rappresenta la traduzione in linguaggio farfallino di `c`;

* una funzione `TraduciTestoInFarfallino(s string) string` che riceve in input un valore `string` nel parametro `s` e restituisce un valore `string` che rappresenta la traduzione in linguaggio farfallino di `s`; la funzione deve utilizzare la funzione `TraduciCarattereInFarfallino()`;
* una funzione `main()` che richiama la funzione `LeggiTesto()` per leggere il testo da **standard input**, traduce in linguaggio farfallino il testo letto (richiamando opportunamente la funzione `TraduciTestoInFarfallino()`) e stampa il testo tradotto.

*Suggerimento:* Per leggere da **standard input** un testo formato da un numero variabile di righe (dove ogni riga è una stringa definita da una sequenza di caratteri arbitrari e terminata da '\n') terminando la lettura quando viene inserito da **standard input** l'indicatore End-Of-File (EOF), è possibile utilizzare il seguente programma.

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
		fmt.Print("stringaLocale == ",stringaLocale,"\n")
		stringa += stringaLocale + "\n"

	}
	fmt.Println("Righe di testo lette: ")
	fmt.Print(stringa)
}
```

##### Esempio d'esecuzione:

```text
$ go run farfallino.go 
Inserisci un testo su più righe (termina con Ctrl+D):
Questo e' un testo di prova
da trasformare IN ALFABETO FARFALLINO
Risultato:
Qufuefestofo efe' ufun tefestofo difi profovafa
dafa trafasfoformafarefe IFIN AFALFAFABEFETOFO FAFARFAFALLIFINOFO
```
