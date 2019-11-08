package main

import "fmt"

func main() {

	var n int
	var stringaInput string

	fmt.Scan(&n, &stringaInput)

	for i := 0; i < n; i++ {
		// ad ogni iterazione del ciclo for stampo la stringa e il carattere '-'
		fmt.Print(stringaInput)
		// per evitare che venga stampato un '-' finale basta controllare
		// di non essere all'ultima iterazione
		if i != n-1 {
			// il carattere '-' viene stampato sempre tranne che all'ultima iterazione
			fmt.Print("-")
		}
	}

	fmt.Println()
}
