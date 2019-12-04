# Laboratorio 8 - Array e slice III
## 1 Qual è l'output?

Qual è l'output del seguente programma?

```go
package main

import (
	"fmt"
	"strings"
)

const Dimensione = 10

func main() {

	var a []int

	for i := 0; i < Dimensione; i++ {
		a = append(a, i+1)
		fmt.Println("Iterazione", i, ":", a, len(a), cap(a))
	}

	fmt.Println(strings.Repeat("=", 10))
	fmt.Println("i)", a, len(a), cap(a))
	a = a[:cap(a)]
	fmt.Println("ii)", a, len(a), cap(a))
	a = a[:Dimensione]

	a = append(a[Dimensione/2:], a[:Dimensione/2]...)

	fmt.Println("iii)", a, len(a), cap(a))
	a = a[:cap(a)]
	fmt.Println("iv)", a, len(a), cap(a))
	a = a[:Dimensione]

	for i := 0; i < Dimensione; i++ {
		a[i] = i + 1
	}

	a = append(a[:Dimensione/5], a[(4*Dimensione)/5:]...)

	fmt.Println(strings.Repeat("=", 10))
	fmt.Println("v)", a, len(a), cap(a))
	a = a[:cap(a)]
	fmt.Println("vi)", a, len(a), cap(a))

}
```
## 2 Qual è l'output?

Qual è l'output del seguente programma?

```go
package main

import (
	"fmt"
	"strings"
)

const Dimensione = 10

func main() {

	var a []int

	for i := 0; i < Dimensione; i++ {
		a = append(a, i+1)
		fmt.Println("Iterazione", i, ":", a, len(a), cap(a))
	}

	fmt.Println(strings.Repeat("=", 10))
	fmt.Println("i)", a, len(a), cap(a))
	a = a[:cap(a)]
	fmt.Println("ii)", a, len(a), cap(a))
	a = a[:Dimensione]

	b := append(a[Dimensione/2:], a[:Dimensione/2]...)

	fmt.Println(strings.Repeat("=", 10))
	fmt.Println("iii)\n a: ", a, len(a), cap(a), "\n b:", b, len(b), cap(b))
	a = a[:cap(a)]
	b = b[:cap(b)]
	fmt.Println("iv)\n a: ", a, len(a), cap(a), "\n b:", b, len(b), cap(b))
	a = a[:Dimensione]
	b = b[:Dimensione]

	c := append(a[:Dimensione/5], a[(4*Dimensione)/5:]...)

	fmt.Println(strings.Repeat("=", 10))
	fmt.Println("v)\n a:", a, len(a), cap(a), "\n b:", b, len(b), cap(b), 
"\n c:", c, len(c), cap(c))
	a = a[:cap(a)]
	b = b[:cap(b)]
	c = c[:cap(c)]
	fmt.Println("vi)\n a:", a, len(a), cap(a), "\n b:", b, len(b), cap(b), 
"\n c:", c, len(c), cap(c))
}
```
## 3 Qual è l'output?

Qual è l'output del seguente programma?

```go
package main

import (
	"fmt"
)

const Dimensione = 10

func main() {

	var a []int

	for i := 0; i < Dimensione; i++ {
		a = append([]int{i+1}, a...)
	}

	fmt.Println(a)

	b := make([]int, Dimensione)

	copy(b[:], a[Dimensione/2:])

	fmt.Println(b)

	c := make([]int, Dimensione)

	for i := 0; i < Dimensione; i += 1 {
		copy(c[Dimensione-1-i:], a[i:i+1])
	}
	fmt.Println(c)
}
```
## 4 Qual è l'output?

Qual è l'output del seguente programma?

