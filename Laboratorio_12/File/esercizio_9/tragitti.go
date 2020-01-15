package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Punto struct {
	etichetta string
	x, y      float64
}

type Tratta struct {
	p1, p2 Punto
}

type TragittoPerTratte struct {
	id        int
	tratte    []Tratta
	lunghezza float64
}

type TragittoPerPunti struct {
	id        int
	punti     []Punto
	lunghezza float64
}

func main() {

	filePunti, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Errore durante l'apertura del file:", err)
		return
	}
	defer filePunti.Close()

	fileTragitti, err := os.Open(os.Args[2])
	if err != nil {
		fmt.Println("Errore durante l'apertura del file:", err)
		return
	}
	defer fileTragitti.Close()

	punti, err := LeggiPunti(filePunti)
	if err != nil {
		fmt.Println("Errore durante la lettura dei punti dal file:", err)
		return
	}

	tragitti, err := LeggiTragitti(fileTragitti, punti)
	if err != nil {
		fmt.Println("Errore durante la lettura dei tragitti dal file:", err)
		return
	}

	fmt.Printf("Tragitto di lunghezza massima:\n%s\n", String(*TragittoMassimo(tragitti)))
}

func LeggiPunti(fileInput *os.File) (punti map[string]Punto, err error) {
	punti = make(map[string]Punto)

	scanner := bufio.NewScanner(fileInput)
	for scanner.Scan() {
		rigaSeparata := strings.Split(scanner.Text(), ";")
		etichetta := rigaSeparata[0]
		x, _ := strconv.ParseFloat(rigaSeparata[1], 64)
		y, _ := strconv.ParseFloat(rigaSeparata[2], 64)

		punti[etichetta] = Punto{etichetta, x, y}
	}
	err = scanner.Err()

	return
}

func LeggiTragitti(fileInput *os.File, punti map[string]Punto) (tragitti map[int]*TragittoPerTratte, err error) {
	tragitti = make(map[int]*TragittoPerTratte)

	scanner := bufio.NewScanner(fileInput)
	for scanner.Scan() {
		rigaSeparata := strings.Split(scanner.Text(), ";")
		idTragitto, _ := strconv.Atoi(rigaSeparata[0])

		etichettaP1 := rigaSeparata[1]
		etichettaP2 := rigaSeparata[2]

		var tragitto *TragittoPerTratte
		var ok bool
		if tragitto, ok = tragitti[idTragitto]; !ok {
			tragitto = NuovoTragittoPerTratte(idTragitto)
			tragitti[idTragitto] = tragitto
			AggiungiTratta(tragitto, punti[etichettaP1], punti[etichettaP2])
		} else {
			AggiungiTratta(tragitto, punti[etichettaP1], punti[etichettaP2])
		}
	}
	err = scanner.Err()

	return
}

func NuovoTragittoPerTratte(id int) (t *TragittoPerTratte) {
	return &TragittoPerTratte{id, nil, 0}
}

func Distanza(p1, p2 Punto) float64 {
	deltaX := p1.x - p2.x
	deltaY := p1.y - p2.y
	return math.Sqrt(deltaX*deltaX + deltaY*deltaY)
}

func AggiungiTratta(t *TragittoPerTratte, p1 Punto, p2 Punto) {
	t.tratte = append(t.tratte, Tratta{p1, p2})
	t.lunghezza += Distanza(p1, p2)
}

func NuovoTragittoPerPunti(TpT *TragittoPerTratte) (TpP *TragittoPerPunti) {

	TpP = &TragittoPerPunti{TpT.id, nil, TpT.lunghezza}

	TpP.punti = append(TpP.punti, TpT.tratte[0].p1)
	TpP.punti = append(TpP.punti, TpT.tratte[0].p2)
	var trattaDaConsiderare []Tratta
	for _, tratta := range TpT.tratte[1:] {
		trattaDaConsiderare = append(trattaDaConsiderare, tratta)
	}

	for len(trattaDaConsiderare) > 0 {

		for i, tratta := range trattaDaConsiderare {
			if TpP.punti[0].etichetta == tratta.p2.etichetta {
				TpP.punti = append([]Punto{tratta.p1}, TpP.punti...)
				trattaDaConsiderare = append(trattaDaConsiderare[:i], trattaDaConsiderare[i+1:]...)
				break
			}
			if TpP.punti[len(TpP.punti)-1].etichetta == tratta.p1.etichetta {
				TpP.punti = append(TpP.punti, tratta.p2)
				trattaDaConsiderare = append(trattaDaConsiderare[:i], trattaDaConsiderare[i+1:]...)
				break
			}
		}

	}

	return

}

func String(t TragittoPerPunti) (stringa string) {
	stringa = fmt.Sprintf("Tragitto %d - Lunghezza %v ", t.id, t.lunghezza)
	for _, p := range t.punti {
		stringa += fmt.Sprintf("%s(%v, %v) ", p.etichetta, p.x, p.y)
	}
	return
}

func TragittoMassimo(tragitti map[int]*TragittoPerTratte) (tragitto *TragittoPerPunti) {
	for _, t := range tragitti {
		fmt.Printf("%s\n",String(*NuovoTragittoPerPunti(t)))
		if tragitto == nil || t.lunghezza > tragitto.lunghezza {
			tragitto = NuovoTragittoPerPunti(t)
		}
	}
	return
}
