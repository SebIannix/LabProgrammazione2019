package main

import (
	"fmt"
)

func main() {

	vocaliMaiuscole, vocaliMinuscole := 0, 0
	consonantiMaiuscole, consonantiMinuscole := 0, 0

	var stringaInput string

	fmt.Scan(&stringaInput)

	for _, lettera := range stringaInput {

		if (lettera >= 'A' && lettera <= 'Z') || (lettera >= 'a' && lettera <= 'z') {
			if lettera >= 'A' && lettera <= 'Z' {
				if lettera == 'A' || lettera == 'E' || lettera == 'I' || lettera == 'O' || lettera == 'U' {
					vocaliMaiuscole++
				} else {
					consonantiMaiuscole++
				}
			} else {
				if lettera == 'a' || lettera == 'e' || lettera == 'i' || lettera == 'o' || lettera == 'u' {
					vocaliMinuscole++
				} else {
					consonantiMinuscole++
				}
			}

		}

	}
	fmt.Println("Vocali maiuscole:", vocaliMaiuscole)
	fmt.Println("Consonanti maiuscole:", consonantiMaiuscole)
	fmt.Println("Vocali minuscole:", vocaliMinuscole)
	fmt.Println("Consonanti minuscole:", consonantiMinuscole)
}
