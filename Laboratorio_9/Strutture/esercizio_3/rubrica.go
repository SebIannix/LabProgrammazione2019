package main

import "fmt"

func main() {

	rubrica := NuovaRubrica()

	rubrica = InserisciContatto(rubrica, "Mario", "Rossi",
		"Via Festa del Perdono", 11, "20122", "Milano", "02503111")
	rubrica = InserisciContatto(rubrica, "Anna", "Rossi",
		"Via Festa del Perdono", 11, "20122", "Milano", "02503111")
	rubrica = InserisciContatto(rubrica, "Carlo", "Rossi",
		"Via Festa del Perdono", 11, "20122", "Milano", "02503111")

	StampaRubrica(rubrica)

	rubrica = EliminaContatto(rubrica, "Mario", "Rossi")

	StampaRubrica(rubrica)

	rubrica = AggiornaContatto(rubrica, "Anna", "Rossi", "Via S. Sofia", 25, "20122", "Milano", "")

	StampaRubrica(rubrica)
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

