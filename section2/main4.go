package main

import "fmt"

func Anything(anything interface{}) {
	fmt.Println(anything)
}

func main() {
	fmt.Println("vim go")
	Anything(2.44)
	Anything("Name")
}