package oop

import "fmt"

type Plane struct {
	*VehicleBase
	fuel string
}

func NewPlane(serialNumber, fuel string, fillFunc func() string) *Plane {
	return &Plane{VehicleBase: NewVehicleBase(serialNumber, fillFunc), fuel: fuel}
}

// rewriting
func (c *Plane) Move() {
	fmt.Println("I am flying")
}

// parent method usage with additional functionality
func (c *Plane) GetInfo() string {
	return c.VehicleBase.GetInfo() + fmt.Sprintf(" My fuel is %v.", c.fuel)
}

// implemented in successor
func (c *Plane) repairBy() string {
	return "air mechanic."
}

func (c *Plane) Repair() {
	c.doRepair(c)
}
