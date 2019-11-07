package main

import "fmt"

func main() {
	var angolo1, angolo2 float64

	fmt.Print("Inserire le ampiezze dei due angoli: ")
	fmt.Scan(&angolo1, &angolo2)

	if sommaAmpiezze := angolo1 + angolo2; sommaAmpiezze < 180 {
		angolo3 := 180 - sommaAmpiezze
		fmt.Print("Ampiezza terzo angolo = ", angolo3, "°\n")
	} else {
		fmt.Println("I due angoli non appartengono ad un triangolo")
	}
}
