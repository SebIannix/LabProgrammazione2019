# Laboratorio 5 - Funzioni
## 1 Qual è l'output?

Qual è l'output del seguente programma?

```go
package main

import "fmt"

func test(x int) (y int) {
	/* Le variabili 'a' e 'b' definite nel 'main' sono out of scope */
	/* La seguente istruzione genererebbe un errore di compilazione */
	//fmt.Println(a, b)
	
	fmt.Println("Funzione 'test' - Inizio.")
	y = x * 10
	fmt.Println("Funzione 'test' - Prima dell'ultima istruzione.")
	return
	//fmt.Println("Funzione 'test' - Dopo dell'ultima istruzione.")
}

func main() {
	fmt.Println("Funzione 'main' - Inizio.")
	var a, b int = 10, 5
	fmt.Println("Funzione 'main' - Valori di 'a' e 'b' prima di chiamare/invocare " +
		"la funzione 'test'.")
	fmt.Println(a, b)
	b = test(a)
	fmt.Println("Funzione 'main' - Valori di 'a' e 'b' dopo aver chiamato/invocato " +
		"la funzione 'test'.")
	fmt.Println(a, b)
	fmt.Println("Funzione 'main' - Fine.")	
}
```
## 2 Qual è l'output?

Qual è l'output del seguente programma?

```go
package main

import "fmt"

func test(x int) (y int) {
	y = x * 10
	x, y = y, x
	return
}

func main() {
	var a, b int = 10, 5
	b = test(a)
	fmt.Println(a, b)
}
```
## 3 Qual è l'output?

Qual è l'output del seguente programma?

```go
package main

import "fmt"

var a int = 10

func test1() int {
	a += 5
	return a
}

func test2(a int) int {
	a += 6
	return a
}

func test3() int {
	return a + 7
}

func main() {
	var a, b, c int
	a, b, c = test1(), test2(a), test3()
	fmt.Println(a, b, c)
}
```
## 4 Trova l'errore

Questo programma contiene degli errori. Corregere gli errori ed eseguire il programma.

```go
package main

import "fmt"

func test(x int) y int, z int
{
y := x + 1
z := x + 2
return
}

func main() {
	var a, b int
	a, b := test(10)
	fmt.Println(a, b)
}
```
## 5 Trova l'errore

Questo programma dovrebbe stampare `20 40` ma non genera l'output desiderato.
Corregere e verificare che l'esecuzione sia conforme al comportamento atteso.

```go
package main

import "fmt"

func test(x, y int) (int, int) {
	return 2*x, 2*y
}

func main() {
	var x, y int = 10, 20
	test(x, y)
	fmt.Println(x, y)
}
```
## 6 Testo invertito

Scrivere un programma che legga da **standard input** un testo formato da un numero variabile di righe, terminando la lettura quando viene inserita da **standard input** una riga vuota (`""`), e lo ristampi dall’ultimo carattere al primo.

Scrivere un programma che: 
* legga da **standard input** un testo su più righe;
* termini la lettura quando viene inserita da **standard input** una riga vuota (`""`);
* ristampi il testo letto (riga vuota esclusa) dall’ultimo carattere al primo.

Il programma deve essere organizzato nelle seguenti funzioni:
* una funzione `LeggiTesto() string` che legge da **standard input** un testo su più righe e terminato da una riga vuota (`""`), restituendo un valore (di tipo) `string` (una stringa) in cui è memorizzato il testo letto (riga vuota esclusa);
* una funzione `InvertiStringa(s string) string` che riceve in input un valore (di tipo) `string` (una stringa) nel parametro `s` e restituisce un valore (di tipo) `string` (una stringa) in cui il primo carattere è l'ultimo che definisce `s`, il secondo carattere è il penultimo che definisce `s`, ...  e l'ultimo carattere è il primo che definisce `s`;
* una funzione `main()` che richiama la funzione `LeggiTesto()` per leggere il testo da **standard input**, inverte il testo letto (inverte la stringa in cui è memorizzato il testo letto) richiamando opportunamente la funzione `InvertiStringa()` e stampa il testo invertito (la stringa invertita).

*Suggerimento:* Per leggere da **standard input** un testo formato da un numero variabile di righe (dove ogni riga è una stringa definita da una sequenza di caratteri arbitrari e terminata da '\n') terminando la lettura quando viene inserita da **standard input** una riga vuota (`""`), è possibile utilizzare il seguente programma.

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	stringa := ""
	fmt.Println("Inserisci una o più righe di testo:")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {

		stringaLocale := scanner.Text()
		if stringaLocale == "" {
			fmt.Print("stringaLocale == \"\": lettura terminata.\n")
			break
		} else {
			fmt.Print("stringaLocale == ",stringaLocale,"\n")
			stringa += stringaLocale + "\n"
		}

	}
	fmt.Println("Righe di testo lette: ")
	fmt.Print(stringa)
}
```

##### Esempio d'esecuzione:

```text
$ go run stringainvertita.go
Inserisci un testo su più righe (termina con riga vuota):
Testo di prova
disposto su due righe

