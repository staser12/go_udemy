package main

import "fmt"

type Car interface {
	Drive()
	Stop()
}

type Lambo struct {
	LamboModel string
}

func NewModel(arg string) Car {
	return &Lambo{arg}
}

type Chevy struct {
	ChevyModel string
}

func (l *Lambo) Drive() {
	fmt.Println("Lambo on the move")
	fmt.Println(l.LamboModel)
}

func (l *Chevy) Drive() {
	fmt.Println("Chevy on the move")
	fmt.Print(l.ChevyModel)
}

func (l *Lambo) Stop() {
	fmt.Println("Stopping Lambo")
}

func main() {
	l := Lambo{"Gallardo"}
	c := Chevy{"369"}
	l.Drive()
	c.Drive()
}