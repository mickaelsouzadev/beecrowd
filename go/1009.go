package main

import (
	"fmt"
)

func main() {
	var name string
	var salary, sales float64

	fmt.Scanf("%s", &name)
	fmt.Scanf("%f", &salary)
	fmt.Scanf("%f", &sales)

	comission := sales * 0.15
	total := salary + comission

	fmt.Printf("TOTAL = R$ %.2f\n", total)
}
