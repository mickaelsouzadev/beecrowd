package main

import (
	"fmt"
)

func main() {

	var n1, n2, av float64
	p1 := 3.5
	p2 := 7.5

	fmt.Scanf("%f", &n1)
	fmt.Scanf("%f", &n2)

	av = ((n1 * p1) + (p2 * n2)) / (p1 + p2)

	fmt.Printf("MEDIA = %.5f\n", av)

}
