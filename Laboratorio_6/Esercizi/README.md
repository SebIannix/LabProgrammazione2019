# Laboratorio 6 - Esercizi
## 1 Qual è l'output?

Qual è l'output del seguente programma?

```go
package main

import (
	"fmt"
)

const Chiave = 5
const NumeroLettere = 26

func main() {
	s := "Guarda, è Zorro!"

	for _, carattere := range s {

		switch {
		case carattere >= 'A' && carattere <= 'Z':
			carattere = 'A' + ((carattere - 'A' + Chiave) % NumeroLettere)
		case carattere >= 'a' && carattere <= 'z':
			carattere = 'a' + ((carattere - 'a' + Chiave) % NumeroLettere)
		}

		fmt.Printf("%c",carattere)

	}

	fmt.Println()

	fmt.Println(s)
}
```
## 2 Qual è l'output?

Qual è l'output del seguente programma?

```go
package main

import (
	"fmt"
)

func main() {

	s := "Papà!"

	var s1 string
	for _, carattere := range s {
		s1 += string(carattere)
	}
	fmt.Println(s1)

	var s2 string
	for _, carattere := range s {
		s2 = string(carattere) + s2
	}
	fmt.Println(s2)

}
```
## 3 Qual è l'output?

Supponendo che l'utente inserisca da **standard input** il valore `5`, cosa dovrebbe produrre in output il seguente programma?

```go
package main

import (
	"fmt"
)

func main() {

	var cifraInInput int32

	for cifraInInput <= 0 || cifraInInput >= 10 {
		fmt.Print("Inserisci un'unica cifra decimale maggiore o uguale a 1: ")
		fmt.Scan(&cifraInInput)
	}
	fmt.Println()

	var carattere rune

	carattere = '0' + (cifraInInput - 1)

	fmt.Printf("Carattere: %c\n", carattere)
    fmt.Printf("Carattere: %d\n", carattere)

	var cifraInOutput int32

	cifraInOutput = (carattere - '0') + 1

	fmt.Printf("Cifra: %d\n", cifraInOutput)
    fmt.Printf("Cifra: %c\n", cifraInOutput)
}
```
## 4 Scacchiera

Scrivere un programma che legga da **standard input** un intero `n>=2` e,  come mostrato nell'**Esempio d'esecuzione**, stampi a video una schacchiera di dimensione `n x n` utilizzando i caratteri `'*'` (asterisco) e `'+'` (più).

##### Esempio d'esecuzione:

```text
$ go run scacchiera.go
6
*+*+*+
+*+*+*
*+*+*+
+*+*+*
*+*+*+
+*+*+*

$ go run scacchiera.go 
3
*+*
+*+
*+*
``` 
## 5 Scala (1)

Scrivere un programma che legga da **standard input** un numero intero `n`  e, come mostrato nell'**Esempio d'esecuzione**, stampi a video una scala utilizzando il carattere `'*'` (asterisco):

* La scala è formata da `n` gradini.
* Ciascun gradino è profondo 3 caratteri e alto 2.
* Il gradino più in alto deve essere stampato senza alcuna rientranza verso destra; considerando poi i successivi gradini dall'alto verso il basso, ciascun di essi rientra (è traslato) verso destra rispetto al precedente di due caratteri `' '` (spazio) .

Se `n` è negativo o nullo, anziché stampare la scala il programma deve stampare un messaggio d'errore.

##### Esempio d'esecuzione:

```text
$ go run scala_1.go
4
***
  *
  ***
    *
    ***
      *
      ***
        *

$ go run scala.go
1
***
  *

$ go run scala.go
0
Dimensione non sufficiente
```
## 6 Scala (2)

Scrivere un programma che legga da **standard input** un numero intero `n`  e, come mostrato nell'**Esempio d'esecuzione**, stampi a video una scala utilizzando il carattere `'*'` (asterisco):

* La scala è formata da `n` gradini.
* Ciascun gradino è profondo 3 caratteri e alto 2.
* Il gradino più in basso deve essere stampato senza alcuna rientranza verso destra; considerando poi i successivi gradini dal basso verso l'alto, ciascun di essi rientra (è traslato) verso destra rispetto al precedente di due caratteri `' '` (spazio) .

Se `n` è negativo o nullo, anziché stampare la scala il programma deve stampare un messaggio d'errore.

##### Esempio d'esecuzione:

```text
$ go run scala.go
3
    ***
    *
  ***
  *
***
*

$ go run scala.go
0
Dimensione non sufficiente

$ go run scala.go
1
***
*
```
## 7 Stringa alternata

Scrivere un programma che legga da **standard input** due stringhe senza spazi `s1` e `s2` e stampi a video la stringa creata alternando i caratteri delle stringhe `s1` e `s2`.

*Esempio:* Se `"ciao!"` e `"MONDO"` sono le stringhe lette, allora la stringa stampata video deve essere `"cMiOaNoD!O"`.

Si assuma che le stringhe lette siano interamente definite da caratteri considerati nello standard US-ASCII.