Testo invertito:
ehgir eud us otsopsid
avorp id otseT
```
## 7 Linguaggio farfallino

Nel linguaggio farfallino ciascuna vocale non accentata (`vocale`) viene sostituita da una sequenza di tre caratteri `vocale-f-vocale`. Per esempio, la vocale `a` viene sostituita dalla sequenza `afa`, la vocale `e` dalla sequenza `efe` e così via. Se una vocale è maiuscola, anche la sequenza di tre caratteri che sostituisce la vocale deve essere definita da caratteri maiuscoli (ad esempio, la vocale `A` viene sostituita dalla sequenza `AFA`).

Scrivere un programma che: 
* legga da **standard input** un testo su più righe (alcune delle quali possono essere delle righe vuote (`""`));
* termini la lettura quando, premendo la combinazione di tasti `Ctrl+D`, viene inserito da **standard input** l'indicatore End-Of-File (EOF);
* ristampi il testo letto dopo averlo tradotto in linguaggio farfallino.

Il programma deve essere organizzato nelle seguenti funzioni:
* una funzione `LeggiTesto() string` che legge da **standard input** un testo su più righe (alcune delle quali possono essere delle righe vuote (`""`)) e terminato dall'indicatore EOF, restituendo un valore `string` in cui è memorizzato il testo letto;

* una funzione `TraduciCarattereInFarfallino(c rune) string` che riceve in input un valore `rune` nel parametro `c` e restituisce un valore `string` che rappresenta la traduzione in linguaggio farfallino di `c`;

* una funzione `TraduciTestoInFarfallino(s string) string` che riceve in input un valore `string` nel parametro `s` e restituisce un valore `string` che rappresenta la traduzione in linguaggio farfallino di `s`; la funzione deve utilizzare la funzione `TraduciCarattereInFarfallino()`;
* una funzione `main()` che richiama la funzione `LeggiTesto()` per leggere il testo da **standard input**, traduce in linguaggio farfallino il testo letto (richiamando opportunamente la funzione `TraduciTestoInFarfallino()`) e stampa il testo tradotto.

*Suggerimento:* Per leggere da **standard input** un testo formato da un numero variabile di righe (dove ogni riga è una stringa definita da una sequenza di caratteri arbitrari e terminata da '\n') terminando la lettura quando viene inserito da **standard input** l'indicatore End-Of-File (EOF), è possibile utilizzare il seguente programma.

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	stringa := ""
	fmt.Println("Inserisci una o più righe di testo:")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {

		stringaLocale := scanner.Text()
		fmt.Print("stringaLocale == ",stringaLocale,"\n")
		stringa += stringaLocale + "\n"

	}
	fmt.Println("Righe di testo lette: ")
	fmt.Print(stringa)
}
```

##### Esempio d'esecuzione:

```text
$ go run farfallino.go 
Inserisci un testo su più righe (termina con Ctrl+D):
Questo e' un testo di prova
da trasformare IN ALFABETO FARFALLINO
Risultato:
Qufuefestofo efe' ufun tefestofo difi profovafa
dafa trafasfoformafarefe IFIN AFALFAFABEFETOFO FAFARFAFALLIFINOFO
```

## 8 Numeri primi

Un numero naturale è primo se è divisibile solo per se stesso e per 1.

Scrivere un programma che legga da **standard input** un numero intero `soglia` e stampi tutti i numeri primi inferiori a `soglia`.

Il programma deve essere organizzato nelle seguenti funzioni:
* una funzione `LeggiNumero() int` che legge da **standard input** un numero intero e ne restituisce il valore;
* una funzione `ÈPrimo(n int) bool` che riceve in input un valore (di tipo) `int` nel parametro `n` e restituisce `true` se `n` è primo (i.e., se il valore di `n` rappresenta un numero primo) e `false` altrimenti;
* una funzione `NumeriPrimi(limite int)` che riceve in input un valore (di tipo) `int` nel parametro `limite` e stampa tutti i numeri primi inferiori a `limite` (i.e., al valore di `limite`); la funzione deve utilizzare la funzione `ÈPrimo()`;
* una funzione `main()` che richiama la funzione `LeggiNumero()` per leggere il numero intero `soglia`. Se `soglia > 0` allora richiama la funzione `NumeriPrimi()` in modo opportuno, altrimenti stampa un messaggio d'errore.

