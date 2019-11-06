# Laboratorio 4 - Tipi di dato
## 1 Qual è l'output?

Qual è l'output del seguente programma?

```go
package main

import "fmt"

func main() {
	var x int
	var y int = 'a'
	x = 'A'
	fmt.Print(x, " ")
	for x := 1; x < 10; x++ {
		fmt.Print(x+y, " ")
	}
	fmt.Println()
}
```
## 2 Qual è l'output?

Qual è l'output del seguente programma?

```go
package main

import "fmt"

func main() {
 var x int = 10
 var y float64 = 2.5
 t := 100
 t, z, k := x/int(y), float64(x)/y, t*2
 fmt.Println(t, z, k)
}
```
## 3 Trova l'errore

Questo programma dovrebbe stampare il valore delle variabili `a`. `b`. `c`, `d` ed `e` ma contiene degli errori. Corregere gli errori e verificare che l'esecuzione produca l'output desiderato.


```go
package main

import "fmt"

func main() {
	var a int = 4
	var b float64 = 12.5
	var e
	var c int
	c := a + b
	d = a/b
	fmt.Println(a, b, c, d, f)
}
```
## 4 Overflow

Qual è l'output del seguente programma?

```go
import "fmt"

func main() {
	var (
		a, b int8 = 30, 100
	)
	somma := a + b
	fmt.Println("La somma di", a, "e", b, "è", somma)
}
```

## 5 Valori rappresentabili

Il seguente programma stampa a video i limiti dei valori rappresentabili con i diversi tipi di dato numerici.
Per ogni tipo di dato, i corrispondenti limiti sono definiti come costanti nel package `math`.
Si stampi a video il valore di tali limiti eseguendo il programma.

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("uint8:", 0, math.MaxUint8)
	fmt.Println("uint16:", 0, math.MaxUint16)
	fmt.Println("uint32:", 0, math.MaxUint32)
	fmt.Println("uint64:", 0, uint64(math.MaxUint64))

	fmt.Println("int8:", math.MinInt8, math.MaxInt8)
	fmt.Println("int16:", math.MinInt16, math.MaxInt16)
	fmt.Println("int32:", math.MinInt32, math.MaxInt32)
	fmt.Println("int64:", math.MinInt64, math.MaxInt64)

    /*  Floating-point limit values: 
        - Max is the largest finite value representable by the type.
        - SmallestNonzero is the smallest positive, non-zero value
        representable by the type. */

	fmt.Println("float32 - SmallestNonzero:", math.SmallestNonzeroFloat32)
	fmt.Println("float32:", -math.MaxFloat32, math.MaxFloat32)
	fmt.Println("float64 - SmallestNonzero:", math.SmallestNonzeroFloat64)
	fmt.Println("float64:", -math.MaxFloat64, math.MaxFloat64)

	fmt.Println("complex64:", complex(-math.MaxFloat32, -math.MaxFloat32), 
        complex(math.MaxFloat32, math.MaxFloat32))
	fmt.Println("complex128:", complex(-math.MaxFloat64, -math.MaxFloat64), 
        complex(math.MaxFloat64, math.MaxFloat64))

}
```
## 6 Uguaglianza tra valori

Scrivere un programma che:  
1) Legga da **standard input** un numero reale.  
2) Calcoli la radice quadrata del numero letto (sia `x` la variabile di tipo `float64` in cui è memorizzato il valore; la radice quadrata di del valore reale può essere calcolata utilizzando la funzione `math.Sqrt` del package `math` nel seguente modo `radiceQuadrata := math.Sqrt(x)`).  
3) Stampi a video `x, "uguale a", radiceQuadrata, "*", radiceQuadrata, "\n"` nel caso in cui `radiceQuadrata*radiceQuadrata` sia uguale a `x` e `x, "diverso da", radiceQuadrata, "*", radiceQuadrata, "\n"` altrimenti.  
 
Perché inserendo 10 da **standard input** il programma stampa a video `10 diverso da 3.1622776601683795 * 3.1622776601683795`?

 
##### Esempio d'esecuzione:

```markdown
$ go run uguaglianza.go
Inserisci un numero: 4
4 uguale a 2 * 2

