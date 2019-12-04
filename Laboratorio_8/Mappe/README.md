# Laboratorio 8 - Mappe
## 1 Qual è l'otuput?

Qual è l'output di questo programma?

```go
package main

import "fmt"

func main() {
	var m map[string]int
	if m == nil {
		fmt.Println("'nil' è lo zero-value di una variabile di tipo " +
			"'reference type' non inizializzata.")
	}
	//m["mamma"] = 5 /* panic: assignment to entry in nil map */
	m = make(map[string]int)
	for _, s := range []string{"questo", "è", "un", "test"} {
		m[s] = len([]rune(s))
	}
	for k, v := range m {
		fmt.Println(k, "->", v)
	}
}
```

## 2 Qual è l'output?

Qual è l'output di questo programma?

```go
package main

import "fmt"

func main() {
	mappa := make(map[string]int)
	// equivalente a: mappa := map[string]int{}

	mappa["A"] = 10
	mappa["B"] -= 5
	mappa["D"] = mappa["D"] + 5
	if v, ok := mappa["B"]; ok {
		fmt.Printf("B è presente con valore %d\n", v)
	} else {
		fmt.Print("B non è presente\n")
	}
	if v, ok := mappa["C"]; ok {
		fmt.Printf("C è presente con valore %d\n", v)
	} else {
		fmt.Print("C non è presente\n")
	}
	if v, ok := mappa["C"]; ok {
		fmt.Printf("C è presente con valore %d\n", v)
	} else {
		fmt.Print("C non è presente\n")
	}
	delete(mappa, "B")
	for k, v := range mappa {
		fmt.Printf("chiave %s valore %d\n", k, v)
	}
}
```

## 3 Istogramma a barre orizzontali (1)

Scrivere un programma che: 
1. legga da **standard input** un testo su più righe (alcune delle quali possono essere delle righe vuote (`""`));
2. termini la lettura quando, premendo la combinazione di tasti `Ctrl+D`, viene inserito da **standard input** l'indicatore End-Of-File (EOF);
3. come mostrato nell'**Esempio di esecuzione**, stampi un istogramma a barre orizzontali per rappresentare il numero di occorrenze di ogni lettera presente nel testo letto:
* una lettera è un carattere il cui codice Unicode, se passato come argomento alla funzione `func IsLetter(r rune) bool` del package `unicode`, fa restituire `true` alla funzione;
* le lettere minuscole sono da considerarsi diverse dalle lettere maiuscole; 
* ogni barra viene rappresentata utilizzando il carattere asterisco (`*`); se il numero di occorrenze della lettera `e` è per esempio `9`, la barra corrispondente sarà formata da `9` caratteri `*`.    

Oltre alla funzione `main()`, devono essere definite ed utilizzate almeno le seguenti funzioni:
* una funzione `LeggiTesto() string` che legge da **standard input** un testo su più righe (alcune delle quali possono essere delle righe vuote (`""`)) e terminato dall'indicatore EOF, restituendo un valore `string` in cui è memorizzato il testo letto;
* una funzione `Istogramma(s string) map[rune]int`  che riceve in input un valore `string` nel parametro `s` e restituisce un valore `map[rune]int` in cui, per ogni lettera presente in `s`, è memorizzato il numero di occorrenze della lettera in `s`.

##### Esempio d'esecuzione:

```text
$ go run istogrammaV1.go
TESTO di prova
disposto su più righe!
Istogramma:
i: ****
a: *
h: *
d: **
r: **
g: *
e: *
p: ***
s: ***
t: *
u: *
o: ***
E: *
S: *
O: *
v: *
ù: *
T: **
```
## 4 Istogramma a barre orizzontali (2)

