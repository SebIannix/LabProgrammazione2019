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

func InterpreteComando(comando string, rubrica map[string]Contatto) map[string]Contatto {
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

func NuovaRubrica() map[string]Contatto {
	return map[string]Contatto{}
}

func StampaRubrica(rubrica map[string]Contatto) {
	fmt.Println("Lista dei", len(rubrica), "contatti")
	contatore := 1
	for _, c := range rubrica {
		fmt.Printf("[%d] - %s %s, Tel. %s, %s %d, %s, %s\n", contatore, c.Nome, c.Cognome, c.Telefono, c.Indirizzo.Via,
			c.Indirizzo.Numero, c.Indirizzo.Cap, c.Indirizzo.Città)
		contatore ++
	}
}

func InserisciContatto(rubrica map[string]Contatto, nome, cognome, via string, numero uint, cap, città, telefono string) map[string]Contatto {
	rubrica[Chiave(nome, cognome)] = Contatto{nome, cognome,
		Indirizzo{via, numero, cap, città}, telefono}
	return rubrica
}

func EliminaContatto(rubrica map[string]Contatto, nome, cognome string) map[string]Contatto {
	delete(rubrica, Chiave(nome, cognome))
	return rubrica
}

func AggiornaContatto(rubrica map[string]Contatto, nome, cognome, via string, numero uint, cap, città, telefono string) map[string]Contatto {
	chiave := Chiave(nome, cognome)
	if _, ok := rubrica[chiave]; ok {
		rubrica[chiave] = Contatto{nome, cognome,
			Indirizzo{via, numero, cap, città}, telefono}
	}
	return rubrica
}

func Chiave(nome, cognome string) string {
	return cognome+"-"+nome
}