##### Esempio d'esecuzione:

```bash
$ go run numeri_primi.go
Inserisci un numero: -3
La soglia inserita non è positiva

$ go run numeri_primi.go
Inserisci un numero: 5
Numeri primi inferiori a 5
2 3 

$ go run numeri_primi.go
Inserisci un numero: 12
Numeri primi inferiori a 12
2 3 5 7 11
```

## 9 Numeri primi gemelli

**Definizione**: Due numeri primi `p` e `q` sono gemelli se `p = q + 2`. 

Scrivere un programma che legga da **standard input** un numero intero `soglia` e stampi tutti i numeri primi gemelli tali che `p` sia inferiore a `soglia`.

Il programma deve essere organizzato nelle seguenti funzioni:
* una funzione `LeggiNumero() int` che legge da **standard input** un numero intero e ne restituisce il valore;
* una funzione `ÈPrimo(n int) bool` che riceve in input un valore `int` nel parametro `n` e restituisce `true` se `n` è primo e `false` altrimenti;
* una funzione `NumeriPrimiGemelli(limite int)` che riceve in input un valore `int` nel parametro `limite` e stampa tutte le coppie di numeri primi gemelli tali che `p` sia inferiore a `limite` (cfr. Definizione); la funzione deve utilizzare la funzione `ÈPrimo()`;
* una funzione `main()` che richiama la funzione `LeggiNumero()` per leggere il numero intero `soglia`. Se `soglia > 0` allora richiama la funzione `NumeriPrimiGemelli()` in modo opportuno, altrimenti stampa un messaggio d'errore.

##### Esempio d'esecuzione:

```text
$ go run numeri_primi_gemelli.go
Inserisci un numero: -4
La soglia inserita non è positiva

$ go run numeri_primi_gemelli.go
Inserisci un numero: 10
Numeri primi gemelli inferiori a 10
(3,5)(5,7)

$ go run numeri_primi_gemelli.go
Inserisci un numero: 10
Numeri primi gemelli inferiori a 10
(3,5) (5,7) 

$ go run numeri_primi_gemelli.go
Inserisci un numero: 20
Numeri primi gemelli inferiori a 20
(3,5) (5,7) (11,13) (17,19)
```
## 10 Numeri perfetti

**Definizione**: Un numero naturale è perfetto se è uguale alla somma dei suoi divisori propri (per esempio, `6` è perfetto perché `6 = 1 + 2 + 3`).

Scrivere un programma che legga da **standard input** un numero intero e stampi se il numero è perfetto oppure no.

Il programma deve essere organizzato nelle seguenti funzioni:
* una funzione `LeggiNumero() int` che legge da **standard input** un numero intero e ne restituisce il valore;
* una funzione `SommaDivisoriPropri(n int) int` che riceve in input un valore (di tipo) `int` nel parametro `n` e restituisce un valore  (di tipo) `int` pari alla somma dei divisori propri di `n` (i.e., alla somma dei divisori propri del numero rappresentato dal valore di `n`), pari a `0` se (il numero rappresentato dal valore di) `n` non ha divisori propri;
* una funzione `ÈPerfetto(n int) bool` che riceve in input un valore  (di tipo) `int` nel parametro `n` e restituisce `true` se `n` è perfetto (i.e., se il valore di `n` rappresenta un numero perfetto) e `false` altrimenti; la funzione deve utilizzare la funzione `SommaDivisoriPropri()`;
* una funzione `main()` che richiama la funzione `LeggiNumero()` per leggere il numero intero `numero` da controllare. Se `numero > 0` allora richiama la funzione `ÈPerfetto()` in modo opportuno, altrimenti stampa un messaggio d'errore.

##### Esempio d'esecuzione:

```text
$ go run numero_perfetto.go
Inserisci un numero: 0
0 non è perfetto

$ go run numero_perfetto.go
Inserisci un numero: 1
1 non è perfetto

$ go run numero_perfetto.go
Inserisci un numero: 6
6 è perfetto

$ go run numero_perfetto.go
Inserisci un numero: 2
2 non è perfetto

$ go run numero_perfetto.go
Inserisci un numero: 28
28 è perfetto
```

 
 
## 11 Numeri amichevoli

