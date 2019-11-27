package main

import (
	"fmt"
	"strconv"
)

func main() {

	slInt := []int64{2, 4, 8, 16, 32}
	sl1 := converti(slInt)
	sl2 := convertiConAppend(slInt)

	fmt.Println(sl1)
	fmt.Println(sl2)
}

func converti(slIn []int64) []string {

	slOut := make([]string,len(slIn))

	for i:=0; i<len(slIn); i++{
		slOut[i] = strconv.FormatInt(slIn[i], 2)
	}

	return slOut
}

func convertiConAppend(slIn []int64) []string {

	var slOut []string

	for i:=0; i<len(slIn); i++{
		slOut = append(slOut, strconv.FormatInt(slIn[i], 2))
	}

	return slOut
}

