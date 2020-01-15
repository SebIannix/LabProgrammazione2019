package main

import (
	"fmt"
	"os"
)

func main() {


	immagine := Immagine{}

	immagine.larghezza = 4
	immagine.altezza = 2
	immagine.valori = []RGB{RGB{200, 100, 100}, RGB{100, 200, 200},
		RGB{240, 240, 240}, RGB{130, 130, 145}, RGB{100, 100, 100},
		RGB{80, 100, 120}, RGB{40, 60, 50}, RGB{10, 20, 30}}

	fOutput, err := os.Create(os.Args[1])
	if err != nil {
		fmt.Println("Errore in creazione del file:", err)
		return
	}
	defer fOutput.Close()

	err = ScriviImmagine(fOutput, immagine)
	if err != nil {
		fmt.Println("Errore durante la scrittura del file:", err)
		return
	}
}

type RGB struct {
	r, g, b uint
}

type Immagine struct {
	larghezza, altezza uint
	valori []RGB
}

func MassimoRGB(valori []RGB) (massimo uint) {
	for _, v := range valori {
		if v.r > massimo {
			massimo = v.r
		}
		if v.g > massimo {
			massimo = v.g
		}
		if v.b > massimo {
			massimo = v.b
		}
	}
	return
}


func ScriviImmagine(file *os.File, immagine Immagine) (err error) {
	if _, err = fmt.Fprintln(file, "P3"); err != nil {
		return
	}
	if _, err = fmt.Fprintf(file, "%d %d\n", immagine.larghezza, immagine.altezza); err != nil {
		return
	}
	if _, err = fmt.Fprintf(file, "%d\n", MassimoRGB(immagine.valori)); err != nil {
		return
	}

	for _, valoreRGB := range immagine.valori {
		if _, err = fmt.Fprintf(file, "%d %d %d\n", valoreRGB.r, valoreRGB.g, valoreRGB.b); err != nil {
			return
		}
	}

	return nil
}