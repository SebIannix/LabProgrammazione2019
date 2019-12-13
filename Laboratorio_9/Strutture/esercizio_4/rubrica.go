package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	listaComandi := LeggiTesto()
	rubrica := NuovaRubrica()

	for _, comando := range listaComandi {
		rubrica = InterpreteComando(comando, rubrica)
	}
}

func LeggiTesto() (testo []string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		riga := scanner.Text()
		if riga == "" {
			break
		}
		testo = append(testo, riga)
	}
	return
}

func Sottostringhe(s string) (sottostringhe []string) {
	return strings.Split(s, ";")
}

func InterpreteComando(comando string, rubrica []Contatto) []Contatto {
	tokenComando := Sottostringhe(comando)
	switch tokenComando[0] {
	case "I":
		numero, _ := strconv.Atoi(tokenComando[4])
		rubrica = InserisciContatto(rubrica, tokenComando[1], tokenComando[2], tokenComando[3], uint(numero), tokenComando[5], tokenComando[6], tokenComando[7])
	case "E":
		rubrica = EliminaContatto(rubrica, tokenComando[1], tokenComando[2])
	case "S":
		StampaRubrica(rubrica)
	case "A":
		numero, _ := strconv.Atoi(tokenComando[4])
		rubrica = AggiornaContatto(rubrica, tokenComando[1], tokenComando[2], tokenComando[3], uint(numero), tokenComando[5], tokenComando[6], tokenComando[7])
	}
	return rubrica
}

type Indirizzo struct {
	Via string
	Numero  uint
	Cap, Città  string
}

type Contatto struct {
	Nome, Cognome   string
	Indirizzo   Indirizzo
	Telefono    string
}

func NuovaRubrica() []Contatto {
	return []Contatto{}
}

func StampaRubrica(rubrica []Contatto) {
	fmt.Println("Lista dei", len(rubrica), "contatti")
	for i, c := range rubrica {
		fmt.Printf("[%d] - %s %s, Tel. %s, %s %d, %s, %s\n", i+1, c.Nome, c.Cognome, c.Telefono, c.Indirizzo.Via,
			c.Indirizzo.Numero, c.Indirizzo.Cap, c.Indirizzo.Città)
	}
}

func InserisciContatto(rubrica []Contatto, nome, cognome, via string, numero uint, cap, città, telefono string) []Contatto {
	return append(rubrica, Contatto{nome, cognome,
		Indirizzo{via, numero, cap, città}, telefono})
}

func EliminaContatto(rubrica []Contatto, nome, cognome string) []Contatto {
	for i, contatto := range rubrica {
		if contatto.Nome == nome && contatto.Cognome == cognome {
			rubrica = append(rubrica[:i], rubrica[i+1:]...)
		}
	}
	return rubrica
}

func AggiornaContatto(rubrica []Contatto, nome, cognome, via string, numero uint, cap, città, telefono string) []Contatto {
	for i, contatto := range rubrica {
		if contatto.Nome == nome && contatto.Cognome == cognome {
			rubrica[i].Telefono = telefono
			rubrica[i].Indirizzo = Indirizzo{via, numero, cap, città}
		}
	}
	return rubrica
}

