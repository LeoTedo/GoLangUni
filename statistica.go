package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var numeri []float64

	fmt.Println("inserisci i numeri")
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		input := scanner.Text()
		numconv, _ := strconv.ParseFloat(input, 4)
		numeri = append(numeri, numconv)
	}
	fmt.Println(max(numeri))
	fmt.Println(min(numeri))
	fmt.Println(med(numeri))
}

func max(num []float64) (max float64) {
	for i := 0; i < len(num); i++ {
		if num[i] >= max {
			max = num[i]
		}
	}
	return max
}

func min(num []float64) (min float64) {
	minimo := 1000000000.00
	for i := 0; i < len(num); i++ {
		if num[i] <= minimo {
			minimo = num[i]
		}
	}
	return minimo
}

func med(num []float64) (med float64) {
	totale := 0.0
	for i := 0; i < len(num)-1; i++ {
		totale += num[i]
	}
	med = totale / float64(len(num))
	return med
}
