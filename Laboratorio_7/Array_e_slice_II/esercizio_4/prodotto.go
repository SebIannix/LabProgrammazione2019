package main

import "fmt"
import "os"
import "strconv"

func main() {

	prodotto := 1
    // Converto ogni valore presente in os.Args
    // ma escludento os.Args[0], che è il nome
    // dell'eseguibile
	for _, v := range os.Args[1:] {
		
		if n, err := strconv.Atoi(v); err == nil {
            // Se riesco a convertire il valore di v (string) in un
            // intero senza errori (err==nil), allora posso usare
            // il valore di n per calcolare il prodotto
			prodotto *= n
		}

	}

	fmt.Println("Il risultato della moltiplicazione tra i numeri interi è", prodotto)

}
