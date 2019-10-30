package main

import (
	"fmt"
)

func main() {
	numeriLetti := 0
	mAritmetica := 0.0

	fmt.Print("Inserisci una sequenza di numeri (interrompi con numero<=0): ")
	for {
		var x float64
		fmt.Scan(&x)

		if x > 0 {
			numeriLetti++
			mAritmetica += x
		} else {
			break
		}
	}

	mAritmetica /= float64(numeriLetti)

	fmt.Println("Media aritmetica:", mAritmetica)

}
