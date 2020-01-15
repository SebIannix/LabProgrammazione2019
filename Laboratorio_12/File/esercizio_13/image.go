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

	conversione := Converti(immagine)

	fOutput, err := os.Create(os.Args[2])
	if err != nil {
		fmt.Println("Errore in apertura del file:", err)
		return
	}
	defer fOutput.Close()

	err = ScriviImmagine(fOutput, conversione)
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

func MediaRGB(valore RGB) RGB {
	media := (valore.r + valore.g + valore.b) / 3
	return RGB{media, media, media}
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

func Converti(immagine Immagine) (conversione Immagine) {
	conversione.larghezza = immagine.larghezza
	conversione.altezza = immagine.altezza

	conversione.valori = make([]RGB, len(immagine.valori))
	for i, valoreRGB := range immagine.valori {
		conversione.valori[i] = MediaRGB(valoreRGB)

	}
	return
}

