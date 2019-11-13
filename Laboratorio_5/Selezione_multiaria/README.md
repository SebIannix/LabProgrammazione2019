# Laboratorio 5 - Selezione multiaria
## 1 Qual è l'output?

Supponendo che l'utente inserisca da **standard input** il valore `10`, cosa dovrebbe produrre in output il seguente programma? E se invece inserisse `4`?

```go
package main

import "fmt"

func main() {

	var voto int

	fmt.Scan(&voto)

	switch voto {
	default:
		fmt.Println("Insufficiente!")
	case 10:
		fallthrough
	case 9:
		fmt.Println("Ottimo!")
	case 8:
		fmt.Println("Distinto!")
	case 7:
		fmt.Println("Buono!")
	case 6:
		fmt.Println("Sufficiente!")
	}
}
```
 
## 2 Qual è l'output?

Supponendo che l'utente inserisca da **standard input** il valore `9`, cosa dovrebbe produrre in output il seguente programma? E se invece inserisse `13`?

```go
package main

import "fmt"

func main() {

	var n int

	fmt.Scan(&n)

	var somma int

	for i:=0; i<n; i++ {

		switch i%2 {
		case 0:
			fmt.Println(i, "pari!")
		case 1:
			continue
		}

		somma += i
	}

	fmt.Println("Somma =", somma)
}
```
 
## 3 Qual è l'output?

Supponendo che l'utente inserisca da **standard input** il valore `3`, cosa dovrebbe produrre in output il seguente programma? E se invece inserisse `100`?

```go
package main

import "fmt"

func main() {

	var n int

	fmt.Scan(&n)

	accumulatore := 1

	for i:=1; i<n; i++ {

		switch {
		case i<5:
			accumulatore *= i
		case i<10:
			accumulatore += i
		}

	}

	fmt.Println("Accumulatore =", accumulatore)
}
```
 
## 4 Soluzione di equazioni di secondo grado

Scrivere un programma che legga da **standard input** tre numeri reali, `a`, `b` e `c`, che rappresentano i coefficienti di un'equazione di secondo grado (con incognita `x`) espressa nella forma `ax² + bx + c = 0`.
Il programma deve calcolare e stampare a video il valore delle radici dell'equazione.

*Suggerimento:* Nel caso in cui l'equazione non ammetta soluzioni reali, potete memorizzare i valori delle radici in variabili di tipo `complex`. 

**Nota:** Si utilizzi il costrutto `switch case`.

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
