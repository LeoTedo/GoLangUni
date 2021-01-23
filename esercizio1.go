package main

import (
	"fmt"
	"strconv"
)

func main() {

	var n1, n2 int
	var ris []string
	fmt.Scan(&n1, &n2)

	for i := n1; i <= n2; i++ {
		//fmt.Println(i)

		if i%3 != 0 && i%5 != 0 && i%7 != 0 && i%15 != 0 && i%21 != 0 && i%35 != 0 && i%105 != 0 {
			ns := strconv.Itoa(i)
			ris = append(ris, ns)
		} else {
			if i%105 == 0 {
				ris = append(ris, "DinDonDan")
				continue
			}
			if i%35 == 0 {
				ris = append(ris, "DonDan")
				continue
			}
			if i%21 == 0 {
				ris = append(ris, "DinDan")
				continue
			}
			if i%15 == 0 {
				ris = append(ris, "DinDon")
				continue
			}
			if i%7 == 0 {
				ris = append(ris, "Dan")
				continue
			}
			if i%5 == 0 {
				ris = append(ris, "Don")
				continue
			}
			if i%3 == 0 {
				ris = append(ris, "Din")
				continue
			}
		}

	}
	fmt.Print(ris)
}
