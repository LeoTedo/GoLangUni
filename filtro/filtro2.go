package main

import (
	"fmt"
)

func main() {

	var num string
	var ris string
	fmt.Scan(&num)
	//fmt.Sprint(string(len(num)))
	for i := 0; i < len(num); i++ {

		if i == len(num)-1 {
			ris += (string(num[len(num)-1]))
			break
		} else {
			if num[i] == num[i+1] {
				continue
			} else {
				ris += (string(num[i]))
			}
		}
	}
	fmt.Println("Output:", ris)
}
