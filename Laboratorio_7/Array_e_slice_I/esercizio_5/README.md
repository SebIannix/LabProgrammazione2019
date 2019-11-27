# Qual è l'outout?

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