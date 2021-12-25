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
	operators := []Operator{
		Addition{Operands{1, 2}},
		Subtraction{Operands{1, 2}},
		Multiplication{Operands{1, 2}},
		Division{Operands{1, 2}},
	}

	for _, operator := range operators {
		fmt.Println(operator.handle())
	}
}

