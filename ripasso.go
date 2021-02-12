package main

import "fmt"

func main() {
	var n, k int
	g := []int{1, 3, 4, 6, 4, 32, 4, 5, 40}
	fmt.Scan(&n, &k)
	fmt.Println(Potenza(n, k))
	fmt.Println(Max(g))
	Prova2()
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

/*func Prova() {
	var n int = 4
	var a [n]int

	for i := 0; i < n; i++ {
		a[i] = i
	}
	fmt.Println(a)
}*/

func Prova2() {

	var a [6]int

	for i, v := range a {
		v = i
	}

	for _, v := range a {
		v *= 2
	}
	fmt.Println(a)
}