Scrivere un programma che: 
1. legga da **standard input** un testo su più righe (alcune delle quali possono essere delle righe vuote (`""`));
2. termini la lettura quando, premendo la combinazione di tasti `Ctrl+D`, viene inserito da **standard input** l'indicatore End-Of-File (EOF);
3. come mostrato nell'**Esempio di esecuzione**, stampi un istogramma a barre orizzontali per rappresentare il numero di occorrenze di ogni lettera presente nel testo letto:
* una lettera è un carattere il cui codice Unicode, se passato come argomento alla funzione `func IsLetter(r rune) bool` del package `unicode`, fa restituire `true` alla funzione;
* le lettere minuscole sono da considerarsi diverse dalle lettere maiuscole; 
* ogni barra viene rappresentata utilizzando il carattere asterisco (`*`); se il numero di occorrenze della lettera `e` è per esempio `9`, la barra corrispondente sarà formata da `9` caratteri `*`;
* le barre devono essere stampate a partire da quella associata alla lettera con codice Unicode più piccolo fino a quella associata alla lettera con codice Unicode più grande.    

*Suggerimento*: 
* Si consideri il seguente programma.

```go
package main

import (
	"fmt"
)

func main() {

	capitali := map[string]string{"Francia": "Parigi", "Italia": "Roma", 
        "Giappone": "Tokio", "Austria": "Vienna"}

	stati := []string{"Austria", "Francia", "Giappone", "Italia"}

	for _, v := range stati {
		fmt.Println(capitali[v])
	}
}
```
Output:
```text 
Vienna
Parigi
Tokio
Roma
```
Il programma stampa a video i nomi delle capitali europee memorizzati in `capitali` in base all'ordine in cui i nomi degli stati corrispondenti sono memorizzati in `stati`.
* La seguente funzione ordina una slice di rune in ordine crescente.

```go
func SortRunes(a []rune) {
	for i:=0;i<len(a)-1;i++{
		indiceMin := i
		for j := i + 1; j<len(a); j++ {
			if a[indiceMin] > a[j] {
				indiceMin = j
			}
		}
		a[i], a[indiceMin] = a[indiceMin], a[i]
	}
}
```

##### Esempio d'esecuzione:

```text
$ go run istogramma.go
Ciao,
come stai?
Occorrenze:
C: *
a: **
c: *
e: *
i: **
m: *
o: **
s: *
t: *
```
## 5 Ripetizioni

Scrivere un programma che:
* legga da **standard input** un testo su più righe (alcune delle quali possono essere delle righe vuote (`""`));
* termini la lettura quando, premendo la combinazione di tasti `Ctrl+D`, viene inserito da **standard input** l'indicatore End-Of-File (EOF);
* stampi a video le seguenti informazioni relative al testo letto:
  1. Il numero di parole distinte presenti nel testo (una parola è una stringa interamente definita da caratteri il cui codice Unicode, se passato come argomento alla funzione `func IsLetter(r rune) bool`, fa restituire `true` alla funzione).
  2. La lista di parole distinte presenti nel testo, riportando per ogni parola il relativo numero di occorrenze nel testo (cfr. **Esecuzione d'esecuzione**).

Il programma deve essere organizzato nelle seguenti funzioni:
* una funzione `LeggiTesto() string` che legge da **standard input** un testo su più righe (alcune delle quali possono essere delle righe vuote (`""`)) e terminato dall'indicatore EOF, restituendo un valore `string` in cui è memorizzato il testo letto;
* una funzione `SeparaParole(s string) []string` che riceve in input un valore `string` nel parametro `s` e restituisce un valore `[]string` in cui sono memorizzate tutte le parole presenti in `s`;
* una funzione `ContaRipetizioni(sl []string) map[string]int` che riceve in input un valore `[]string` nel parametro `sl` e restituisce un valore `map[string]int` in cui, per ogni parola presente in `sl`, è memorizzato il numero di occorrenze della parola in `sl`.

##### Esempio d'esecuzione:

```text
$ go run ripetizioni.go
Ciao come stai?
io sto bene, tu?            
anche io sto bene, grazie
Parole distinte: 9
io: 2
sto: 2
tu: 1
anche: 1
come: 1
stai: 1
bene: 2
grazie: 1
Ciao: 1
```
