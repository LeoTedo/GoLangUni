package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	numero := os.Args[1]
	togli := os.Args[2]
	max := 0
	//found := 0
	//var err error
	numstr := strings.Split(string(numero), "")
	r, _ := strconv.Atoi(togli)
	for i := 0; i < r; i++ {
		fmt.Println("--", r)

		for k := 0; k < len(string(numero))-1; k++ {
			s1, _ := strconv.Atoi(numstr[k])
			s2, _ := strconv.Atoi(numstr[k+1])
			if s1 >= s2 && s1 >= max {
				max = s1
			}
			if s2 >= s1 && s2 >= max {
				max = s2
			}
			//maxconv := string(max)
			fmt.Println(max)
		}
		for f := 0; f < len(string(numero)); f++ {
			finale, _ := strconv.Atoi(numstr[f])
			if finale == max {
				//fmt.Println("check")
				numstr[f] = " "
			}
		}
		max = 0
	}
	fmt.Println("numero migliore scritto da cani: ", numstr)
}