**Definizione**: Due numeri naturali `x` e `y`, con `x < y`, sono detti amichevoli se la somma dei divisori propri di ciascuno è uguale
all’altro; ad esempio `(220, 284)` è una coppia di amichevoli, essendo `284 = 1 + 2 + 4 + 5 + 10 + 11 + 20 + 22 + 44 + 55 + 110` (dove `1, 2, 4, 5, 10, 11, 20, 22, 44, 55, 110` sono i divisori di `220`) e `220 = 1 + 2 + 4 + 71 + 142` (dove `1, 2, 4, 71, 142` sono i divisori di `284`).

Scrivere un programma che legga da **standard input** un numero intero `soglia` e stampi tutte le coppie di numeri amichevoli tali che `y` sia inferiore a `soglia`.

Il programma deve essere organizzato nelle seguenti funzioni:
* una funzione `LeggiNumero() int` che legge da **standard input** un numero intero e ne restituisce il valore;
* una funzione `SommaDivisoriPropri(n int) int` che riceve in input un valore `int` nel parametro `n` e restituisce un valore `int` pari alla somma dei divisori propri di `n` (`0` se `n` non ha divisori propri);
* una funzione `SonoAmichevoli(n, m int) bool` che riceve in input due valori `int` nei parametri `n` ed `m` e restituisce `true` se `n` ed `m` sono amichevoli e `false` altrimenti (utilizzando la funzione `SommaDivisoriPropri()`);
* una funzione `NumeriAmichevoli(limite int)` che riceve in input un valore `int` nel parametro `limite` e stampa tutte le coppie di numeri amichevoli tali che `y` sia inferiore a `limite` (cfr. Definizione); la funzione deve utilizzare la funzione `SonoAmichevoli()`;
* una funzione `main()` che richiama la funzione `LeggiNumero()` per leggere il numero intero `soglia`. Se `soglia > 0` allora richiama la funzione `NumeriAmichevoli()` in modo opportuno, altrimenti stampa un messaggio d'errore.

##### Esempio d'esecuzione:

```text
$ go run numeri_amichevoli.go
Inserisci un numero: 300
Numeri amichevoli a 300
(220,284) 

$ go run numeri_amichevoli.go
Inserisci un numero: 0
La soglia inserita non è positiva
```

## 12 Garibaldi

Scrivere un programma che: 
* legga da **standard input** un testo su più righe;
* termini la lettura quando viene inserita da **standard input** una riga vuota (`""`);
* ristampi il testo letto (riga vuota esclusa) sostituendo tutte le vocali non accentate (minuscole o maiuscole) con delle `u`.

Il programma deve essere organizzato nelle seguenti funzioni:
* una funzione `LeggiTesto() string` che legge da **standard input** un testo su più righe e terminato da una riga vuota (`""`), restituendo un valore (di tipo) `string` (una stringa) in cui è memorizzato il testo letto (riga vuota esclusa);
* una funzione `TrasformaCarattere(c rune) rune` che riceve in input un valore (di tipo) `rune` nel parametro `c` e restituisce un valore (di tipo) `rune` uguale a `u` se `c` è una vocale non accentata (minuscola o maiuscola) (i.e., se il valore di `c` rappresenta una vocale non accentata (minuscola o maiuscola)) e uguale a `c` altrimenti.
* una funzione `Garibaldi(s string) string` che riceve in input un valore (di tipo) `string` (una stringa) nel parametro `s` e restituisce un valore (di tipo) `string` (una stringa) in cui l'i-esimo carattere è uguale all'i-esimo carattere di `s` se diverso da una vocale non accentata e uguale a `u` altrimenti; la funzione deve utilizzare la funzione `TrasformaCarattere()`;
* una funzione `main()` che richiama la funzione `LeggiTesto()` per leggere il testo da **standard input**, trasforma il testo letto (richiamando opportunamente la funzione `Garibaldi()`) e stampa il testo trasformato.

##### Esempio d'esecuzione:

```text
$ go run garibaldi.go
Inserisci un testo su più righe (termina con riga vuota):
Garibaldi fu ferito
fu ferito in una gamba,
Garibaldi che comanda
che comanda i bersaglier

Risultato trasformazione:
Gurubuldu fu furutu
fu furutu un unu gumbu,
Gurubuldu chu cumundu
chu cumundu u bursugluur
```

## 13 Spaziatura

Scrivere un programma che: 
* legga da **standard input** un testo su più righe (alcune delle quali possono essere delle righe vuote (`""`));
* termini la lettura quando, premendo la combinazione di tasti `Ctrl+D`, viene inserito da **standard input** l'indicatore End-Of-File (EOF);
* ristampi il testo letto con spaziatura, ovvero inserendo uno spazio `' '` tra ogni coppia di caratteri che **non** sono spazi.

