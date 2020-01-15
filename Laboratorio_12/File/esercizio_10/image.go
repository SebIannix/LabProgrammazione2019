package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	fInput, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Errore in apertura del file:", err)
		return
	}
	defer fInput.Close()

	immagine, err := LeggiImmagine(fInput)
	if err != nil {
		fmt.Println("Errore in lettura immagine:", err)
		return
	}

	fmt.Println("Immagine letta:", immagine)
}

type RGB struct {
	r, g, b uint
}

type Immagine struct {
	larghezza, altezza uint
	valori []RGB
}

func LeggiImmagine(file *os.File) (immagine Immagine, err error) {
	var riga string
	// scarto la prima riga
	if _, err = fmt.Fscanf(file, "%s", &riga); err != nil {
		return
	}
	// leggo larghezza e altezza dalla seconda riga
	if _, err = fmt.Fscanln(file, &immagine.larghezza, &immagine.altezza); err != nil {
		return
	}

	// scarto la terza riga
	if _, err = fmt.Fscanf(file, "%s", &riga); err != nil {
		return
	}

	// leggo i valori RGB
	immagine.valori = make([]RGB, immagine.larghezza * immagine.altezza)
	for i := range immagine.valori {
		_, err = fmt.Fscan(file, &immagine.valori[i].r, &immagine.valori[i].g, &immagine.valori[i].b)
		if err == io.EOF {
			break
		}
		if err != nil {
			return
		}
	}
	return immagine, nil
}


