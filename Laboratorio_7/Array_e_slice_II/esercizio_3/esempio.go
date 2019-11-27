package main

import (
	"os"
	"fmt"
)

func main() {
	fmt.Printf("tipo os.Args: %T\n", os.Args)
	for i, v := range os.Args {
		fmt.Printf("os.Args[%d]: tipo %T valore %s\n", i, v, v)
	}
}
