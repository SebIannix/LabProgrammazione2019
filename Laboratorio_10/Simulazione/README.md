# Laboratorio 10 - Simulazione
## Esercizio Filtro - Codifica compressa

Scrivere un programma che legga da **riga di comando** la codifica compressa di una stringa di caratteri e stampi a video la stringa non compressa.  

La codifica compressa di una stringa di caratteri è definita da una sequenza di coppie di valori `ch n`, dove `ch` rappresenta un carattere, mentre `n>0` è il numero di volte che il carattere `ch` viene ripetuto.

*Esempio:*
Alla sequenza di valori `e 3 a 4`, corrisponde la stringa `eeeaaaa`, ovvero il carattere `'e'`, ripetuto 3 volte, seguito dal carattere `'a'`, ripetuto 4 volte. 

Si assuma che ogni carattere appartenenga all'alfabeto inglese e sia quindi codificato all'interno dello standard US-ASCII (integrato nello standard Unicode).

Si assuma inoltre che la sequenza di valori specificata a riga di comando sia nel formato corretto e includa almeno una coppia `ch n`.


##### Esempi d'esecuzione:

```text
$ go run esercizio_filtro.go e 3 r 4 Y 3
eeerrrrYYY

$ go run esercizio_filtro.go R 3 A 1 B 10
RRRABBBBBBBBBB

$ go run esercizio_filtro.go W 2
WW

$ go run esercizio_filtro.go r 3 P 4
rrrPPPP
```

##### Test automatico:

L'esercizio filtro è considerato esatto **solo se** eseguendo il comando `go test` si ottiene il seguente output:
```text
$ go test
PASS
ok
```

Invece, nel caso in cui l'output dovesse essere simile al seguente
```text
$ go test
--- FAIL: TestFiltro (0.00s)
        esercizio_filtro_test.go:40: 
                Codifica compressa stringa: e 3 r 4 Y 3
                Expected result: <eeerrrrYYY>
                Given result: <eeeerrrrrYYYY>
```
significa che almeno un caso tra quelli riportati nell'esempio d'esecuzione non è stato eseguito in modo corretto, ed il filtro è considerato **errato**.

## Esercizio 1 - Numeri pari superiori al minimo dispari

Scrivere un programma che legga da **riga di comando** una sequenza di valori positivi. Sia `minDispari` il minimo valore dispari tra i valori interi letti.
Il programma deve stampare a video i valori interi pari, e superiori a `minDispari`, presenti nella sequenza di valori letta.

Si assuma che tra i valori interi letti da riga di comando sia sempre presente almeno un valore dispari.

Oltre alla funzione `main()`, devono essere definite ed utilizzate almeno le seguenti funzioni:
* una funzione `MinimoDispari(sl []int) int` che riceve in input un valore `[]int` nel parametro `sl` e restituisce il minimo valore dispari presente in `sl`.

##### Esempio d'esecuzione:

```text
$ go run esercizio_1.go 3 2 4      
4 

$ go run esercizio_1.go 2 5 4 6 7 8
6 8 

$ go run esercizio_1.go 2 6 8 1    
2 6 8 

$ go run esercizio_1.go 1 3 5  

```
## Esercizio 2 - Scala di numeri

Scrivere un programma che legga da **standard input** due numeri interi `n` e `p` e, come mostrato nell'**Esempio d'esecuzione**, stampi a video una scala utilizzando in sequenza i caratteri corrispondenti alle cifre decimali:

* La scala è formata da `n` gradini.
* Ciascun gradino è profondo `p` caratteri e alto 2.
* Il gradino più in alto deve essere stampato senza alcuna rientranza verso destra; considerando poi i successivi gradini dall'alto verso il basso, ciascuno di essi rientra (è traslato) verso destra rispetto al precedente di `p-1` caratteri `' '` (spazio).
* La sequenza di caratteri utilizzata per rappresentare la scala parte dal carattere `0`. Il carattere successivo al carattere `0` è il carattere `1`. Il carattere successivo al carattere `1` è il carattere `2`, e così via. Il carattere successivo al carattere `9` è il carattere `0`.

Si assuma che vengano sempre inseriti da **standard input** valori interi positivi, sia per `n` sia per `p`.

##### Esempio d'esecuzione

```text
$ go run esercizio_2.go 
1 4
0123
   4

$ go run esercizio_2.go
4 3
012
  3
  456
    7
    890
      1
      234
        5

$ go run esercizio_2.go
3 4
0123
   4
   5678
      9
      0123
         4
```
## Esercizio 3 - Conta cifre e lettere

