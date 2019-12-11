# Laboratorio 9 - Strutture
## 1 Qual è l'output?

Qual è l'output del seguente programma?

```go
package main

import (
	"fmt"
)

type Persona struct {
	Nome    string
	Cognome string
}

func main() {

	var p1, p2, p5 Persona
    // ... equivalente a:
	//p1 = Persona{}
	//p2 = Persona{}
	//p5 = Persona{}

	p1 = Persona{"Rick", "Sanchez"}

	p2.Nome = "Mario"
	p2.Cognome = "Rossi"

	p3 := Persona{"Morty", "Smith"}
	p4 := Persona{Cognome: "Bianchi", Nome: "Paolo"}

	p6 := Persona{Cognome: "Rossi"}

	fmt.Println("p1:", p1)
	fmt.Println("p5:", p5)

	nome := "Giuseppe"
	cognome := "Verdi"

	p5 = Persona{Cognome: cognome, Nome: nome}
    // ... equivalente a:
    // p5 = Persona{nome, cognome}

	fmt.Printf("Nome: %-10s - Cognome: %-10s\n", p1.Nome, p1.Cognome)
	fmt.Printf("%T\t%v\n", p3, p3)
	fmt.Printf("%#v\n", p4)
	fmt.Printf("%#v\n", p6)

}
```

## 2 Qual è l'output?

Qual è l'output del seguente programma?

```go
package main

import (
	"fmt"
)

type Persona struct {
	Nome    string
	Cognome string
}

func main() {

	anagrafica := []Persona{
		{"Rick", "Sanchez"},
		{"Morty", "Smith"},
		{"Jerry", "Smith"},
		{"Summer", "Smith"},
		{"Beth", "Smith"}}

	fmt.Printf("Tipo anagrafica %T\n", anagrafica)
	fmt.Println()

	for _, p := range anagrafica {
		fmt.Println(p.Nome, p.Cognome)
	}
	fmt.Println()

	CambiaPrimoCognome(anagrafica)

	for _, p := range anagrafica {
		fmt.Println(p.Nome, p.Cognome)
	}
}

func CambiaPrimoCognome(a []Persona){
	a[0].Cognome = "Martinez"
}
```

## 3 Rubrica (1)

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




## 4 Rubrica (2)

Si consideri una rubrica in cui:
* Ogni contatto deve avere un cognome, un nome, un indirizzo ed un numero di telefono.
* Ogni indirizzo deve contenere le seguenti informazioni: via, numero civico, CAP e città.
* Non possono esistere due contatti con lo stesso cognome e lo stesso nome.

La rubrica può essere gestita tramite le seguenti operazioni: 

1. I;*cognome*;*nome*;*via*;*numero_civico*;*CAP*;*città*;*numero_di_telefono*

   Un nuovo contatto specificato dai parametri: 

   *cognome* *nome* *via* *numero_civico* *CAP* *città* *numero_di_telefono*

   deve essere **inserito** nella rubrica.

   Se i parametri *cognome* e *nome* identificano un contatto già esistente all'interno della rubrica,  i dati relativi al contatto devono essere **aggiornati** come specificato dai parametri:
 
   *via* *numero_civico* *CAP* *città* *numero_di_telefono*

2. E;*cognome*;*nome*

   Il contatto identificato dai parametri: 

   *cognome* *nome*

   deve essere **eliminato** dalla rubrica.

   Se i parametri *cognome* e *nome* non identificano un contatto esistente all'interno della rubrica, l'operazione non ha alcun effetto sulla rubrica.

3. S

   I contatti presenti nella rubrica devono essere **stampati** a video.

4. A;*cognome*;*nome*;*via*;*numero_civico*;*CAP*;*città*;*numero_di_telefono*

   I dati relativi al contatto identificato dai parametri: 

   *cognome* *nome*

   devono essere **aggiornati** come specificato dai parametri:

   *via* *numero_civico* *CAP* *città* *numero_di_telefono*.

   Se i parametri *cognome* e *nome* non identificano un contatto già esistente all'interno della rubrica, un nuovo contatto specificato dai parametri: 
 
   *cognome* *nome* *via* *numero_civico* *CAP* *città* *numero_di_telefono*

   deve essere **inserito** nella rubrica. 

Scrivere un programma che:
* legga da **standard input** un testo su più righe;
* termini la lettura quando viene inserita da standard input una riga vuota (`""`).

Ogni riga di testo è una stringa in uno dei seguenti possibili formati:

