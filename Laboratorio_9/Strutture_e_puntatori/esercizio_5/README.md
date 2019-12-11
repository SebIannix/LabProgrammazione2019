# Frazioni

Al fine di modellare l'entità matematica `frazione`, si definisca un nuovo tipo `Frazione` come una struttura avente due campi `Numeratore` e `Denominatore` di tipo `int`.

Implementare le funzioni:

* `NuovaFrazione(numeratore, denominatore int) *Frazione` che restituisce una nuova istanza del tipo `Frazione` inizializzata in base ai valori dei parametri `numeratore` e `denominatore`;
* `String(f Frazione) string` che riceve in input un'instanza del tipo `Frazione` nel parametro `f` e restituisce un valore `string` che corrisponde alla rappresentazione  `string` di `f` nel formato `NUMERATORE/DENOMINATORE`, dove `NUMERATORE` (`DENOMINATORE`) è il valore `string` corrispondente al valore del campo `Numeratore` (`Denominatore`) di `f`;

affinché la funzione `main()` di seguito riportata possa essere eseguita generando l'output atteso.

Funzione `main()`:
```go
func main() {
	frazione := NuovaFrazione(34, 18)
	fmt.Println(String(*frazione))
}
```  
Output atteso:
```text
34/18
```