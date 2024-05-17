/*
Write a program that creates two custom struct types called 'triangle' and 'square'
-> the square type should be a struct with a field called 'sideLength' of type float64
-> the triangle type should be a struct with fields called 'base' and 'height' of type float64
-> both types should have a function called getArea associated with them.
*/
package main

import "fmt"

type Square struct {
	sideLength float64
}

type Triangle struct {
	base   float64
	height float64
}

type Shape interface {
	getArea() float64
}

func main() {
	s := Square{sideLength: 10}
	t := Triangle{base: 10, height: 10}

	printArea(s)
	printArea(t)
}

func (s Square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t Triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func printArea(s Shape) {
	fmt.Println(s.getArea())
}