$ go run uguaglianza.go 
Inserisci un numero: 10
10 diverso da 3.1622776601683795 * 3.1622776601683795
``` 
## 7 Valore di default delle variabili

Che cosa succede se si prova a utilizzare una variabile senza averla inizializzata? Si esegua il seguente programma per stampare a video i valori di default (valori iniziali/zero-values) per i tipi `int`, `float64`, `string` e `bool`.

```go
package main

import "fmt"

func main() {

	var numeroIntero int
	var numeroReale float64
	var stringa string
	var valoreLogico bool

	fmt.Print("Valore di default per il tipo int: _", numeroIntero, "_\n")
	fmt.Print("Valore di default per il tipo float64: _", numeroReale, "_\n")
	fmt.Print("Valore di default per il tipo string: _", stringa, "_\n")
	fmt.Print("Valore di default per il tipo bool: _", valoreLogico, "_\n")

}
```
## 8 Qual è l'errore?

Si consideri il seguente frammento di codice. 

```go
package main

import "fmt"

func main() {

	var a, b int
	fmt.Scan(&a, &b)

	var c ???

	c = a == b
	
	if c {
		fmt.Println("uguali")
	} else {
		fmt.Println("diversi")
	}
}
```

Di che tipo deve essere la variabile `c` affinché la compilazione del codice non generi errori?
## 9 Operatori e precedenza

Si consideri il seguente frammento di codice.

```go
var a, b, c int = 10, 5, 4
fmt.Println( a / b * c )
```

Quale valore viene stampato a video?
## 10 Divisione per zero

Si consideri il seguente frammento di codice.

```go
var numero float64 = 0
fmt.Println(numero)
numero = 1 / numero
fmt.Println(numero)
numero = 1 / numero
fmt.Println(numero)
numero = 1 / numero
numero = numero - numero
fmt.Println(numero)
```
Quali valori vengono stampati a video?
Quali valori verrebbero stampati a video se la variabile `numero` fosse di tipo `int`?


## 11 Troncamento

Scrivere un programma che legga da **standard input** un numero reale e ne stampi il valore **troncato** alla seconda cifra decimale.

*Suggerimento:* Il valore troncato alla seconda cifra decimale di un numero reale può essere ottenuto effettuando le seguenti operazioni:  
1. Moltiplicare il numero reale per 100   
2. Convertire il valore ottenuto al passo 1) in un valore di tipo `int`   
3. Riconvertire il valore ottenuto al passo 2) in un valore di tipo `float64`   
4. Dividere il valore ottenuto al passo 3) per 100  

oppure:

1. Moltiplicare il numero reale per 100  
2. Utilizzare la funzione `math.Trunc` del package `math` per troncare all'intero valore ottenuto al passo 1) (si utilizzi il comando `go doc math.Trunc` per capire come utilizzare la funzione)  
3. Dividere il valore ottenuto al passo 2) per 100 

##### Esempio d'esecuzione:

```text
$ go run troncamento.go 
Inserisci il valore da troncare: 10.34762
Valore troncato = 10.34

$ go run troncamento.go        
Inserisci il valore da troncare: 8.34267
Valore troncato = 8.34
``` 


## 12 Arrotondamento

Scrivere un programma che legga da **standard input** un numero reale e ne stampi il valore **arrotondato** alla seconda cifra decimale.

*Suggerimento:* Il valore arrotondato alla seconda cifra decimale di un numero reale può essere ottenuto effettuando le seguenti operazioni:

1. Moltiplicare il numero reale per 100
2. Sommare 0.5 al valore ottenuto al passo 1) 
3. Convertire il valore ottenuto al passo 2) in un valore di tipo `int` 
4. Riconvertire il valore ottenuto al passo 3) in un valore di tipo `float64` 
5. Dividere il valore ottenuto al passo 4) per 100

oppure:

1. Moltiplicare il numero reale per 100
2. Utilizzare la funzione `math.Round` del package `math` per arrontondare all'intero valore ottenuto al passo 1) (si utilizzi il comando `go doc math.Round` per capire come utilizzare la funzione)
3. Dividere il valore ottenuto al passo 2) per 100

##### Esempio d'esecuzione:

```text
$ go run arrotondamento.go 
Inserisci il valore da arrotondare: 10.34762
Valore arrotondato = 10.35

