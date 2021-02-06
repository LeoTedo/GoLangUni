package main

import "fmt"

func main() {
	var n, k int
	g := []int{1, 3, 4, 6, 4, 32, 4, 5, 40}
	fmt.Scan(&n, &k)
	fmt.Println(Potenza(n, k))
	fmt.Println(Max(g))
}

func Potenza(n, k int) int {
	if k == 0 {
		return 1
	} else {
		for i := 1; i < k; i++ {
			n *= n
		}
		return n
	}
}

func Max(g []int) int {
	var massimo int
	minimo := 1000000
	for i := 0; i < len(g); i++ {
		if g[i] >= massimo {
			massimo = g[i]
		}
		if g[i] <= minimo {
			minimo = g[i]
		}
	}
	fmt.Println(minimo)
	return massimo
}
