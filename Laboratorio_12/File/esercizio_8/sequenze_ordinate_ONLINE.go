package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func main() {

	startTime := time.Now()

	fileList, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while opening file %s: %v\n", os.Args[1], err)
		return
	}
	defer fileList.Close()

	err, inFile, inErr := OpenFiles(fileList)
	for i := 0; i < len(inFile); i++ {
		defer inFile[i].Close()
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Errore while reading file %s: %v\n", os.Args[1], err)
		return
	}
	if Errors(inFile,inErr) {
		LogErrors(inFile, inErr)
		return
	}

	fileOutput, errOutput := os.Create("sequenza_risultante.txt")
	if errOutput != nil {
		fmt.Fprintf(os.Stderr, "Errore while creating file %s: %v\n", "sequenza_risultante.txt", errOutput)
		return
	}
	defer fileOutput.Close()

	errOutput = OutputSequencesMix(inFile,inErr, fileOutput)
	if errOutput != nil {
		fmt.Fprintf(os.Stderr, "Errore while writing file %s: %v\n", "sequenza_risultante.txt", errOutput)
		return
	}
	if Errors(inFile,inErr) {
		LogErrors(inFile, inErr)
		return
	}

	endTime := time.Now()

	fmt.Printf("Execution time: %v\n",endTime.Sub(startTime).Seconds())

}

func OpenFiles(file *os.File) (err error, inFile []*os.File, inErr []error) {
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	inFilename := strings.Split(scanner.Text(), ";")

	if err = scanner.Err(); err != nil {
		return
	}

	for _, filename := range inFilename {
		f, e := os.Open(filename)
		if e != nil {
			return

		}
		inFile = append(inFile, f)
		inErr = append(inErr, e)
	}

	return nil, inFile, inErr

}

func Errors(inFile []*os.File, inErr []error) bool {
	for i := 0; i < len(inFile); i++ {
		if inErr[i] != nil && inErr[i] != io.EOF {
			return true
		}
	}
	return false
}

func LogErrors(inFile []*os.File, inErr []error) {
	for i := 0; i < len(inFile); i++ {
		if inErr[i] != nil && inErr[i] != io.EOF {
			fmt.Fprintf(os.Stderr, "File %s: %s\n", inFile[i], inErr[i])
		}
	}
}


func ReadNextNumbers(inFile []*os.File, inErr []error, nextNumberToRead[]bool, number []int) {
	for i:=0; i<len(inFile); i++{
		if inErr[i] != io.EOF && nextNumberToRead[i] {
			_, inErr[i] = fmt.Fscan(inFile[i], &number[i])
			nextNumberToRead[i] = false
		}
	}
}

func NumberToOutput(inErr []error) bool{
	for i:=0; i<len(inErr); i++ {
		if inErr[i] != io.EOF{
			return true
		}
	}
	return false
}

func InFileIndexOfNumberToOutput(inErr []error, number []int) int {

	var indexOfMin int

	for i:=0; i<len(inErr); i++ {
		if inErr[i] == nil {
			indexOfMin = i
			break;
		}
	}
	for i:=indexOfMin+1; i<len(inErr); i++ {
		if inErr[i] == nil && number[i]<number[indexOfMin]{
			indexOfMin = i
		}
	}

	return indexOfMin
}

func OutputSequencesMix(inFile []*os.File, inErr []error, fileOutput *os.File) (errOutput error) {

	n := len(inFile)

	number := make([]int, n)
	nextNumberToRead := make([]bool, n)
	for i:=0; i<n; i++{
		nextNumberToRead[i] = true
	}
	again := NumberToOutput(inErr)
	for again {
		ReadNextNumbers(inFile, inErr, nextNumberToRead, number)
		if Errors(inFile,inErr) {
			LogErrors(inFile, inErr)
			return
		}

		if again = NumberToOutput(inErr); again {
			index := InFileIndexOfNumberToOutput(inErr, number)
			nextNumberToRead[index] = true
			_, errOutput = fmt.Fprintf(fileOutput, "%d ", number[index])
			if errOutput != nil {
				return
			}
		}
	}
	return nil
}