```go
package main

import (
	"fmt"
)

func main() {

	var a [3][2]int = [3][2]int{{1, 2}, {10, 20}, {100, 200}}
	//a := [3][2]int{{1, 2}, {10, 20}, {100, 200}}
	//a := [...][2]int{{1, 2}, {10, 20}, {100, 200}}

	fmt.Println("a:", a)
	/* a è un array bi-dimensionale di tipo [3][2]int 
        con lunghezza pari a 3 */
	fmt.Printf("Tipo di a: %T\n", a)

	fmt.Println()

	for i, r := range a {
		fmt.Printf("a[%d]: %v\n", i, r)
	}
	for i := 0; i < len(a); i++ {
		/* a[i], con 0 <= i < len(a), è di tipo [2]int */
		fmt.Printf("Tipo di a[%d]: %T\n", i, a[i])
	}

	fmt.Println()

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			fmt.Printf("%d ", a[i][j])
		}
		fmt.Println()
	}
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			/* a[i][j], con 0 <= i < len(a) e con 0 <= j < len(a[i]), 
                è di tipo int */
			fmt.Printf("a[%d][%d] è il valore alla riga %d e " +
				"colonna %d dell'array bi-dimensionale a: %d\n", i, j, i, j, a[i][j])
		}
	}

}
```
## 5 Qual è l'output?

Qual è l'output del seguente programma?

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	var slc [][]int
	fmt.Printf("Tipo di slc: %T\n", slc)

	if slc == nil {
		fmt.Println("'nil' è lo zero-value di una variabile di tipo " +
			"'reference type' non inizializzata.")
	}
	fmt.Printf("slc: %v %d %d\n", slc, len(slc), cap(slc))
	fmt.Println()

	n := 4
	slc = CreaSliceBidimensionale(n)

	InizializzaSliceBidimensionale(slc)

	fmt.Println("slc:")
	for _, r := range slc {
		fmt.Printf("%v\n", r)
	}

	slc_soprasoglia := FiltraSliceBidimensionale(slc, 2)

	fmt.Println("slc_soprasoglia:")
	for _, r := range slc_soprasoglia {
		fmt.Printf("%v\n", r)
	}

}

/* 'CreaSliceBidimensionale' riceve in input un valore int nel parametro l 
e restituisce un valore s di tipo [][]int (una slice bi-dimensionale
di interi) con lunghezza/capacità pari a l in cui s[i], con 0 <= i < l, 
è un valore di tipo []int (una slice di interi) con lunghezza/capacità pari a l
*/
func CreaSliceBidimensionale(l int) [][]int {
	var s [][]int
	/* s è una slice bi-dimensionale di tipo [][]int*/
	fmt.Printf("Tipo di s: %T\n", s)
	if s == nil {
		// 'nil' è lo zero-value di una variabile di tipo 
		// 'reference type' non inizializzata
		fmt.Printf("s == nil \t=>\t")
		fmt.Printf("s: %v %d %d\n", s, len(s), cap(s))
	}
	s = make([][]int, l)
	/* s è una slice bi-dimensionale di tipo [][]int con lunghezza/capacità 
        pari a l */
	fmt.Printf("s: %v %d %d\n\n", s, len(s), cap(s))

	for i := 0; i < l; i++ {
		/* s[i], con 0 <= i < dim, è di tipo []int */
		fmt.Printf("Tipo di s[%d]: %T\n", i, s[i])
		if s[i] == nil {
			// 'nil' è lo zero-value di una variabile di tipo 
			// 'reference type' non inizializzata
			fmt.Printf("s[%d] == nil \t=>\t", i)
			fmt.Printf("s[%d]: %v %d %d\n", i, s[i], len(s[i]), cap(s[i]))
		}
		s[i] = make([]int, l)
		/* s[i] è una slice con lunghezza/capacità pari a l */
		fmt.Printf("s[%d]: %v %d %d\n\n", i, s[i], len(s[i]), cap(s[i]))
	}
	return s
}