1. I;*cognome*;*nome*;*via*;*numero_civico*;*CAP*;*città*;*numero_di_telefono*
2. E;*cognome*;*nome*
3. S
4. A;*cognome*;*nome*;*via*;*numero_civico*;*CAP*;*città*;*numero_di_telefono*

Il programma deve gestire una rubrica eseguendo in sequenza le operazioni corrispondenti alle righe di testo letto.

Oltre alla funzione `main()`, devono essere definite ed utilizzate almeno le seguenti funzioni:

* una funzione `LeggiTesto() []string` che legge da **standard input** un testo su più righe e terminato da una riga vuota (`""`), restituendo un valore `[]string` in cui sono memorizzate le stringhe corrispondenti alle righe del testo letto;
* una funzione `Sottostringhe(s string) []string` che riceve in input un valore `string` nel parametro `s` e restituisce un valore `[]string` in cui sono memorizzate tutte le sottostringhe presenti in `s` separate da un carattere `;` (si faccia riferimento alla documentazione della funzione `strings.Split` del package `strings`);
* una funzione `NuovaRubrica() (r []Contatto)` che restituisce un valore `[]Contatto` nella varibile `r` che rappresenta una rubrica vuota; 
* una funzione `InserisciContatto(r []Contatto, cognome, nome string, via string, numero uint, cap, città string, telefono string) []Contatto` che, in base alle modalità previste al punto `1`, inserisce nella rubrica `r` un'istanza del tipo `Contatto` inizializzata in base ai valori dei parametri `cognome`, `nome`, `via`, `numero`, `CAP`, `città`, `telefono`;
* una funzione `EliminaContatto(r []Contatto, cognome, nome string) []Contatto` che, in base alle modalità previste al punto `2`, elimina dalla rubrica `r` l'istanza del tipo `Contatto` i cui campi `Cognome` e `Nome` hanno lo stesso valore di quello dei parametri `cognome` e `nome`;
* una funzione `AggiornaContatto(rubrica []Contatto, cognome, nome string, via string, numero uint, cap, città string, telefono string) []Contatto` che, secondo le modalità previste al punto `4`, aggiorna in base ai valori dei parametri `via`, `numero`, `CAP`, `città`, `telefono` l'istanza del tipo `Contatto` presente nella rubrica `r` i cui campi `Cognome` e `Nome` hanno lo stesso valore di quello dei parametri `cognome` e `nome`.

##### Esempio d'esecuzione:
```text
$ go run rubrica.go
I;Mario;Rossi;Via Celoria;18;20122;Milano;02503111
S
I;Elena;Bianchi;Via Celoria;18;20122;Milano;02503111
S
E;Mario;Rossi;
S
A;Elena;Bianchi;Via Festa del perdono;7;20122;Milano;02503111
S

Lista dei 1 contatti
[1] - Mario Rossi, Tel. 02503111, Via Celoria 18, 20122, Milano
Lista dei 2 contatti
[1] - Mario Rossi, Tel. 02503111, Via Celoria 18, 20122, Milano
[2] - Elena Bianchi, Tel. 02503111, Via Celoria 18, 20122, Milano
Lista dei 1 contatti
[1] - Elena Bianchi, Tel. 02503111, Via Celoria 18, 20122, Milano
Lista dei 1 contatti
[1] - Elena Bianchi, Tel. 02503111, Via Festa del perdono 7, 20122, Milano
```
## 5 Rubrica (3)

Si consideri una rubrica in cui:
* Ogni contatto deve avere un cognome, un nome, un indirizzo ed un numero di telefono.
* Ogni indirizzo deve contenere le seguenti informazioni: via, numero civico, CAP e città.
* Non possono esistere due contatti con lo stesso cognome e lo stesso nome.

La rubrica può essere gestita tramite le seguenti operazioni: 

1. I;*cognome*;*nome*;*via*;*numero_civico*;*CAP*;*città*;*numero_di_telefono*

   Un nuovo contatto specificato dai parametri: 

   *cognome* *nome* *via* *numero_civico* *CAP* *città* *numero_di_telefono*

   deve essere **inserito** nella rubrica.

   Se i parametri *cognome* e *nome* identificano un contatto già esistente all'interno della rubrica,  i dati relativi al contatto devono essere **aggiornati** come specificato dai parametri:
 
   *via* *numero_civico* *CAP* *città* *numero_di_telefono*

