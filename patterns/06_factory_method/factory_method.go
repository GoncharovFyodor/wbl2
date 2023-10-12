package main

import "fmt"

type Person interface {
	SayHello()
}

type Creator interface {
	FactoryMethod() Person
}

type ConcretePerson struct{}

func (c *ConcretePerson) SayHello() {
	fmt.Println("Hello!")
}

type ConcreteCreator struct{}

func (c *ConcreteCreator) FactoryMethod() Person {
	return &ConcretePerson{}
}

func main() {
	creator := &ConcreteCreator{}
	person := creator.FactoryMethod()
	person.SayHello()
}
