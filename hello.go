package main

import (
	"fmt"
	"math"
)

const pi = math.Pi

type Figure interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) Area() float64 {
	return r.height * r.width
}

func (c Circle) Area() float64 {
	return c.radius * c.radius * pi
}

func (r Rectangle) Perimeter() float64 {
	return (r.height + r.width) * 2
}

func (c Circle) Perimeter() float64 {
	return c.radius * 2 * pi
}

func main() {
	rectangle := Rectangle{
		width:  5,
		height: 3,
	}

	circle := Circle{
		radius: 4,
	}

	fmt.Printf("Прямиоугольник:\nПлощадь:%.2f\nПериметр:%.2f\n\n", rectangle.Area(), rectangle.Perimeter())
	fmt.Printf("Круг:\nПлощадь:%.2f\nПериметр:%.2f\n\n", circle.Area(), circle.Perimeter())
}
