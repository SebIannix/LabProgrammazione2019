# Laboratorio 7 - Array e slice II
## 1 Qual è l'outout?

Qual è l'output di questo programma?

```go
package main

import "fmt"

func main() {
	var a = [6]int{1, 2, 3, 4, 5, 6}
	var b []int
	stampaArray(a)
	modificaArray(a)
	stampaArray(a)
	fmt.Println("==============")
	b = a[:]
	stampaSlice(b)
	modificaSlice(a[:])
	stampaSlice(b)
	stampaArray(a)
	fmt.Println("++++++++++++++")
	modificaArrayConPuntatore(&a)
	stampaArray(a)
}
func stampaArray(a [6]int) {
	for _, v := range a {
		fmt.Print(v, " ")
	}
	fmt.Println()
}
func modificaArray(a [6]int) {
	a[0] = 10
}
func stampaSlice(a []int) {
	for _, v := range a {
		fmt.Print(v, " ")
	}
	fmt.Println()
}
func modificaSlice(a []int) {
	a[0] = 10
}
func modificaArrayConPuntatore(a *[6]int) {
	a[0] = 100	//i.e., (*a)[0] = 100
}
```
## 2 Qual è l'outout?

Qual è l'output di questo programma?

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {

	slInt := []int64{2, 4, 8, 16, 32}
	sl1 := converti(slInt)
	sl2 := convertiConAppend(slInt)

	fmt.Println(sl1)
	fmt.Println(sl2)
}

func converti(slIn []int64) []string {

	slOut := make([]string,len(slIn))

	for i:=0; i<len(slIn); i++{
		slOut[i] = strconv.FormatInt(slIn[i], 2)
	}

	return slOut
}

func convertiConAppend(slIn []int64) []string {

	var slOut []string

	for i:=0; i<len(slIn); i++{
		slOut = append(slOut, strconv.FormatInt(slIn[i], 2))
	}

	return slOut
}
```
## 3 Qual è l'output?

Supponendo che l'utente inserisca da **riga di comando** i valori `1 a 5.6 ciao true`, cosa dovrebbe produrre in output il seguente programma?

```go
package main

import (
	"os"
	"fmt"
)

func main() {
	fmt.Printf("tipo os.Args: %T\n", os.Args)
	for i, v := range os.Args {
		fmt.Printf("os.Args[%d]: tipo %T valore %s\n", i, v, v)
	}
}
```

## 4 Filtra e moltiplica

Scrivere un programma che legga da **riga di comando** una sequenza di valori e stampi a video il risultato della moltiplicazione tra i valori che rappresentano numeri interi.

##### Esempio d'esecuzione:

```text
$ go run prodotto.go 4 3           
Il risultato della moltiplicazione tra i numeri interi è 12

$ go run prodotto.go 6  
Il risultato della moltiplicazione tra i numeri interi è 6

$ go run prodotto.go 1 3 a 5 ciao 2
Il risultato della moltiplicazione tra i numeri interi è 30
```

## 5 Minimo, massimo e valor medio

Scrivere un programma che legga da **riga di comando** una sequenza di valori interi e stampi a video il valore minimo, massimo e medio tra i valori letti.

Oltre alla funzione `main()`, devono essere definite ed utilizzate almeno le seguenti funzioni:
* una funzione `Minimo(sl []int) int` che riceve in input un valore `[]int` nel parametro `sl` e restituisce il minimo valore intero presente in `sl`.
* una funzione `Massimo(sl []int) int` che riceve in input un valore `[]int` nel parametro `sl` e restituisce il massimo valore intero presente in `sl`.
* una funzione `Media(sl []int) float64` che riceve in input un valore `[]int` nel parametro `sl` e restituisce un valore reale pari alla media aritmetica dei valori interi presenti in `sl`.

**Suggerimento**: Il numero di valori interi che il programma deve considerare è pari a `len(os.Args)-1`.

*Nota:* Per creare dinamicamente una slice, si utilizzi la funzione `make`.

##### Esempio d'esecuzione:

```text
$ go run min_max_media.go 1 2 3 4
Minimo: 1
Massimo: 4
Valore medio: 2.5

$ go run min_max_media.go -1 10 6 
Minimo: -1
Massimo: 10
Valore medio: 5

