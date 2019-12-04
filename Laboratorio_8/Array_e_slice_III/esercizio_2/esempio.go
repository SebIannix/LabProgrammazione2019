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
	fmt.Println("v)\n a:", a, len(a), cap(a), "\n b:", b, len(b), cap(b), "\n c:", c, len(c), cap(c))
	a = a[:cap(a)]
	b = b[:cap(b)]
	c = c[:cap(c)]
	fmt.Println("vi)\n a:", a, len(a), cap(a), "\n b:", b, len(b), cap(b), "\n c:", c, len(c), cap(c))
}