/* 'InizializzaSliceBidimensionale' riceve in input un valore [][]int nel 
parametro s ed inizializza s[i][j] (con 0 <= i < len(s) e 0 <= j <len(s[i])) 
con un valore intero estratto in maniera casuale nell'insieme {0,1}
ossia
'InizializzaSliceBidimensionale' inizializza una slice bi-dimensionale di 
interi con valori interi estratti in maniera casuale nell'insieme {0,1} */
func InizializzaSliceBidimensionale(s [][]int) {

	rand.Seed(int64(time.Now().Nanosecond()))

	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s[i]); j++ {
			s[i][j] = rand.Intn(2)
		}
	}

}

/* 'FiltraSliceBidimensionale' riceve in input un valore [][]int ed un valore 
int nei parametri s e soglia, rispettivamente, e restituisce un 
valore [][]int definito dai valori s[i] di s (con 0 <= i < len(s)) che 
includono valori interi la cui somma è maggiore o uguale a soglia */
func FiltraSliceBidimensionale(s [][]int, soglia int) [][]int {

	var s_soprasoglia [][]int

	for i := 0; i < len(s); i++ {
		somma := 0
		for j := 0; j < len(s[i]); j++ {
			somma += s[i][j]
		}
		if somma >= soglia {
			s_soprasoglia = append(s_soprasoglia, s[i])
		}
	}

	return s_soprasoglia
}
```
## 6 Qual è l'output?

Qual è l'output del seguente programma?

```go
package main

import "fmt"

func main() {

	//strFrom := "ABDEF"
	//strFrom[2] = 'C' /* error: "cannot assign to str[2]" */

	/* (1) */
	strFrom1 := "ABDEF"
	// len(sliceDiRune) è pari al numero di caratteri Unicode presenti in strFrom1
	sliceDiRune := []rune(strFrom1)
	sliceDiRune[2] = 'C'
	strTo1 := string(sliceDiRune)
	fmt.Println(strTo1)
	/* (1) - end */

	/* (2) */
	strFrom2 := "ABDEF"
	//strTo2 := strFrom2[:2]+string('C')+strFrom2[3:]
	strTo2 := strFrom2[:2]+"C"+strFrom2[3:]
	fmt.Println(strTo2)
	/* (2) - end */

}
```
## 7 Tavola pitagorica

Scrivere un programma che legga da **riga di comando** un numero intero `n` e, come mostrato nell'**Esempio di esecuzione**, stampi a video la corrispondente tavola pitagorica `n x n`.

Oltre alla funzione `main()`, devono essere definite ed utilizzate almeno le seguenti funzioni:

* una funzione `CreaTavolaPitagorica(n int) [][]int` che riceve in input un valore `int` nel parametro `n` e restituisce un valore di tipo `[][]int` in cui sono memorizzati i valori di una tavola pitagorica `n x n`;
* una funzione `StampaTavolaPitagorica(s [][]int)` che riceve in input un valore di tipo `[][]int` nel parametro `s` e, come mostrato nell'**Esempio di esecuzione**, stampa la tavola pitagorica corrispondete ai valori memorizzati `s`.

##### Esempio d'esecuzione:

```text
$ go run tavola_pitagorica.go 5
   1    2    3    4    5 
   2    4    6    8   10 
   3    6    9   12   15 
   4    8   12   16   20 
   5   10   15   20   25 

$ go run tavola_pitagorica.go 10
   1    2    3    4    5    6    7    8    9   10 
   2    4    6    8   10   12   14   16   18   20 
   3    6    9   12   15   18   21   24   27   30 
   4    8   12   16   20   24   28   32   36   40 
   5   10   15   20   25   30   35   40   45   50 
   6   12   18   24   30   36   42   48   54   60 
   7   14   21   28   35   42   49   56   63   70 
   8   16   24   32   40   48   56   64   72   80 
   9   18   27   36   45   54   63   72   81   90 
  10   20   30   40   50   60   70   80   90  100
