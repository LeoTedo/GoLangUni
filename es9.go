package main

import "fmt"

func main() {

	var b [3]int = [3]int{2, 20, 200}
	fmt.Println(b)

	var a [3][2]int
	a = [3][2]int{{1, 3}, {10, 30}, {100, 300}}
	fmt.Println(a)

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			fmt.Printf("%3d ", a[i][j])
		}
		fmt.Printf("\n")
	}

	var m, n int = 5, 3
	var s [][]int
	s = make([][]int, m)
	for i := 0; i < m; i++ {
		s[i] = make([]int, n)
	}

	fmt.Println(s)

	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s[i]); j++ {
			fmt.Printf("%3d ", s[i][j])
		}
		fmt.Printf("\n")
	}
}
