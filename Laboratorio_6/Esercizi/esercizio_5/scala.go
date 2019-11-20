package main

import (
	"fmt"
)

func main() {

	var numeroGradini int
	fmt.Scan(&numeroGradini)

	if numeroGradini>0 {
		for gradino := 0; gradino < numeroGradini; gradino++ {
			StampaGradino(gradino)
		}
	} else {
		fmt.Println("Dimensione non sufficiente")
	}

}

// Stampa un gradino traslato di gradinoAttuale*2 spazi
func StampaGradino(gradinoAttuale int) {
  for i:=0; i<gradinoAttuale; i++ {
    fmt.Print("  ")
  }
	fmt.Println("***")
  for i:=0; i<gradinoAttuale; i++ {
    fmt.Print("  ")
  }
	fmt.Println("  *")
}
