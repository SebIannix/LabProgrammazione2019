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
