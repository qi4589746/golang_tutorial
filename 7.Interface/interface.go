package main

import "fmt"

type Car interface {
	Drive()
	Stop()
}

// Lambo struct

type Lambo struct {
	LamboModel string
}

// 將Lambo結構賦予Car的Interface
// 與下方Chevy最大差異為，必定要實做Interface中的function
func NewModel(arg string) Car {
	return &Lambo{arg}
}

func (lam *Lambo) Stop() {
	fmt.Println(lam.LamboModel + ": Lambo is Stopping")
}

func (lam *Lambo) Drive() {
	fmt.Println(lam.LamboModel + ": Lambo on the move")
}

// Chevy struct

type Chevy struct {
	ChevyModel string
}

func (che *Chevy) Drive() {
	fmt.Println(che.ChevyModel + ": Chevy on the move")
}

// func (che *Chevy) Stop() {
// 	fmt.Println(che.ChevyModel + ": Chevy is Stopping")
// }

func main() {
	l := NewModel("lambo156")
	c := Chevy{"Chevy873"}
	l.Drive()
	l.Stop()
	c.Drive()
}
