package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	numero := os.Args[1]
	togli := os.Args[2]
	max := 0

	//found := 0
	//var err error
	numstr := strings.Split(string(numero), "")
	var numstr2 string
	var numstrX string
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

	for f := 0; f < len(numstr); f++ {
		numstrX += numstr[f]
	}

	numstrR := []rune(string(numstrX))
	for i, v := range numstrR {
		if unicode.IsSpace(v) == false {
			numstr2 = numstr2 + string(numstrR[i])
		} else {
			//fmt.Println("bruh")
		}
	}
	fmt.Println("numero migliore scritto carino: ", numstr2)
}

/*for i := 0; i < len(numstr); i++ {
	if numstr[i] == " " {
		numstr2 = numstr2 + numstr[i]
	} else {
		fmt.Println("bruh")
	}
}*/
