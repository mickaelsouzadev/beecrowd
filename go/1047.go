package main

import (
	"fmt"
)

func main() {

	var h1, m1, h2, m2 int

	fmt.Scanf("%d %d %d %d", &h1, &m1, &h2, &m2)

	h := h2 - h1
	m := m2 - m1

	if h == 0 && m <= 0 {
		h = 24
	}

	if m < 0 {
		h = h - 1
		m += 60
	}

	if h < 0 {
		h += 24
	}

	fmt.Printf("O JOGO DUROU %d HORA(S) E %d MINUTO(S)\n", h, m)
}
