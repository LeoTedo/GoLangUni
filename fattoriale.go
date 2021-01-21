package main

import (
	"fmt"
	"os"
	"strconv"
)

func fattoriale(n int) int {
	if n > 0 {
		n = n * fattoriale(n-1)
	} else {
		n += 1
	}
	return n
}

func main() {
	n := os.Args[1]
	nc, _ := strconv.Atoi(n)
	fmt.Println("er fattoriale:", fattoriale(nc))
}
