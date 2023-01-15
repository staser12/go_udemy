package main

import "fmt"

func main() {
	// var c chan int
	c := make(chan int)

	go func() {
		c <- 1
	} ()

	val := <- c
	fmt.Println(val)

	go func() {
		c <- 2
	} ()

	val = <- c
	fmt.Println(val)

	fmt.Println(c)
}