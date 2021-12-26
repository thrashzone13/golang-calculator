package main

import (
	"fmt"
)

type Operator interface {
	handle() int
}

type Operands struct {
	x int
	y int
}

type Subtraction struct {
	Operands
}

type Addition struct {
	Operands
}

type Multiplication struct {
	Operands
}

type Division struct {
	Operands
}

func (operation Division) handle() int {
	return operation.x / operation.y
}

func (operation Subtraction) handle() int {
	return operation.x - operation.y
}

func (operation Addition) handle() int {
	return operation.x + operation.y
}

func (operation Multiplication) handle() int {
	return operation.x * operation.y
}

func main() {
	var x, y int
	var o string
	var operator Operator

	fmt.Println("Enter first num:")
	fmt.Scan(&x)
	fmt.Println("Enter second num:")
	fmt.Scan(&y)

	operands := Operands{x, y}

	fmt.Println("Enter operator (+, -, *, /):")
	for {
		if operator != nil {
			break
		}
		fmt.Scan(&o)
		switch o {
		case "+":
			operator = Addition{operands}
		case "_":
			operator = Subtraction{operands}
		case "/":
			operator = Division{operands}
		case "*":
			operator = Multiplication{operands}
		default:
			fmt.Println("Enter a valid operator (+, -, *, /)")
		}
	}

	fmt.Println("Result is: ", operator.handle())
}
