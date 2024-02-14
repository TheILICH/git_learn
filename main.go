package main

import (
	"fmt"
)

func main() {

	N := 100
	p := make([]bool, N)

	for i := 0; i < N; i++ {
		p[i] = true
	}

	p[0], p[1] = false, false
	for i := 0; i < N; i++ {
		if !p[i] {
			continue
		}
		for j := i * i; j < N; j += i {
			p[j] = false
		}
	}

	var res []int
	for i := 0; i < N; i++ {
		if p[i] {
			res = append(res, i)
		}
	}

	fmt.Printf("PRIMES BETWEEN 1 AND %d ARE => \n%v\n", N, res)


}
