# Laboratorio 7 - Array e slice I
## 1 Qual è l'output?

Qual è l'output del seguente codice?

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {

	var numeroIntero int = 10

	fmt.Printf("Tipo: %T\n", numeroIntero)
	fmt.Printf("- formato default: %v\n", numeroIntero)
	fmt.Printf("- formato base 10: %d\n", numeroIntero)
	fmt.Printf("- formato base 2: %b\n", numeroIntero)

	fmt.Println()

	var numeroReale float64 = 14.5
	fmt.Printf("Tipo: %T\n", numeroReale)
	fmt.Printf("- formato default: %v\n", numeroReale)
	fmt.Printf("- formato con decimali: %f\n", numeroReale)
	fmt.Printf("- formato con numero fissato di decimali: %.2f\n", numeroReale)
	fmt.Printf("- formato esponenziale: %e\n", numeroReale)

	fmt.Println()

	var valoreLogico bool = false
	fmt.Printf("Tipo: %T\n", valoreLogico)
	fmt.Printf("- formato default: %v\n", valoreLogico)
	fmt.Printf("- formato true/false: %t\n", valoreLogico)

	fmt.Println()

	var carattere rune = 'A'
	fmt.Printf("Tipo: %T\n", carattere)
	fmt.Printf("- formato default: %v\n", carattere)
	fmt.Printf("- formato carattere: %c\n", carattere)
	fmt.Printf("- formato intero: %d\n", carattere)
	fmt.Printf("- formato unicode: %U\n", carattere)

	fmt.Println()

	var stringaTesto string = "Hello\tworld!"
	fmt.Printf("Tipo: %T\n", stringaTesto)
	fmt.Printf("- formato default: %v\n", stringaTesto)
	fmt.Printf("- formato stringa: %s\n", stringaTesto)

	fmt.Println()

	var posizioniDecimali = 2
	fmt.Printf("Printf con formato creato dinamicamente %." + 
            strconv.Itoa(posizioniDecimali) + "f\n", 10.458)
}
```
## 2 Qual è l'output?

Qual è l'output di questo programma?

```go
package main

import "fmt"

func main() {
	var a = [4]int{1, 2, 3, 4}
    var b = [...]rune{'a', 'b', 'c'}
	
    fmt.Println(a)

	for i:=0; i<len(a); i++ {
		fmt.Println("Indice", i, " - Valore:", a[i])
	}

	fmt.Println()

	for i, v := range b {
		fmt.Println("Indice", i, " - Valore:", v)
	}

}
```
## 3 Qual è l'output?

Qual è l'output di questo programma?

```go
package main

import "fmt"

func main() {
	a := [...]string{5:"E", 3:"C"}

	fmt.Println(a)

	for i:=len(a)-1; i>=0; i-- {
		fmt.Println("Indice", i, " - Valore:", a[i])
	}
	fmt.Println()

	for i, _ := range a {
		fmt.Println("Indice", len(a)-1-i, " - Valore:", a[len(a)-1-i])
	}
	fmt.Println()

	for i, _ := range a {
		fmt.Println("Indice", i, " - Valore:", a[i])
	}
	fmt.Println()

	for i := range a {
		fmt.Println("Indice", i, " - Valore:", a[i])
	}

}
```
## 4 Qual è l'outout?

Qual è l'output di questo programma?

```go
package main

import "fmt"

func main() {
	var a [6]int
    
    for i:=0;i<len(a);i++{
        a[i] = i*2
    }   

    fmt.Printf("a - %T: %v\n",a,a)

    b := a

    fmt.Printf("b - %T: %v\n",b,b)

	sl1 := a[:]    
    sl2 := a[2:]

    fmt.Printf("\na - %v\n",a)
    fmt.Printf("sl1 - %T: %v\n",sl1,sl1)

    sl1[0] = -9
    sl2[0] = 100

    fmt.Printf("a - %v\n",a)

    fmt.Printf("\nb - %v\n",b)

	for _, v := range b {
		v *= 2
	}
    
    fmt.Printf("b - %v\n",b)
}
```
## 5 Qual è l'outout?

Qual è l'output di questo programma?

```go
package main

import "fmt"

func main() {
	sl := []int{1, 2, 3, 4, 5, 6, 7}

	fmt.Printf("sl - %T: %v\n", sl, sl)

	sl1 := sl[:]
	sl2 := sl[1:3]

	fmt.Printf("len(sl1) = %d, cap(sl1) = %d\n", len(sl1), cap(sl1))
	fmt.Printf("sl1 - %T: %v\n", sl1, sl1)
	fmt.Printf("len(sl2) = %d, cap(sl2) = %d\n", len(sl2), cap(sl2))
	fmt.Printf("sl2 - %T: %v\n", sl2, sl2)

	sl2 = sl2[:len(sl2)+1]   // reslicing
	fmt.Printf("\nsl2 - %T: %v\n", sl2, sl2)

	sl2 = sl2[:cap(sl2)]
	fmt.Printf("sl2 - %T: %v\n", sl2, sl2)

	sl1 = sl1[1:]
	fmt.Printf("\nsl1 - %T: %v\n", sl1, sl1)

	/* una slice s non può essere modificata per accedere ad elementi
	   dell'array (a cui si riferisce) che precedono quello contenuto in
	   s[0]; l’istruzione s = s[-1:] genera un errore */

}
```
## 6 Qual è l'outout?

Qual è l'output di questo programma?

```go
package main

import "fmt"
import "math"

func main() {
    
    var n int = 5

    //a := [n]int   /* non corretto */ 
 
    sl := make([]int, n)

    for i:=0;i<len(sl);i++{
         sl[i] = i
    }      

    fmt.Printf("%T: %v\n", sl, sl)
}
```
## 7 Stampa in ordine inverso


Scrivere un programma che, dopo aver letto da **standard input** un numero intero `n`, chiede all'utente di inserire `n` numeri interi (sempre da **standard input**). 

Il programma deve stampare gli `n` numeri interi in ordine inverso rispetto a quello di inserimento.

*Nota:* Per creare dinamicamente una slice, si utilizzi la funzione `make`.

##### Esempio d'esecuzione:

```text
$ go run apmats.go 
9
Inserisci 9 numeri:
1 -12 3 -4 5 -6 7 -7 9
Numeri in ordine inverso:
9 -7 7 -6 5 -4 3 -12 1 

$ go run apmats.go 
5
Inserisci 5 numeri:
1 2 3 4 5
Numeri in ordine inverso:
5 4 3 2 1 
```
