package main

import (
	"fmt"
	"math"
)

func main() {

	var a, b, c float64

	fmt.Scan(&a, &b, &c)

	discriminante := b*b - 4*a*c

	x1 := (-b + math.Sqrt(discriminante)) / (2 * a)
	x2 := (-b - math.Sqrt(discriminante)) / (2 * a)
	fmt.Println("Radici dell'equazione:", x1, x2)

}
