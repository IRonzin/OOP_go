package oop

import "fmt"

type Car struct {
	*VehicleBase
	fuel string
}

func NewCar(serialNumber, fuel string) *Car {

	return &Car{VehicleBase: NewVehicleBase(serialNumber, func() string {
		return fuel
	}), fuel: fuel}
}

//rewriting
func (c *Car) Move() {
	fmt.Println("I am driving")
}

//parent method usage with additional functionality
func (c *Car) GetInfo() string {
	return c.VehicleBase.GetInfo() + fmt.Sprintf(" My fuel is %v.", c.fuel)
}

//implemented in successor
func (c *Car) repairBy() string {
	return "car mechanic."
}

func (c *Car) Repair() {
	c.doRepair(c)
}
