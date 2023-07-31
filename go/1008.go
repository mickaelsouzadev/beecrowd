package main

import (
	"fmt"
)

func main() {

	var number, hour int
	var price float64

	fmt.Scanf("%d", &number)
	fmt.Scanf("%d", &hour)
	fmt.Scanf("%f", &price)

	salary := float64(hour) * price

	fmt.Printf("NUMBER = %d\n", number)
	fmt.Printf("SALARY = U$ %.2f\n", salary)
}