2. E;*cognome*;*nome*

   Il contatto identificato dai parametri: 

   *cognome* *nome*

   deve essere **eliminato** dalla rubrica.

   Se i parametri *cognome* e *nome* non identificano un contatto esistente all'interno della rubrica, l'operazione non ha alcun effetto sulla rubrica.

3. S

   I contatti presenti nella rubrica devono essere **stampati** a video.

4. A;*cognome*;*nome*;*via*;*numero_civico*;*CAP*;*città*;*numero_di_telefono*

   I dati relativi al contatto identificato dai parametri: 

   *cognome* *nome*

   devono essere **aggiornati** come specificato dai parametri:

   *via* *numero_civico* *CAP* *città* *numero_di_telefono*.

   Se i parametri *cognome* e *nome* non identificano un contatto già esistente all'interno della rubrica, un nuovo contatto specificato dai parametri: 
 
   *cognome* *nome* *via* *numero_civico* *CAP* *città* *numero_di_telefono*

   deve essere **inserito** nella rubrica. 

Scrivere un programma che:
* legga da **standard input** un testo su più righe;
* termini la lettura quando viene inserita da standard input una riga vuota (`""`).

Ogni riga di testo è una stringa in uno dei seguenti possibili formati:

1. I;*cognome*;*nome*;*via*;*numero_civico*;*CAP*;*città*;*numero_di_telefono*
2. E;*cognome*;*nome*
3. S
4. A;*cognome*;*nome*;*via*;*numero_civico*;*CAP*;*città*;*numero_di_telefono*

Il programma deve gestire una rubrica eseguendo in sequenza le operazioni corrispondenti alle righe di testo letto.

Oltre alla funzione `main()`, devono essere definite ed utilizzate almeno le seguenti funzioni:

* una funzione `NuovaRubrica() (r map[string]Contatto)` che restituisce un valore `map[string]Contatto` nella variabile `r` che rappresenta una rubrica vuota (un'istanza del tipo `Contatto` può essere indicizzata in `r` utilizzando il valore chiave `cognome + "-" + nome`); 
* una funzione `InserisciContatto(r map[string]Contatto, cognome, nome string, via string, numero uint, cap, città string, telefono string) map[string]Contatto` che, in base alle modalità previste al punto `1`, inserisce nella rubrica `r` un'istanza del tipo `Contatto` inizializzata in base ai valori dei parametri `cognome`, `nome`, `via`, `numero`, `CAP`, `città`, `telefono`;
* una funzione `EliminaContatto(r map[string]Contatto, cognome, nome string) map[string]Contatto` che, in base alle modalità previste al punto `2`, elimina dalla rubrica `r` l'istanza del tipo `Contatto` i cui campi `Cognome` e `Nome` hanno lo stesso valore di quello dei parametri `cognome` e `nome`;
* una funzione `StampaRubrica(r map[string]Contatto)` che stampa a video le istanze del tipo `Contatto` presenti all'interno della rubrica `r`;
* una funzione `AggiornaContatto(rubrica map[string]Contatto, cognome, nome string, via string, numero uint, cap, città string, telefono string) map[string]Contatto` che, secondo le modalità previste al punto `4`, aggiorna in base ai valori dei parametri `via`, `numero`, `CAP`, `città`, `telefono` l'istanza del tipo `Contatto` presente nella rubrica `r` i cui campi `Cognome` e `Nome` hanno lo stesso valore di quello dei parametri `cognome` e `nome`.

##### Esempio d'esecuzione:
```text
$ go run rubrica.go
I;Mario;Rossi;Via Celoria;18;20122;Milano;02503111
S
I;Elena;Bianchi;Via Celoria;18;20122;Milano;02503111
S
E;Mario;Rossi;
S
A;Elena;Bianchi;Via Festa del perdono;7;20122;Milano;02503111
S

Lista dei 1 contatti
[1] - Mario Rossi, Tel. 02503111, Via Celoria 18, 20122, Milano
Lista dei 2 contatti
[1] - Mario Rossi, Tel. 02503111, Via Celoria 18, 20122, Milano
[2] - Elena Bianchi, Tel. 02503111, Via Celoria 18, 20122, Milano
Lista dei 1 contatti
[1] - Elena Bianchi, Tel. 02503111, Via Celoria 18, 20122, Milano
Lista dei 1 contatti
[1] - Elena Bianchi, Tel. 02503111, Via Festa del perdono 7, 20122, Milano
```