Se le stringhe lette non sono definite dallo stesso numero di caratteri, si deve utilizzare il carattere '-' come carattere di riempimento:

*Esempio:* Se `"ciao"` e `"mondo!"`, sono le stringhe lette, allora la stringa stampata video deve essere `"cmioanod-o-!"`.

##### Esempio d'esecuzione
```text
$ go run stringa_alternata.go
ciao
mondo
cmioanod-o

$ go run stringa_alternata.go 
ciaone
mondo
cmioanodnoe-

$ go run stringa_alternata.go
esame
go
egsoa-m-e-
```
## 8 Robot

Un robot è in grado di muoversi nelle quattro direzioni nord (N), sud (S), est (E) e ovest (O). 

In particolare, il robot accetta comandi che consistono in una coppia di valori `d p`, dove `d` è un carattere che indica la direzione e può assumere uno dei quattro valori `'N'`, `'S'`, `'E'`, `'O'`, mentre `0 < p < 10` è un intero che indica il numero di passi che il robot deve compiere in quella direzione. 

Scrivere un programma che, dopo aver letto da **standard input** un numero intero `n`, chiede all'utente di inserire `n` comandi da impartire al robot (sempre da **standard input**). 

Dopo aver letto gli `n` comandi, il programma deve stampare a video:
* il numero totale di passi che deve compiere il robot in ognuna delle quattro direzioni (se il robot non si deve mai muovere in una certa direzione, l'output relativo a tale direzione non va stampato);
* la sequenza di comandi inversi che si dovrebbe impartire al robot per farlo ritornare al punto di partenza lungo lo stesso percorso.

*Suggerimento*: La sequenza di comandi:
```text
O 2
O 3 
N 4 
```
può essere memorizzata nella stringa `"O2O3N4"`. In particolare, una qualsiasi sequenza di comandi può essere memorizzata in una stringa interamente definita da caratteri considerati nello standard US-ASCII.

##### Esempio d'esecuzione:

```text
$ go run robot.go
Inserisci il numero di comandi:
5
Inserisci i comandi:
O 2 
O 3 
N 4 
S 5 
E 1 

Movimenti totali:
N 4
S 5
E 1
O 5
Comandi inversi:
O 1, N 5, S 4, E 3, E 2

$ go run robot.go 
Inserisci il numero di comandi:
8
Inserisci i comandi:
O 2  
N 1
E 5 
S 4 
E 5 
N 3 
S 1
O 3 

Movimenti totali:
N 4
S 5
E 10
O 5
Comandi inversi:
E 3, N 1, S 3, O 5, N 4, O 5, S 1, E 2

$ go run robot.go  
Inserisci il numero di comandi:
4
Inserisci i comandi:
E 3
N 1
N 4
E 2

Movimenti totali:
N 5
E 5
Comandi inversi:
O 2, S 4, S 1, O 3

$ go run robot.go 
Inserisci il numero di comandi:
5
Inserisci i comandi:
N 3 
N 2 
O 3 
S 2 
O 1 

Movimenti totali:
N 5
S 2
O 4
Comandi inversi:
E 1, N 2, E 3, S 2, S 3
```

## 9 Numero nascosto

Scrivere un programma che legga da **standard input** una riga di testo e stampi in output il doppio del numero nascosto all'interno della riga di testo, ovvero il doppio del numero che si ottiene concatenando le cifre presenti all'interno della riga di testo.

*Suggerimento:* Per convertire una stringa in un numero intero  si faccia riferimento alla documentazione della funzione  `strconv.Atoi()` del package `strconv`.

##### Esempio d'esecuzione:

```text
$ go run numero_nascosto.go
Ci8ao 97com3 va?
Doppio del numero nascosto: 17946 (8973 * 2)

$ go run numero_nascosto.go
Ch3 831 t3mp0
Doppio del numero nascosto: 766260 (383130 * 2)

$ go run numero_nascosto.go
c140n3
Doppio del numero nascosto: 2806 (1403 * 2)
```
## 10 Numeri abbondanti

**Definizione**: Un numero naturale è abbondante se è inferiore alla somma dei suoi divisori propri (per esempio, `12` è abbondante perché  la somma dei suoi divisori propri (`1`, `2`, `3`, `4`, `6`) è `16`).

Scrivere un programma che legga da **standard input** un numero intero `soglia` e stampi tutti i numeri abbondanti inferiori a `soglia`.

Il programma deve essere organizzato nelle seguenti funzioni:
* una funzione `LeggiNumero() int` che legge da **standard input** un numero intero e ne restituisce il valore;
* una funzione `SommaDivisoriPropri(n int) int` che riceve in input un valore `int` nel parametro `n` e restituisce un valore `int` pari alla somma dei divisori propri di `n` (pari a `0` se `n` non ha divisori propri);
* una funzione `ÈAbbondante(n int) bool` che riceve in input un valore `int` nel parametro `n` e restituisce `true` se `n` è abbondante, `false` altrimenti; la funzione deve utilizzare la funzione `SommaDivisoriPropri()`;
* una funzione `NumeriAbbondanti(limite int)` che riceve in input un valore `int` nel parametro `limite` e stampa tutti i numeri abbondanti inferiori a `limite`; la funzione deve utilizzare la funzione `ÈAbbondante()`;
* una funzione `main()` che richiama la funzione `LeggiNumero()` per leggere il numero intero `soglia`. Se `soglia > 0` allora richiama la funzione `NumeriAbbondanti()` in modo opportuno, altrimenti stampa un messaggio d'errore.

##### Esempio d'esecuzione:

```text
$ go run numeri_abbondanti.go
20
Numeri abbondanti: 12 18

$ go run numeri_abbondanti.go 
10
Numeri abbondanti:


$ go run numeri_abbondanti.go
30
Numeri abbondanti: 12 18 20 24
```
## 11 Statistiche testo

Scrivere un programma che: 
* legga da **standard input** un testo su più righe (alcune delle quali possono essere delle righe vuote (`""`));
* termini la lettura quando, premendo la combinazione di tasti `Ctrl+D`, viene inserito da **standard input** l'indicatore End-Of-File (EOF);
* stampi a video le seguenti statistiche relative al testo letto:
  1. Il numero di parole presenti nel testo (una parola è una stringa interamente definita da caratteri il cui codice Unicode, se passato come argomento alla funzione `func IsLetter(r rune) bool`, fa restituire `true` alla funzione).
  2. La lunghezza media delle parole presenti nel testo (misurata in caratteri).

Il programma deve essere organizzato nelle seguenti funzioni:
* una funzione `LeggiTesto() string` che legge da **standard input** un testo su più righe (alcune delle quali possono essere delle righe vuote (`""`)) e terminato dall'indicatore EOF, restituendo un valore `string` in cui è memorizzato il testo letto;
* una funzione `StatisticheParole(s string) (int, int)` che riceve in input un valore `string` nel parametro `s` e restituisce due valori `int` pari rispettivamente al numero di parole presenti in `s` e alla loro lunghezza totale;
* una funzione `main()` che richiama la funzione `LeggiTesto()` per leggere il testo da **standard input**, calcola le statistiche relative al testo letto (richiamando opportunamente la funzione `StatisticheParole()`) e stampa a video le statistiche calcolate.

##### Esempio d'esecuzione:

```text
$ go run statistiche.go
Inserisci un testo su più righe (termina con Ctrl+D):
Testo di prova 

su cui calcolare le statistiche.
:-)
Statistiche:
Numero parole: 8
Lunghezza media: 0.20512820512820512
```

## 12 Terna pitagorica

**Definizione**: Se `a`, `b` e `c` sono numeri naturali e `a² + b² = c²`, si dice che la terna di numeri `a`, `b` e `c` è una terna pitagorica (il nome deriva dal teorema di Pitagora: In un triangolo rettangolo di cateti `a`, `b` e di ipotenusa `c` si ha che `a² + b² = c²`). 

Scrivere un programma che legga da **standard input** un intero `n>0` e stampi a video tutte le terne pitagorighe tali che `a<n`, `b<n` e `c<n`. 
 
 Oltre alla funzione `main()`, devono essere definite ed utilizzate almeno le seguenti funzioni:
 * una funzione `ÈTernaPitagoriga(a int, b int, c int) bool` che riceve in input tre valori `int` nei parametri `a`, `b` e `c`, rispettivamente, e restituisce `true` se `c²` è uguale a `a² + b² ` e `false` altrimenti.

##### Esempio d'esecuzione:
 
```text
$ go run terna_pitagorica.go
Inserisci la soglia: 10
Terne pitagoriche:
(3, 4, 5)
(4, 3, 5)

$ go run terna_pitagorica.go
Inserisci la soglia: 20
Terne pitagoriche:
(3, 4, 5)
(4, 3, 5)
(5, 12, 13)
(6, 8, 10)
(8, 6, 10)
(8, 15, 17)
(9, 12, 15)
(12, 5, 13)
(12, 9, 15)
(15, 8, 17)
```
## 13 Parola palindroma

**Definizione**: Una parola è palindroma se può essere letta normalmente, da sinistra verso destra, sia viceversa, cioè da destra verso sinistra.

Scrivere un programma che:
* legga da **standard input** una stringa senza spazi;
* stampi a video il messaggio `Palindroma` nel caso in cui la stringa letta sia palindroma e `Non palindroma` altrimenti.

Oltre alla funzione `main()`, devono essere definite ed utilizzate almeno le seguenti funzioni:
* una funzione `ÈPalindroma(s string) bool` che riceve in input un valore `string` nel parametro `n` e restituisce `true` se `s` è palindroma e `false` altrimenti.

##### Esempio d'esecuzione:

```text
$ go run palindroma.go
anna
Palindroma

$ go run palindroma.go
anni
Non palindroma

$ go run palindroma.go
osso
Palindroma

$ go run palindroma.go
èssè
Palindroma

$ go run palindroma.go
èsse
Non palindroma
```