```
## 8 Date (1)

Scrivere un programma che:
* legga da **standard input** un testo su più righe;
* termini la lettura quando viene inserita da standard input una riga vuota (`""`).

Ogni riga di testo è una una stringa senza spazi che codifica una data in uno dei seguenti possibili formati:

1. aaaa/m/g
2. aaaa/m/gg
3. aaaa/mm/g
4. aaaa/mm/gg

Una volta terminata la fase di lettura, il programma deve stampare la codifica nel formato aaaa-mm-gg delle date lette.

Oltre alla funzione `main()` devono essere definite ed utilizzate almeno le seguenti funzioni:
* una funzione `LeggiTesto() []string` che legge da **standard input** un testo su più righe e terminato da una riga vuota (`""`), restituendo un valore `[]string` in cui sono memorizzate le righe del testo letto;
* una funzione `Formatta(s string) string` che riceve in input un valore `string` nel parametro `s` in cui è codificata una data nel formato **aaaa/m/g**, **aaaa/m/gg**, **aaaa/mm/g** o **aaaa/mm/gg**, e restituisce un valore `string` in cui la data in input è codificata nel formato **aaaa-mm-gg**.

##### Esempio d'esecuzione:

```text
$ go run date.go
$ go run date.go
2019/04/03
2019/6/4
2019/07/5
2019/4/05

