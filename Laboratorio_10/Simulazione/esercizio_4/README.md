# Tragitto

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