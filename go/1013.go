package main

import (
	"fmt"
)

func main() {
	var a, b, c, bigger int

	fmt.Scanf("%d %d %d", &a, &b, &c)

	bigger = a

	if bigger < b {
		bigger = b
	}

	if bigger < c {
		bigger = c
	}

	fmt.Printf("%d eh o maior\n", bigger)
}
