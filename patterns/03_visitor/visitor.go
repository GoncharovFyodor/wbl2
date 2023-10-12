package main

import "fmt"

// Интерфейс фигуры (имеет метод accept, принимающий посетителя)
type Shape interface {
	getType() string
	accept(Visitor)
}

// Квадрат
type Square struct {
	side int
}

func (s *Square) accept(v Visitor) {
	v.visitForSquare(s)
}

func (s *Square) getType() string {
	return "Квадрат"
}

// Круг
type Circle struct {
	radius int
}

func (c *Circle) accept(v Visitor) {
	v.visitForCircle(c)
}

func (c *Circle) getType() string {
	return "Circle"
}

// Прямоугольник
type Rectangle struct {
	l int
	b int
}

func (t *Rectangle) accept(v Visitor) {
	v.visitForRectangle(t)
}

func (t *Rectangle) getType() string {
	return "Прямоугольник"
}

// Интерфейс посетителя
type Visitor interface {
	visitForSquare(*Square)
	visitForCircle(*Circle)
	visitForRectangle(*Rectangle)
}

type AreaCalculator struct {
	area int
}

func (a *AreaCalculator) visitForSquare(s *Square) {
	fmt.Println("Расчет площади квадрата")
}

func (a *AreaCalculator) visitForCircle(s *Circle) {
	fmt.Println("Расчет площади круга")
}

func (a *AreaCalculator) visitForRectangle(s *Rectangle) {
	fmt.Println("Расчет площади прямоугольника")
}

type MiddleCoordinates struct {
	x int
	y int
}

func (a *MiddleCoordinates) visitForSquare(s *Square) {
	fmt.Println("Расчет координат центра квадрата")
}

func (a *MiddleCoordinates) visitForCircle(c *Circle) {
	fmt.Println("Расчет координат центра круга")
}

func (a *MiddleCoordinates) visitForRectangle(t *Rectangle) {
	fmt.Println("Расчет координат центра прямоугольника")
}

func main() {
	square := &Square{side: 2}
	circle := &Circle{radius: 3}
	rectangle := &Rectangle{l: 2, b: 3}
	areaCalculator := &AreaCalculator{}
	square.accept(areaCalculator)
	circle.accept(areaCalculator)
	rectangle.accept(areaCalculator)

	fmt.Println()
	middleCoordinates := &MiddleCoordinates{}
	square.accept(middleCoordinates)
	circle.accept(middleCoordinates)
	rectangle.accept(middleCoordinates)
}
