package main

import "fmt"

type Operation int

const (
	Add = iota
	Subtract
	Multiply
	Divide
)

func (op Operation) calculate(lhs, rhs int) int {
	switch op {
	case Add:
		return lhs + rhs
	case Subtract:
		return lhs - rhs
	case Multiply:
		return lhs * rhs
	case Divide:
		return lhs / rhs
	default:
		panic("Unknown operation")
	}
}

func main() {
	add := Operation(Add)
	fmt.Println(add.calculate(2, 2))

	sub := Operation(Subtract)
	fmt.Println(sub.calculate(10, 3))

	mul := Operation(Multiply)
	fmt.Println(mul.calculate(4, 3))

	div := Operation(Divide)
	fmt.Println(div.calculate(100, 2))

}
