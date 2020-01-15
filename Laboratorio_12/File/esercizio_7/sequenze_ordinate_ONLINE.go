package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {


	startTime := time.Now()

	file1, err1 := os.Open(os.Args[1])
	if err1 != nil {
		fmt.Println("Errore in apertura del file:", err1)
		return
	}
	defer file1.Close()

	file2, err2 := os.Open(os.Args[2])
	if err2 != nil {
		fmt.Println("Errore in apertura del file:", err2)
		return
	}
	defer file2.Close()

	fileOutput, errOutput := os.Create("sequenza_risultante.txt")
	if errOutput != nil {
		fmt.Println("Errore in creazione del file:", errOutput)
		return
	}
	defer fileOutput.Close()

	err1, err2, errOutput = OutputSequencesMix(file1, file2, fileOutput)
	if err1 != nil || err1 != nil || errOutput != nil {
		fmt.Fprintf(os.Stderr, "File %s: %s \nFile %s: %s\nFile %s: %s\n", file1, err1, file2, err2, fileOutput, errOutput)
	}

	endTime := time.Now()

	fmt.Printf("Execution time: %v\n",endTime.Sub(startTime).Seconds())

}

func OutputSequencesMix(file1, file2 *os.File, fileOutput *os.File) (err1, err2, errOutput error) {
	var numero1, numero2 int

	again := true
	read1 := true
	read2 := true
	for again {
		if err1 != io.EOF && read1 {
			_, err1 = fmt.Fscan(file1, &numero1)
			read1 = false
		}
		if err2 != io.EOF && read2 {
			_, err2 = fmt.Fscan(file2, &numero2)
			read2 = false
		}

		if (err1 != nil && err1 != io.EOF) ||  (err2 != nil && err2 != io.EOF) {
			//fmt.Fprintf(os.Stderr, "File %s: %s \n File %s: %s\n", fileInput1, err1, fileInput2, err2)
			return
		}

		var numeroDaStampare int
		switch {
		case err1 == nil && err2 == nil:
			if numero1 < numero2 {
				numeroDaStampare = numero1
				read1 = true
			} else {
				numeroDaStampare = numero2
				read2 = true
			}
		case err1 == nil && err2 == io.EOF:
			numeroDaStampare = numero1
			read1 = true
		case err1 == io.EOF && err2 == nil:
			numeroDaStampare = numero2
			read2 = true
		default:
			again = false
		}

		if again {
			_, errOutput = fmt.Fprintf(fileOutput, "%d ", numeroDaStampare)
			if errOutput != nil {
				return
			}
		}
	}
	return nil, nil, nil
}