$ go run arrotondamento.go
Inserisci il valore da arrotondare: 8.32467
Valore arrotondato = 8.32
``` 


## 13 Troncamento/Arrotondamento generalizzati

Scrivere una versione generalizzata dei programmi **Troncamento** e **Arrotondamento** che legga da **standard input** un intero `n` oltre al numero reale. L'intero `n` specifica che il troncamento e l'arrotondamento devono avvenire  alla `n`-esima cifra decimale.

##### Esempio d'esecuzione:

```text
$ go run generalizzazione.go
Inserisci il valore: 10.34762
Inserisci il numero di cifre dopo la virgola: 4
Valore troncato = 10.3476
Valore arrotondato = 10.3476

$ go run generalizzazione.go
Inserisci il valore:  10.34762
Inserisci il numero di cifre dopo la virgola: 3
Valore troncato = 10.347
Valore arrotondato = 10.348
``` 


## 14 Soluzione di equazioni di primo grado

Scrivere un programma che legga da **standard input** due numeri reali, `a` e `b`, che rappresentano i coefficienti di un'equazione
di primo grado (con incognita `x`) espressa nella forma `ax + b = 0`.
Il programma deve calcolare e stampare a video il valore della radice dell'equazione (il valore di `x` per cui l’uguaglianza risulta verificata). 

*Suggerimento:* Il valore della radice è pari a `-b/a`.

##### Esempio d'esecuzione:

```text
$ go run equazione.go
3 6
La radice è -2

$ go run equazione.go
4 2
La radice è -0.5

$ go run equazione.go
10 -4
La radice è 0.4

$ go run equazione.go
10 -10
La radice è 1
``` 

Cosa succede se `a=0` e `b≠0`?

Cosa succede se `a=0` e `b=0`?
## 15 Soluzione di equazioni di secondo grado con radici reali

Scrivere un programma che legga da **standard input** tre numeri reali, `a`, `b` e `c`, che rappresentano i coefficienti di un'equazione di secondo grado (con incognita `x`) espressa nella forma `ax² + bx + c = 0`.
Il programma deve calcolare e stampare a video il valore delle radici dell'equazione (si assuma che l'equazione ammetta sempre soluzioni reali).

##### Esempio d'esecuzione:

```text
$ go run equazioni.go
1 5 4
Radici dell'equazione: -1 -4

$ go run equazioni.go
1 0 -4
Radici dell'equazione: 2 -2
``` 

Utilizzando il programma appena scritto, provate a risolvere l'equazione `x² + 1 = 0`. Che cosa succede?
## 16 Soluzione di equazioni di secondo grado

Scrivere un programma che legga da **standard input** tre numeri reali, `a`, `b` e `c`, che rappresentano i coefficienti di un'equazione di secondo grado (con incognita `x`) espressa nella forma `ax² + bx + c = 0`.
Il programma deve calcolare e stampare a video il valore delle radici dell'equazione.

*Suggerimento:* Nel caso in cui l'equazione non ammetta soluzioni reali, potete memorizzare i valori delle radici in variabili di tipo `complex`. 

##### Esempio d'esecuzione:

```text
$ go run equazioni.go 
1 5 4
Due radici reali -1 -4

$ go run equazioni.go
1 0 4
Due radici complesse (0+2i) (-0-2i)

$ go run equazioni.go
1 4 4
Due radici reali coincidenti -2 -2
``` 
