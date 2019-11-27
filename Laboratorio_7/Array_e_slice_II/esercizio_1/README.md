# Qual è l'outout?

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