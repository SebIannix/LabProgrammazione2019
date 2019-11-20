package main

import "fmt"

func main() {

	var dimensione int
	fmt.Scan(&dimensione)

	for riga := 0; riga < dimensione; riga++ {
		for colonna := 0; colonna < dimensione; colonna ++ {
      if riga%2 == 0 {
        if colonna%2 == 0 {
          // se riga pari e colonna pari
          fmt.Print("*")
        } else {
          // se riga pari e colonna dispari
          fmt.Print("+")
        }
      } else {
        if colonna%2 == 0 {
          // se riga dispari e colonna pari
          fmt.Print("+")
        } else {
          // se riga dispari e colonna dispari
          fmt.Print("*")
        }
      }
		}
		fmt.Println()
	}


}