$ go run min_max_media.go -1 -2 -3
Minimo: -3
Massimo: -1
Valore medio: -2
```
## 6 Fattoriale

**Definizione**: Si definisce fattoriale di un numero naturale, il prodotto dei numeri interi positivi minori o uguali a tale numero. Il fattoriale di `k` è uguale a `1*2*3*...*(k-3)*(k-2)*(k-1)*k`.

Scrivere un programma che legga da **riga di comando** un numero intero `n` e stampi a video il fattoriale di tutti i numeri compresi tra `1` e `n` (estremi inclusi).

Oltre alla funzione `main()`, devono essere definite ed utilizzate almeno le seguenti funzioni:
* una funzione `Fattoriali(n int) (f []int)` che riceve in input un valore `int` nel parametro `n` e restituisce il valore `f` di tipo `[]int` in cui in `f[0]` è memorizzato il fattoriale di `1`, in `f[1]` è memorizzato il fattoriale di `2`, ..., in `f[n-1]` è memorizzato il fattoriale di `n`.

*Nota:* Per creare dinamicamente una slice, si utilizzi la funzione `make`.

##### Esempio d'esecuzione:

```text
$ go run fattoriale.go 2
Fattoriali: [1 2]

$ go run fattoriale.go 3
Fattoriali: [1 2 6]

$ go run fattoriale.go 10
Fattoriali: [1 2 6 24 120 720 5040 40320 362880 3628800]
```
## 7 Somma di prodotti

Scrivere un programma che legga da **riga di comando** una sequenza di numeri interi di lunghezza pari. Data la sequenza, il programma deve moltiplicare ciascun numero in una posizione di indice pari per il successivo numero in posizione di indice dispari e sommare i prodotti ottenuti. 

*Esempio:* Se `10 2 3 4 5 6` è la sequenza letta, allora la somma calcolata deve essere `10*2 + 3*4 + 5*6 = 62`.

Il programma deve infine stampare a video il valore della somma calcolata.
 
Oltre alla funzione `main()`, devono essere definite ed utilizzate almeno le seguenti funzioni:
* una funzione `Calcola(sl []int) int` che riceve in input un valore `[]int` nel parametro `sl` e restituisce un valore di tipo `int` pari alla somma dei prodotti ottenuti moltiplicando ciascun numero in una posizione di indice pari di `sl` per il successivo numero in posizione di indice dispari. 

*Nota:* Per creare dinamicamente una slice, si utilizzi la funzione `make`.

##### Esempio d'esecuzione:

```text
$ go run somma_prodotto.go 1 2 3 4 5 6
La somma è 44

$ go run somma_prodotto.go 7 3 1 8
La somma è 29
```
## 8 Somma di prodotti pari

Scrivere un programma che:
 
 * legga da **riga di comando** una sequenza di numeri interi;
 * stampi a video il risultato della somma dei prodotti pari associati alle coppie non ordinate di numeri che si possono definire a partire dai numeri letti (data la coppia non ordinata di numeri `(numero_1, numero_2)`, il valore del prodotto associato è `numero_1 * numero_2`).

*Esempio:* Se `10 1 31 4` è la sequenza letta, le coppie non ordinate di numeri che si possono definire a partire dai numeri letti sono: `(10, 1); (10, 31); (10, 4); (1, 31); (1, 4); (31, 4)`. Di queste, quelle il cui prodotto è pari sono: `(10, 1); (10, 31); (10, 4); (1, 4); (31, 4)`. La somma dei prodotti pari è `488`.

Oltre alla funzione `main()`, devono essere definite ed utilizzate almeno le seguenti funzioni:
* una funzione `Calcola(sl []int) int` che riceve in input un valore `[]int` nel parametro `sl` e restituisce un valore di tipo `int` pari alla somma dei prodotti pari associati alle coppie non ordinate di numeri che si possono definire a partire dai numeri presenti in `sl`. 

*Nota:* Per creare dinamicamente una slice, si utilizzi la funzione `make`.

##### Esempio d'esecuzione:

```text
$ go run prodotti_pari.go 1 2 3 4 5 6
La somma è 152

