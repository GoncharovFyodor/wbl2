package main

import "fmt"

// Структура "Автомобиль"
type Car struct {
	name     string
	color    string
	maxSpeed int
}

func (c *Car) GetName() string {
	return c.name
}

func (c *Car) GetColor() string {
	return c.color
}

func (c *Car) GetMaxSpeed() int {
	return c.maxSpeed
}

// Интерфейс Builder длля структуры "Автомобиль"
type CarBuilder interface {
	SetName(name string)
	SetColor(color string)
	SetMaxSpeed(maxSpeed int)
	GetCar() *Car
}

// Реализация интерфейса CarBuilder
type ConcreteCarBuilder struct {
	car *Car
}

func NewConcreteCarBuilder() *ConcreteCarBuilder {
	return &ConcreteCarBuilder{car: &Car{}}
}

func (c *ConcreteCarBuilder) SetName(name string) {
	c.car.name = name
}

func (c *ConcreteCarBuilder) SetColor(color string) {
	c.car.color = color
}

func (c *ConcreteCarBuilder) SetMaxSpeed(maxSpeed int) {
	c.car.maxSpeed = maxSpeed
}

func (c *ConcreteCarBuilder) GetCar() *Car {
	return c.car
}

// Директор, который управляет процессом конструирования
type Director struct {
	builder CarBuilder
}

func NewDirector(builder CarBuilder) *Director {
	return &Director{
		builder: builder,
	}
}

func (d *Director) Construct() {
	d.builder.SetName("BMW")
	d.builder.SetColor("white")
	d.builder.SetMaxSpeed(290)
}

func main() {
	builder := NewConcreteCarBuilder()
	director := NewDirector(builder)
	director.Construct()

	car := builder.GetCar()
	fmt.Printf("Car: %s, color %s, max speed: %d\n", car.GetName(), car.GetColor(), car.GetMaxSpeed())
}