2019-04-03
2019-06-04
2019-07-05
2019-04-05
```
## 9 Coppie

Scrivere un programma che:
* legga da **riga di comando** un numero intero `n`;
* utilizzi le funzioni `CreaSliceBidimensionale(l int) [][]int` e `InizializzaSliceBidimensionale([][]int) ` descritte nell'Esercizio 5 (Laboratorio 8 - Array e slice III) per inizializzare una variabile `s` di tipo `[][]int` con lunghezza/capacità pari a `n` in cui ogni elemento `s[i]`, con `0 <= i < l`, è un valore di tipo `[]int` con lunghezza/capacità pari a `n`;
* stampi a video tutte le coppie di indici `(i, j)`, con `0 <= i < l` e `0 <= j < l`, tali che `s[i][j]` è uguale a `1`.

Oltre alle funzioni `main()`, `CreaSliceBidimensionale()`, e `InizializzaSliceBidimensionale()` devono essere definite ed utilizzate almeno le seguenti funzioni:
* una funzione `Coppie(s [][]int) (coppie [][]int)` che riceve in input un valore `[][]int` nel parametro `s` e restituisce il valore di tipo `[][]int` nella variabile `coppie` in cui sono memorizzate tutte le coppie di indici `(i, j)`, con `0 <= i < l` e `0 <= j < l`, tali che `s[i][j]` è uguale a `1` (`coppie[k]`, con `0 <= k < len(coppie)`, è un valore di tipo` []int` di lunghezza `2`).

##### Esempio d'esecuzione:

```text
$ go run coppie.go 4
[1 1]
[1 3]
[2 1]
[2 2]
[3 2]
[3 3]
```
## 10 Mazzo di carte

Scrivere un programma che:
1. Legga da **riga di comando** un numero intero `n` tale che `0 < n < 10`.
2. Inizializzi una stringa che rappresenti un mazzo di carte formato dalle sole carte di cuori.
    1. Le carte di cuori corrispondono ai caratteri con codice Unicode compreso nell'intervallo che va da `'\U0001F0B1'` a `'\U0001F0BA'`, estremi inclusi. 
    2. Le carte del mazzo non sono mischiate: la prima carta del mazzo è l'asse di cuori; la seconda carta del mazzo è il due di cuori;... l'ultima carta del mazzo è il dieci di cuori.
3. Simuli l'estrazione casuale (ed in sequenza) di `n` carte dal mazzo, stampando a video le carte estratte e quelle rimaste nel mazzo dopo ogni estrazione. 

Oltre alla funzione `main()`, devono essere definite ed utilizzate almeno le seguenti funzioni:

* una funzione `LeggiNumero() int` che legge da **riga di comando** un numero intero e ne restituisce il valore;
* una funzione `GeneraMazzo() string` che restituisce un valore `string` che rappresenta un mazzo di carte formato dalle sole carte di cuori;
* una funzione `EstraiCarta(mazzo string) (cartaEstratta rune, mazzoResiduo string)` che riceve in input un valore `string` nel parametro `mazzo` e restituisce un valore di tipo `rune` nella variabile `cartaEstratta` ed un valore di tipo `string` nella variabile `mazzoResiduo`, che rappresentano rispettivamente la carta estratta in modo casuale dal mazzo di carte rappresentato da `mazzo` ed il mazzo residuo di carte dopo l'estrazione;
* una funzione `EstraiCarte(mazzo string, n int)` che riceve in input un valore `string` nel parametro `mazzo` ed un valore `int` nel parametro `n` e simula l'estrazione casuale (ed in sequenza) di `n` carte dal mazzo di carte rappresentato da `mazzo`, stampando a video le carte estratte e quelle rimaste nel mazzo dopo ogni estrazione; la funzione deve utilizzare la funzione `EstraiCarta()`.

##### Esempio d'esecuzione:

```text
$ go run carte.go 6
Estratta la carta 🂱 - Carte rimaste nel mazzo: 🂲🂳🂴🂵🂶🂷🂸🂹🂺
Estratta la carta 🂸 - Carte rimaste nel mazzo: 🂲🂳🂴🂵🂶🂷🂹🂺
Estratta la carta 🂹 - Carte rimaste nel mazzo: 🂲🂳🂴🂵🂶🂷🂺
Estratta la carta 🂷 - Carte rimaste nel mazzo: 🂲🂳🂴🂵🂶🂺
Estratta la carta 🂵 - Carte rimaste nel mazzo: 🂲🂳🂴🂶🂺
Estratta la carta 🂴 - Carte rimaste nel mazzo: 🂲🂳🂶🂺
```


## 11 Date (2)

Scrivere un programma che:
* legga da **standard input** un testo su più righe;
* termini la lettura quando viene inserita da standard input una riga vuota ( "" ).

Ogni riga di testo è una una stringa senza spazi che codifica una data in uno dei seguenti possibili formati:

1. aaaa/m/g
2. aaaa/m/gg
3. aaaa/mm/g
4. aaaa/mm/gg
5. g/m/aaaa
6. gg/m/aaaa
7. g/mm/aaaa
8. gg/mm/aaaa

Una volta terminata la fase di lettura, il programma deve stampare la codifica nel formato aaaa-mm-gg delle date lette.

Oltre alla funzione `main()` devono essere definite ed utilizzate almeno le seguenti funzioni:
* una funzione `daInvertire(s string) bool` che riceve in input un valore `string` nel parametro `s` in cui è codificata una data nel formato **aaaa/m/g**, **aaaa/m/gg**, **aaaa/mm/g**, **aaaa/mm/gg**, **g/m/aaaa**, **gg/m/aaaa**, **g/mm/aaaa**, o **gg/mm/aaaa** e restituisce `true` se in `s` è codificata una data nel formato **g/m/aaaa**, **gg/m/aaaa**, **g/mm/aaaa**, o **gg/mm/aaaa**, `false` altrimenti.
* una funzione `Inverti(s string) string` che riceve in input un valore  `string` nel parametro `s` e restituisce un valore `string` in cui il primo carattere è l'ultimo che definisce `s`, il secondo carattere è il penultimo che definisce `s`, ...  e l'ultimo carattere è il primo che definisce `s`.


##### Esempio d'esecuzione:

```text
$ go run date.go       
2019/04/03
2019/6/4
2019/07/5
2019/4/05
5/5/2019
05/4/2019
7/05/2019
07/09/2019

2019-04-03
2019-06-04
2019-07-05
2019-04-05
2019-05-05
2019-04-05
2019-05-07
2019-09-07
```
