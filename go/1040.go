package main

import (
	"fmt"
)

func main() {

	var n1, n2, n3, n4 float64

	fmt.Scanf("%f %f %f %f", &n1, &n2, &n3, &n4)

	p1, p2, p3, p4 := 2.0, 3.0, 4.0, 1.0

	av := ((p1 * n1) + (p2 * n2) + (p3 * n3) + (p4 * n4)) / (p1 + p2 + p3 + p4)

	fmt.Printf("Media: %.1f\n", av)

	if av < 5.0 {
		fmt.Println("Aluno reprovado.")
	}

	if av > 7.0 {
		fmt.Println("Aluno aprovado.")
	}

	if av >= 5.0 && av < 7.0 {
		Exam(av)
	}
}

func Exam(av float64) {
	fmt.Println("Aluno em exame.")

	var exam float64

	fmt.Scanf("%f", &exam)

	fmt.Printf("Nota do exame: %.1f\n", exam)
	finalAv := (av + exam) / 2

	if finalAv < 5.0 {
		fmt.Println("Aluno reprovado.")
	}

	if finalAv >= 5.0 {
		fmt.Println("Aluno aprovado.")
	}

	fmt.Printf("Media final: %.1f\n", finalAv)
}