Scrivere un programma che: 
* legga da **standard input** un testo su più righe (alcune delle quali possono essere delle righe vuote (`""`));
* termini la lettura quando, premendo la combinazione di tasti `Ctrl+D`, viene inserito da **standard input** l'indicatore End-Of-File (EOF);
* stampi a video le seguenti informazioni relative al testo letto:
1. La lista di lettere distinte presenti nel testo, riportando per ogni lettera il relativo numero di occorrenze nel testo (cfr. **Esecuzione d'esecuzione**). Una lettera è un carattere il cui codice Unicode, se passato come argomento alla funzione `func IsLetter(r rune) bool` del package `unicode`, fa restituire `true` alla funzione. Le lettere minuscole sono da considerarsi diverse dalle lettere maiuscole.
2. Il numero totale di lettere presenti nel testo.
3. La lista delle cifre decimali distinte presenti nel testo, riportando per ogni cifra decimale il relativo numero di occorrenze nel testo (cfr. **Esecuzione d'esecuzione**). Una cifra decimale è un carattere il cui codice Unicode, se passato come argomento alla funzione `func IsDigit(r rune) bool` del package `unicode`, fa restituire `true` alla funzione.
4. Il numero totale di cifre decimali presenti nel testo.

##### Esempio d'esecuzione:

```text
$ go run esercizio_3.go 
Via Celoria 18, 
20133 - Milano
 
REPORT:
Totale lettere: 16
V -> 1
a -> 3
C -> 1
e -> 1
l -> 2
o -> 2
r -> 1
n -> 1
i -> 3
M -> 1
Totale cifre: 7
1 -> 2
8 -> 1
2 -> 1
0 -> 1
3 -> 2

$ go run esercizio_3.go 
Via Celoria 18
REPORT:
Totale lettere: 10
a -> 2
C -> 1
e -> 1
l -> 1
o -> 1
r -> 1
V -> 1
i -> 2
Totale cifre: 2
1 -> 1
8 -> 1

$ go run esercizio_3.go 
44 gatti 
in fila per 3

col resto di 2
REPORT:
Totale lettere: 24
i -> 4
l -> 2
p -> 1
e -> 2
r -> 2
s -> 1
d -> 1
t -> 3
n -> 1
f -> 1
c -> 1
o -> 2
g -> 1
a -> 2
Totale cifre: 4
4 -> 2
3 -> 1
2 -> 1
```

## Esercizio 4 - Tragitto

Sul piano cartesiano, ad ogni punto individuato da una coppia di numeri reali, chiamati rispettivamente *ascissa* e *ordinata*, può essere associata un'*etichetta* simbolica, generalmente una lettera maiuscola.

Scrivere un programma che:
* legga da **standard input** una sequenza di righe di testo;
* termini la lettura quando, premendo la combinazione di tasti `Ctrl+D`, viene inserito da **standard input** l'indicatore End-Of-File (EOF);

Ogni riga di testo è una stringa nel seguente formato:

*etichetta*;*x*;*y*
   
La tripla di valori separati dal carattere `;` specifica un punto sul piano cartesiano:
   
1. *etichetta*: una stringa che specifica l'etichetta simbolica associata al punto (ad es.: "A", "B", ...)
2. *x*: un valore reale che specifica l'ascissa del punto;
3. *y*: un valore reale che specifica l'ordinata del punto.

Si ipotizzi che vengano inserite da **standard input** le seguenti di righe di testo:

```text
A;10.0;2.0
B;11.5;3.0
C;8.0;1.0
```

Ogni coppia di punti specificata da righe consecutive
```text
A;10.0;2.0
B;11.5;3.0
```
```text
B;11.5;3.0
C;8.0;1.0
```
descrive una tratta di un tragitto che va dal primo punto (`A`) all'ultimo punto (`C`). 

La lunghezza di ciascuna tratta del tragitto è pari alla distanza euclidea tra i due punti che definiscono la tratta. 
Per esempio, la lunghezza della prima tratta, quella che va dal punto `A` al punto `B`, è pari alla distanza euclidea tra i punti `A` e `B`: ((x<sub>A</sub>-x<sub>B</sub>)<sup><small>2</small></sup> + (y<sub>A</sub>-y<sub>B</sub>)<sup><small>2</small></sup>)<sup><small>1/2</small></sup>.

Una volta terminata la fase di lettura, il programma deve:
* stampare a video la lunghezza totale del tragitto descritto dalla sequenza di punti specificati dalle righe di testo lette;
* stampare a video la rappresentazione `string` del primo punto che si incontra dopo aver percorso più della metà del tragitto (il primo tra quelli presenti nella sequenza di punti che descrive il tragitto).

Si assuma che:
* le righe di testo lette da **standard input** siano nel formato corretto;
* la tripla di valori presente in ogni riga specifichi correttamente un punto sul piano cartesiano;
* vengano lette da **standard input** almeno 2 righe di testo.

Oltre alla funzione `main()`, devono essere definite ed utilizzate almeno le seguenti funzioni:
* una funzione `NuovoPercorso() (percorso []Punto)` che:
  1. legge da **standard input** una sequenza di righe di testo nel formato *etichetta*;*x*;*y*, terminando la lettura quando viene letto l'indicatore End-Of-File (EOF);
  2. restituisce un valore `[]Punto` nella variabile `percorso` in cui è memorizzata la sequenza di istanze del tipo `Punto` inizializzate in base alle triple di valori specificate nelle righe di testo lette;
* una funzione `Distanza(p1, p2 Punto) float64` che riceve in input due instanze del tipo  `Punto` nei parametri `p1` e `p2` e restituisce un valore `float64` pari alla distanza euclidea tra i punti rappresentati da `p1` e `p2`;
* una funzione `String(p Punto) string` che riceve in input un'instanza del tipo `Punto` nel parametro `p` e restituisce un valore `string` che corrisponde alla rappresentazione `string` di `p` nel formato `ETICHETTA = (X, Y)`, dove `ETICHETTA` è il valore `string` che specifica l'etichetta simbolica di `p`, mentre `X` ed `Y` sono i valori `float64` che specificano rispettivamente l'ascissa e l'ordinata di `p`.

##### Esempio d'esecuzione:

```text
$ go run esercizio_5.go
A;10.0;2.0
B;11.5;3.0
C;8.0;1.0
D;3;4
E;1;0
F;-1;5
Lunghezza percorso: 21.522157168860655
Punto oltre metà: D = (3, 4)

$ go run esercizio_5.go
A;0;0
B;4;0
C;4;4
D;0;4
E;0;0
Lunghezza percorso: 16
Punto oltre metà: D = (0, 4)
```
