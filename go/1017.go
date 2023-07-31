package main

import (
	"fmt"
)

func main() {
	const KmByLiters float64 = 12
	var liters, time, speed float64

	fmt.Scanf("%f", &time)
	fmt.Scanf("%f", &speed)

	liters = time * speed / KmByLiters

	fmt.Printf("%.3f\n", liters)
}
