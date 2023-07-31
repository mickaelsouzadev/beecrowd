package main

import (
	"fmt"
)

func main() {

	var n int

	evenCount := 0

	for i := 0; i < 5; i++ {
		fmt.Scanf("%d", &n)

		if n%2 == 0 {
			evenCount++
		}
	}

	fmt.Println(evenCount, "valores pares")
}
