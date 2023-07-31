package main

import (
	"fmt"
)

func main() {

	var n, hour, minutes, seconds int

	fmt.Scanf("%d", &n)

	hour = n / 3600
	minutes = n / 60
	seconds = n % 60

	if hour > 0 {
		minutes = minutes % 60
	}

	fmt.Printf("%d:%d:%d\n", hour, minutes, seconds)
}
