package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	lista := os.Args[0:]
	//var err error
	for i := 0; i < len(lista)-1; i++ {
		num, _ := strconv.Atoi(lista[i+1])
		for k := 0; k < num; k++ {
			fmt.Print(lista[i])
			//fmt.Println()
		}
	}
}