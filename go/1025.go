package main

import (
	"fmt"
	"sort"
)

func main() {

	var N, Q int
	caseQ := 1

	for {
		fmt.Scanf("%d %d", &N, &Q)

		if N == 0 && Q == 0 {
			break
		}

		marbles := make([]int, N)
		founded := make([]int, Q)
		keys := make([]int, Q)

		for i := 0; i < N; i++ {
			var mN int

			fmt.Scanf("%d", &mN)

			marbles[i] = mN
		}

		sort.Ints(marbles)

		for j := 0; j < Q; j++ {
			var mQ int

			fmt.Scanf("%d", &mQ)

			founded[j] = mQ

			pos := sort.SearchInts(marbles, mQ)
			if pos < len(marbles) && marbles[pos] == mQ {
				keys[j] = pos + 1
				continue
			}

			keys[j] = 0
		}

		if Q != 0 {
			fmt.Printf("CASE# %d:\n", caseQ)
		}

		for l := 0; l < len(founded); l++ {
			if keys[l] == 0 {
				fmt.Printf("%d not found\n", founded[l])
			} else {
				fmt.Printf("%d found at %d\n", founded[l], keys[l])
			}
		}

		caseQ++
	}
}
