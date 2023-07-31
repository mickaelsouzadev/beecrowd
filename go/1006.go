package main

import (
	"fmt"
)

func main() {

	var n1, n2, n3 float64

	p1, p2, p3 := 2.0, 3.0, 5.0

	fmt.Scanf("%f", &n1)
	fmt.Scanf("%f", &n2)
	fmt.Scanf("%f", &n3)

	av := ((p1 * n1) + (p2 * n2) + (p3 * n3)) / (p1 + p2 + p3)

	fmt.Printf("MEDIA = %.1f\n", av)
}
