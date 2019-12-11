# Laboratorio 9 - Strutture e puntatori
## 1 Qual è l'output?

Qual è l'output del seguente programma?

```go
package main

import "fmt"

func main() {
	var a, b, c int
	var ptr *int
	a, b, c = 10, 15, 20
	ptr = &a
	*ptr += 10
	a += 10
	ptr = &b
	*ptr += 10
	*ptr = c
	*ptr += 10
	fmt.Println(a, b, c)
}
```

## 2 Qual è l'output?

Qual è l'output di questo programma?

```go
package main

import "fmt"

func main() {
	var a, b = 15, 20
	f(&a, &b)
	fmt.Println(a, b)
}
func f(x, y *int) {
	*x, *y = *y, *x
}
```

## 3 Trova l'errore

Questo programma dovrebbe stampare `20 100` ma non genera l'output desiderato. Corregere e verificare che l'esecuzione del programma generi l'output atteso.

```go
package main

import "fmt"

func main() {
	var a, b int = 10, 20
	var ptr1, ptr2 *int
	ptr1 = a
	*ptr1 += 10
	*ptr2 = 100
	ptr2 = &b
	fmt.Println(a, b)
}
```

## 4 Qual è l'output?

Qual è l'output di questo programma?

```go
package main

import (
	"fmt"
	"strings"
)

type Persona struct {
	Nome, Cognome string
}

func main() {

	p1 := Persona{"Rick", "Sanchez"}

	fmt.Println("Stampe per p1:")
	fmt.Println(p1)
	f(p1)
	fmt.Println(p1)
	g(&p1)
	fmt.Println(p1)

	p2 := NuovaPersona("Mario", "Rossi")

	fmt.Println("\nStampe per p2:")
	fmt.Println(p2)
	f(*p2)
	fmt.Println(p2)
	g(p2)
	fmt.Println(p2)
}

func f(p Persona) {
	p.Nome, p.Cognome = strings.ToUpper(p.Nome), strings.ToUpper(p.Cognome)
}

func g(p *Persona) {
	p.Nome, p.Cognome = strings.ToUpper(p.Nome), strings.ToUpper(p.Cognome)
	// ... equivalente a:
	//(*p).Nome, (*p).Cognome = strings.ToUpper((*p).Nome), strings.ToUpper((*p).Cognome)

}

// Factory method of type Persona
func NuovaPersona(nome, cognome string) (p *Persona) {

	if p == nil {
		fmt.Println("'nil' è lo zero-value di una variabile di tipo " +
			"'reference type' non inizializzata.")
	}
	p = new(Persona) // ... equivalente a: p = &Persona{}
	p.Nome = nome
	p.Cognome = cognome
	return p
	// ... equivalente a:
	//p = &Persona{Nome:nome,Cognome:cognome}
	//return p
	// ... equivalente a:
	//return &Persona{nome, cognome}
}
```

## 5 Frazioni

Al fine di modellare l'entità matematica `frazione`, si definisca un nuovo tipo `Frazione` come una struttura avente due campi `Numeratore` e `Denominatore` di tipo `int`.

Implementare le funzioni:

* `NuovaFrazione(numeratore, denominatore int) *Frazione` che restituisce una nuova istanza del tipo `Frazione` inizializzata in base ai valori dei parametri `numeratore` e `denominatore`;
* `String(f Frazione) string` che riceve in input un'instanza del tipo `Frazione` nel parametro `f` e restituisce un valore `string` che corrisponde alla rappresentazione  `string` di `f` nel formato `NUMERATORE/DENOMINATORE`, dove `NUMERATORE` (`DENOMINATORE`) è il valore `string` corrispondente al valore del campo `Numeratore` (`Denominatore`) di `f`;

affinché la funzione `main()` di seguito riportata possa essere eseguita generando l'output atteso.

Funzione `main()`:
```go
func main() {
	frazione := NuovaFrazione(34, 18)
	fmt.Println(String(*frazione))
}
```  
Output atteso:
```text
34/18
```
## 6 Riduzione di una frazione ai minimi termini

Sulla base del codice sviluppato relativamente all'Esercizio 2 (Laboratorio 9 - Strutture II), scrivere un programma che:
* legga da **riga di comando**  due numeri interi `n` e `d`;
* riduca ai minimi termini la frazione che ha per numeratore e denominatore i valori `n` e `d`; 
* stampi a video, nel formato `NUMERATORE/DENOMINATORE`, la frazione ridotta ai minimi termini.

Oltre alla funzione `main()`, devono essere definite ed utilizzate almeno le seguenti funzioni:
* una funzione `Riduci(f *Frazione)` che riceve in input un'instanza del tipo `Frazione` nel parametro `f` e, se necessario, modifica opportunamente il valore dei campi `f.numeratore` e `f.denominatore` affinché `f` rappresenti una frazione ridotta ai minimi termini.

##### Esempio d'esecuzione:
```text
$ go run riduci.go 10 10
1/1

$ go run riduci.go 34 18
17/9

$ go run riduci.go 12 36
1/3
```
## 7 Moltiplicazione tra frazioni

Sulla base del codice sviluppato relativamente all'Esercizio 2 e 3 (Laboratorio 9 - Strutture II), scrivere un programma che:
* legga da **standard input** un testo su più righe;
* termini la lettura quando viene inserita da standard input una riga vuota (`""`).

In ogni riga sono specificati due numeri interi che rappresentano il numeratore e il denominatore di una frazione.

Una volta terminata la fase di lettura, il programma deve:
1. calcolare la frazione che si ottiene moltiplicando tra di loro le frazioni corrispondenti alle coppie di interi letti;
2. ridurre ai minimi termini la frazione calcolata al punto 1;
3. stampare a video, nel formato `NUMERATORE/DENOMINATORE`, la frazione ridotta ai minimi termini ottenuta al punto 2.

Oltre alla funzione `main()`, devono essere definite ed utilizzate almeno le seguenti funzioni:

* una funzione `LeggiFrazioni() []Frazione` che legge da **standard input**  un testo su più righe e terminato da una riga vuota (`""`), restituendo un valore `[]Frazione` in cui sono memorizzate tutte le istanze del tipo `Frazione` inizializzate in base ai numeri interi `numeratore` e `denominatore` specificati in ciascuna delle righe lette;
* una funzione `Moltiplica(f1, f2 Frazione) *Frazione` che riceve in input due istanze del tipo `Frazione` nei parametri `f1` e `f2` e restituisce una nuova istanza del tipo `Frazione` i cui campi `numeratore` e `denominatore` sono inizializzati con i valori  `f1.numeratore * f2.numeratore` e `f1.denominatore * f2.denominatore`;
* una funzione `MoltiplicaN(fN []Frazione) *Frazione` che riceve in input un valore `[]Frazione` nel parametro `fN` e restituisce una nuova istanza del tipo `Frazione` corrispondente alla frazione che si ottiene moltiplicando tra di loro le frazioni relative alle istanze del tipo `Frazione` presenti in `fN`; la funzione deve utilizzare la funzione `Moltiplica()`.

##### Esempio d'esecuzione:
```text
$ go run prodotto.go 
Inserisci numeratore e denominatore delle frazioni:
1 2
4 3
2 5

4/15

$ go run prodotto.go
Inserisci numeratore e denominatore delle frazioni:
1 1
2 3
1 8

1/12
```

