package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {

	startTime := time.Now()

	file1, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Errore in apertura del file:", err)
		return
	}
	defer file1.Close()

	file2, err := os.Open(os.Args[2])
	if err != nil {
		fmt.Println("Errore in apertura del file:", err)
		return
	}
	defer file2.Close()

	seq1, err := LeggiSequenza(file1)
	if err != nil {
		fmt.Println("Errore in lettura del file:", err)
		return
	}
	seq2, err := LeggiSequenza(file2)
	if err != nil {
		fmt.Println("Errore in lettura del file:", err)
		return
	}

	risultato := CalcolaMixSequenze(seq1, seq2)

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

func CalcolaMixSequenze(s1, s2 []int) (risultato []int) {
	for len(s1) > 0 && len(s2) > 0 {
		if s1[0] < s2[0] {
			risultato = append(risultato, s1[0])
			s1 = s1[1:]
		} else {
			risultato = append(risultato, s2[0])
			s2 = s2[1:]
		}
	}
	risultato = append(risultato, s1...)
	risultato = append(risultato, s2...)
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
