package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
	"time"
)

func main() {
	startTime := time.Now()
	fileLista, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Errore in apertura del file:", err)
		return
	}
	defer fileLista.Close()

	sequenze, err := LeggiSequenze(fileLista)
	if err != nil {
		fmt.Println("Errore in lettura del file:", err)
		return
	}

	risultato := CalcolaMixSequenze(sequenze)

	fileOutput, err := os.Create("sequenza_risultante.txt")
	if err != nil {
		fmt.Println("Errore in creazione del file:", err)
		return
	}
	defer fileOutput.Close()

	err = StampaSequenza(fileOutput, risultato)
	if err != nil {
		fmt.Println("Errore in scrittura del file:", err)
		return
	}
	endTime := time.Now()

	fmt.Printf("Execution time: %v\n",endTime.Sub(startTime).Seconds())
}

func LeggiSequenze(file *os.File) (sequenze [][]int, err error) {
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	listaFileSequenza := strings.Split(scanner.Text(), ";")
	if err = scanner.Err(); err != nil {
		return
	}

	for _, nomeFileSequenza := range listaFileSequenza {
		fileSequenza, err := os.Open(nomeFileSequenza)
		if err != nil {
			return nil, err
		}
		defer fileSequenza.Close()

		sequenza, err := LeggiSequenza(fileSequenza)
		if err != nil {
			return nil, err
		}

		sequenze = append(sequenze, sequenza)
	}

	return sequenze, nil
}

func LeggiSequenza(fileInput *os.File) (sequenza []int, err error) {
	for {
		var n, numero int
		n, err = fmt.Fscan(fileInput, &numero)
		if n > 0 {
			sequenza = append(sequenza, numero)
		}
		if err == io.EOF {
			return sequenza, nil
		} else if err != nil {
			return nil, err
		}
	}
}

func Minimo(sequenze [][]int) (minimo, indiceSequenza int) {
	minimo = math.MaxInt64
	indiceSequenza = -1
	for i, sequenza := range sequenze {
		if len(sequenza) > 0 && minimo > sequenza[0] {
			minimo = sequenza[0]
			indiceSequenza = i
		}
	}
	return
}

func CalcolaMixSequenze(sequenze [][]int) (risultato []int) {
	copiaSequenze := make([][]int, len(sequenze))
	copy(copiaSequenze, sequenze)

	for {
		minimo, indice := Minimo(copiaSequenze)
		if indice < 0 {
			break
		}

		risultato = append(risultato, minimo)
		copiaSequenze[indice] = copiaSequenze[indice][1:]
	}

	return
}

func StampaSequenza(fileOutput *os.File, sequenza []int) (err error) {
	for _, v := range sequenza {
		_, err = fmt.Fprintf(fileOutput, "%d ", v)
		if err != nil {
			return
		}
	}
	return
}