Il programma deve essere organizzato nelle seguenti funzioni:
* una funzione `LeggiTesto() string` che legge da **standard input** un testo su più righe (alcune delle quali possono essere delle righe vuote (`""`)) e terminato dall'indicatore EOF, restituendo un valore `string` in cui è memorizzato il testo letto;
* una funzione `Spazia(s string) string` che riceve in input un valore `string` nel parametro `s` e restituisce un valore `string` che rappresenta la versione spaziata di `s`;
* una funzione `main()` che richiama la funzione `LeggiTesto()` per leggere il testo da **standard input**, spazia il testo letto (richiamando opportunamente la funzione `Spazia()`) e stampa il testo spaziato.

*Suggerimento:* per sapere se un carattere è uno spazio si faccia riferimento alla documentazione della funzione `unicode.IsSpace` del package `unicode`.

##### Esempio d'esecuzione:

```text
$ go run spaziatura.go 
Inserisci un testo su più righe (termina con Ctrl+D):
Testo di prova 
da stampare con spaziatura
Risultato:
T e s t o d i p r o v a
d a s t a m p a r e c o n s p a z i a t u r a
```

## 14 Il cifrario di Cesare

Giulio Cesare usava per le sue corrispondenze riservate un codice di sostituzione molto semplice, nel quale la lettera chiara viene sostituita dalla lettera che la segue di tre posti nell’alfabeto: la lettera A è sostituita dalla D, la B dalla E, e così via fino alle ultime lettere che sono cifrate con le prime.

Chiaro:   A B C D E F G H I J K L M N O P Q R S T U V W X Y Z

Cifrato:  D E F G H I J K L M N O P Q R S T U V W X Y Z A B C

Più in generale si dice cifrario di Cesare un codice nel quale ogni lettera del messaggio chiaro viene spostata di un numero fisso `k` di posti (non necessariamente tre), dove `k` è detto **chiave di cifratura**. 

Scrivere un programma che:
* legga da **standard input** un numero intero `k` (la chiave di cifratura);
* legga da **standard input** un messaggio in chiaro su più righe (alcune delle quali possono essere delle righe vuote (`""`)), terminando la lettura quando, premendo la combinazione di tasti `Ctrl+D`, viene inserito da **standard input** l'indicatore End-Of-File (EOF);
* stampi il messaggio cifrato (ottenuto con chiave di cifratura `k`) corrispondente al messaggio in chiaro letto.

Il programma deve essere organizzato nelle seguenti funzioni:
* una funzione `LeggiNumero() int` che legge da **standard input** un numero intero e ne restituisce il valore;
* una funzione `LeggiTesto() string` che legge da **standard input** un testo su più righe (alcune delle quali possono essere delle righe vuote (`""`)) e terminato dall'indicatore EOF, restituendo un valore `string` in cui è memorizzato il testo letto;
* una funzione `CifraCarattere(c rune, chiave int) rune` che riceve in input un valore `rune` nel parametro `c` ed un valore `int` nel parametro `chiave`, e restituisce un valore `rune` uguale a `c` nel caso in cui `c` non sia una lettera dell'alfabeto inglese, uguale al valore cifrato corrispondente a `c` (ottenuto con chiave di cifratura `chiave`) altrimenti (in particolare, il valore cifrato deve essere minuscolo se `c` è minuscolo e maiuscolo se `c` è maiuscolo);
* una funzione `CifraTesto(s string, chiave int) string` che riceve in input un valore `string` nel parametro `s` ed un valore `int` nel parametro `chiave`, e restituisce un valore `string` uguale al valore cifrato corrispondente a `s` (ottenuto con chiave di cifratura `chiave`); la funzione deve utilizzare la funzione `CifraCarattere()`;
* una funzione `main()` che richiama le funzioni `LeggiNumero()` e `LeggiTesto()` per leggere rispettivamente la chiave di cifratura ed il messaggio in chiaro da **standard input**, cifra il messaggio letto (richiamando opportunamente la funzione `CifraTesto()`) e stampa il messaggio cifrato.

##### Esempio d'esecuzione:
 
 ```text
$ go run cifrario_cesare.go 
Inserisci un numero: 1
Inserisci un testo su più righe (termina con CTRL D):
Testo di esempio
diviso su righe diverse
Testo cifrato:
Uftup ej ftfnqjp
ejwjtp tv qsjhif ejwfstf

$ go run cifrario_cesare.go
Inserisci un numero: -2
Inserisci un testo su più righe (termina con CTRL D):
AbC

dEf
Testo cifrato:
YzA

bCd
```