$ go run prodotti_pari.go 1 2 3 4 5  
La somma è 62
```
## 9 Filtra voti

Scrivere un programma che:
 * legga da **riga di comando** una sequenza di valori (i valori numerici interi che compaiono all'interno della sequenza rappresentano voti in una scala di valutazione tra 0 e 100; i valori numerici interi superiori a 60 corrispondono a voti sufficienti); 
 * stampi a video le due sottosequenze di valori numerici interi che corrisponodo rispettivamente a voti insufficienti e sufficienti.
 
Oltre alla funzione `main()`, devono essere definite ed utilizzate almeno le seguenti funzioni:
* una funzione `LeggiNumeri() (numeri []int)` che restituisce un valore di tipo `[]int` in cui sono memorizzati i valori numerici interi specificati a **riga di comando**;
* una funzione `FiltraVoti(voti []int) (sufficienti, insufficienti []int)` che riceve in input un valore `[]int` nel parametro `voti` e restituisce due valori di tipo `[]int`, `sufficienti` e `insufficienti`, in cui sono memorizzati rispettivamente i voti sufficienti e insufficienti presenti in `voti`. 
 
##### Esempio d'esecuzione:

```text
$ go run filtro.go 80 75 60 55 
Voti sufficienti: [80 75 60]
Voti insufficienti: [55]

$ go run filtro.go 100 98 59 40
Voti sufficienti: [100 98]
Voti insufficienti: [59 40]
```
## 10 Numeri casuali

Scrivere un programma che:
1) Legga da **riga di comando** un numero intero `soglia`;
2) Generi in modo casuale una sequenza di lunghezza arbitraria di numeri interi compresi nell'intervallo che va da 0 a 100, estremi inclusi. Il processo di generazione si interrompe quando viene generato un numero inferiore a `soglia`.
3) Stampi a video tutti i numeri generati.
4) Stampi a video tutti i numeri generati superiori a `soglia`.

Oltre alla funzione `main()`, devono essere definite ed utilizzate almeno le seguenti funzioni:
* una funzione `Genera(soglia int) []int` che riceve in input un valore `int` nel parametro `soglia` e restituisce un valore di tipo `[]int` in cui è memorizzata una sequenza di lunghezza arbitraria di numeri interi, generata in base alle specifiche di cui al punto 2.

*Suggerimento:* per generare in modo casuale un numero intero, potete utilizzare le funzioni dei package `math/rand` e `time` come mostrato nel seguente frammento di codice:
```go
/* inizializzazione del generatore di numeri casuali */
rand.Seed(int64(time.Now().Nanosecond())) 
/* generazione di un numero casuale compreso nell'intervallo 
   che va da 0 a 99 (estremi inclusi) */
numeroGenerato := rand.Intn(100)
```

##### Esempio d'esecuzione:

```text
$ go run numeri_random.go 20
Valori generati [21 72 44 64 30 13]
Valori sopra soglia: [21 72 44 64 30]
```
## 11 Filtra testo

Scrivere un programma che: 
* legga da **standard input** un testo su più righe (alcune delle quali possono essere delle righe vuote (`""`));
* termini la lettura quando, premendo la combinazione di tasti `Ctrl+D`, viene inserito da **standard input** l'indicatore End-Of-File (EOF);
* ristampi le righe del testo letto in cui compaiono cifre.

Il programma deve essere organizzato nelle seguenti funzioni:
* una funzione `LeggiTesto() []string` che legge da **standard input** un testo su più righe (alcune delle quali possono essere delle righe vuote (`""`)) e terminato dall'indicatore EOF, restituendo un valore `[]string` in cui sono memorizzate le stringhe corrispondenti alle righe del testo letto;
* una funzione `ContieneCifre(s string) bool` che riceve in input un valore `string` nel parametro `s` e restituisce `true` se almeno un carattere in `s` è una cifra, `false` altrimenti;
* una funzione `FiltraTesto(sl []string) []string` che riceve in input un valore `[]string` nel parametro `sl` e restituisce un valore `[]string` in cui sono  memorizzate le stringhe di `sl` in cui compaiono cifre; la funzione deve utilizzare la funzione `ContieneCifre()`;
* una funzione `main()` che richiama la funzione `LeggiTesto()` per leggere il testo da **standard input**, filtra il testo letto (richiamando opportunamente la funzione `FiltraTesto()`) e () stampa il testo filtrato.

*Suggerimento:* per sapere se un carattere è una cifra si può far riferimento alla documentazione della funzione `unicode.IsDigit` del package `unicode`.

```text
$ go run filtra_testo.go
riga senza cifre
riga con 1 cifra         
2a riga con 2 cifre
nessuna cifra 
Testo filtrato:
riga con 1 cifra
2a riga con 2 cifre
``` 
