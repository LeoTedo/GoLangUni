package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n, k, x int
	g := []int{1, 3, 4, 6, 4, 32, 4, 5, 40}
	fmt.Scan(&n, &k)
	fmt.Println(Potenza(n, k))
	fmt.Println(Max(g))
	fmt.Scan(&x)
	fmt.Println(Taglio(x))
	for i, v := range Occo(x) {
		fmt.Println("numero", i, "è apparso", v, "volte")
	}
	//fmt.Println(Occo(x))
	//Prova2()
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

	for _, v := range a {
		v *= 2
	}
	fmt.Println(a)
}

func Taglio(n int) int {
	var convtagliato, convtagliato2 string
	conv := strconv.Itoa(n)
	convtagliato = conv[1 : len(conv)-1]
	convtagliato2 = conv[:1] + conv[len(conv)-1:len(conv)]
	tagliato, _ := strconv.Atoi(convtagliato)
	tagliato2, _ := strconv.Atoi(convtagliato2)
	fmt.Println(tagliato2)
	return tagliato
}

func Occo(n int) map[string]int {
	oc := make(map[string]int)
	conv := strconv.Itoa(n)

	for _, v := range conv {
		oc[string(v)]++
	}

	return oc
}
