package main

import "fmt"

type Vehicle interface {
	Manufacture() string
}
type Car struct{}

func (car *Car) Manufacture() string {
	return "Manufactured a car"
}

type Motorcycle struct{}

func (moto *Motorcycle) Manufacture() string {
	return "Manufactured a motorcycle"
}

type Bicycle struct{}

func (bicycle *Bicycle) Manufacture() string {
	return "Manufactured a bicycle"
}

type VehicleManufacture interface {
	Create() Vehicle
}

type CarManufacture struct{}

func (f *CarManufacture) Create() Vehicle {
	return &Car{}
}

type MotorcycleManufacture struct{}

func (f *MotorcycleManufacture) Create() Vehicle {
	return &Motorcycle{}
}

type BicycleManufacture struct{}

func (f *BicycleManufacture) Create() Vehicle {
	return &Bicycle{}
}

func main() {
	carFactory := &CarManufacture{}
	motoFactory := &MotorcycleManufacture{}
	bicycleFactory := &BicycleManufacture{}

	car := carFactory.Create()
	motorcycle := motoFactory.Create()
	bicycle := bicycleFactory.Create()

	fmt.Println(car.Manufacture())
	fmt.Println(motorcycle.Manufacture())
	fmt.Println(bicycle.Manufacture())
}
