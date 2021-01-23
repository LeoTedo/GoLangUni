package main

import (
	"fmt"
)

func main() {

	var num string
	var numS []string
	fmt.Scan(&num)

	for i := 0; i < len(string(num))-1; i++ {
		//numS = append(numS, string(num[i]))
		if num[i] == num[i+1] {
			numS = append(numS, string(num[i]))
			i++
		} else {
			//nc, _ := strconv.Atoi(num[i])
			numS = append(numS, string(num[i]))
		}
	}
	fmt.Println(numS)
}
