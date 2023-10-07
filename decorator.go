package main

import (
	"fmt"
)

type Coffee interface {
	Cost() float64
}

type Espresso struct{}

func (e *Espresso) Cost() float64 {
	return 1000.00
}

type CoffeeDecorator interface {
	Cost() float64
}

type MilkDecorator struct {
	Coffee Coffee
}

func (m *MilkDecorator) Cost() float64 {
	return m.Coffee.Cost() + 500.00
}

type SyrupDecorator struct {
	Coffee Coffee
}

func (s *SyrupDecorator) Cost() float64 {
	return s.Coffee.Cost() + 300.00
}

func main() {
	coffee := &Espresso{}
	coffeeWithMilk := &MilkDecorator{Coffee: coffee}
	coffeeWithMilkAndSyrup := &SyrupDecorator{Coffee: coffeeWithMilk}

	fmt.Printf("Cost of decorated coffee: %.2f KZT", coffeeWithMilkAndSyrup.Cost())
}
