package main

import (
	"fmt"
	"time"
)

type Clock struct {
	hour int
	min  int
	sec  int
}

var h, m, s int

func main() {
	fmt.Println("inserisci tempo con ore minuti e secondi divisi da uno spazio:")
	fmt.Scan(&h, &m, &s)
	tempo := new(Clock)
	tempo.hour = h
	tempo.min = m
	tempo.sec = s

	Countdown(*tempo)

}

func Countdown(c Clock) {
	for {
		if c.sec > 0 {
			c.sec--
		} else if c.min > 0 {
			c.sec = 59
			c.min--
		} else if c.hour > 0 {
			c.min = 59
			ChangeMin(c.hour)
		} else {
			fmt.Println("DRIIIINNNN")
			break
		}

		time.Sleep(time.Duration(1) * time.Second)
		fmt.Println(c)
	}
	return
}

func ChangeMin(m int) int {
	m--
	return m
}

func ChangeHour(h int) int {
	h--
	return h
}

//ciao bello di mamma
