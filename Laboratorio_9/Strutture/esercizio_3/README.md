# Rubrica (1)

Si consideri una rubrica in cui:
* Ogni contatto deve avere un cognome, un nome, un indirizzo ed un numero di telefono.
* Ogni indirizzo deve contenere le seguenti informazioni: via, numero civico, CAP e città.
* Non possono esistere due contatti con lo stesso cognome e lo stesso nome.

L'entità indirizzo può essere modellata con il tipo di dato:

```go
type Indirizzo struct {
	Via string
	Numero  uint
	Cap, Città  string
}
```

L'entità contatto può essere modellata con il tipo di dato:
```go
type Contatto struct {
	Cognome, Nome   string
	Indirizzo   Indirizzo
    Telefono    string
}
```

Implementare le funzioni:

* `NuovaRubrica() (r []Contatto)` che restituisce un valore `[]Contatto` nella varibile `r` che rappresenta una rubrica vuota (una rubrica in cui, inizialmente, non è presente alcun contatto); 
* `InserisciContatto(r []Contatto, cognome, nome string, via string, numero uint, cap, città string, telefono string) []Contatto` che inserisce nella rubrica `r` un'istanza del tipo `Contatto` inizializzata in base ai valori dei parametri `cognome`, `nome`, `via`, `numero`, `CAP`, `città`, `telefono`;
* `EliminaContatto(r []Contatto, cognome, nome string) []Contatto` che elimina dalla rubrica `r` l'istanza del tipo `Contatto` i cui campi `Cognome` e `Nome` hanno lo stesso valore di quello dei parametri `cognome` e `nome`;
* `StampaRubrica(r []Contatto)` che stampa a video le istanze del tipo `Contatto` presenti all'interno della rubrica `r`;
* `AggiornaContatto(rubrica []Contatto, cognome, nome string, via string, numero uint, cap, città string, telefono string) []Contatto` che aggiorna in base ai valori dei parametri `via`, `numero`, `CAP`, `città`, `telefono` l'istanza del tipo `Contatto` presente nella rubrica `r` i cui campi `Cognome` e `Nome` hanno lo stesso valore di quello dei parametri `cognome` e `nome`;

affinché la funzione `main()` di seguito riportata possa essere eseguita generando l'output atteso.

Funzione `main()`:
```go
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

	rubrica = AggiornaContatto(rubrica, "Anna", "Rossi", 
        "Via S. Sofia", 25, "20122", "Milano", "")

	StampaRubrica(rubrica)
}
```

Output atteso:
```text
Lista dei 3 contatti
[1] - Mario Rossi, Tel. 02503111, Via Festa del Perdono 11, 20122, Milano
[2] - Anna Rossi, Tel. 02503111, Via Festa del Perdono 11, 20122, Milano
[3] - Carlo Rossi, Tel. 02503111, Via Festa del Perdono 11, 20122, Milano

Lista dei 2 contatti
[1] - Anna Rossi, Tel. 02503111, Via Festa del Perdono 11, 20122, Milano
[2] - Carlo Rossi, Tel. 02503111, Via Festa del Perdono 11, 20122, Milano

Lista dei 2 contatti
[1] - Anna Rossi, Tel. , Via S. Sofia 25, 20122, Milano
[2] - Carlo Rossi, Tel. 02503111, Via Festa del Perdono 11, 20122, Milano
```



