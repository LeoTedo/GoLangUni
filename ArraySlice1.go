package main

import (
	"fmt"
	"os"
	"strconv"
)

func MinDispari(sl []int) (min int) {

	var i int

	for i = 0; i < len(sl); i++ {
		if sl[i]%2 == 1 {
			min = sl[i]
		}
	}
	i++
	for i = 0; i < len(sl); i++ {
		if (sl[i]%2 == 1) && (sl[i] < min) {
			min = sl[i]
		}
	}
	return
}

func main() {
	var sl []int
	for _, v := range os.Args[1:] {
		val, err := strconv.Atoi(v)
		if err == nil {
			sl = append(sl, val)
		}
	}

	var minimoDispari int
	minimoDispari = MinDispari(sl)

	for _, v := range sl {
		if (v%2 == 0) && (v > minimoDispari) {
			fmt.Print(v, " ")
		}
		fmt.Println(" ")
	}

}
