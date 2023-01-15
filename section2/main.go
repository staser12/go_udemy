package main

import "fmt"

func swap(m1, m2 *int) {
	var temp int
	temp = *m2
	*m2 = *m1
	*m1 = temp
}

type Car struct {
	Name string
	Age int
	ModelNo int
}

func (c Car) Print() {
	fmt.Print(c)
}

func main() {
	/* m1, m2 := 2, 3
	swap(&m1, &m2)
	fmt.Println(m1)
	fmt.Println(m2) */
	c := Car{"chevy", 1, 2}
	fmt.Print(c)
	c.Print()

}