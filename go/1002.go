package main

import (
	"fmt"
	"math"
)

func main() {
	const Pi float64 = 3.14159
	var ray float64

	fmt.Scanf("%f", &ray)

	a := Pi * math.Pow(ray, 2)

	fmt.Printf("A=%.4f\n", a)
}
