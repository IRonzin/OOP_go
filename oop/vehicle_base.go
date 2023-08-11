package oop

import "fmt"

type VehicleBase struct {
	serialNumber string
	doFillFuel   func() string
}

func NewVehicleBase(serialNumber string, doFillFuel func() string) *VehicleBase {
	var fillFunc func() string
	if doFillFuel == nil {
		fillFunc = func() string {
			panic("not implemented")
		}
	} else {
		fillFunc = doFillFuel
	}

	return &VehicleBase{serialNumber: serialNumber, doFillFuel: fillFunc}
}

//implemented in successor
func (c *VehicleBase) repairBy() string {
	panic("not implemented")
}

func (b *VehicleBase) doRepair(v IVehicle) {
	fmt.Println("I am repired by " + v.repairBy())
}

func (b *VehicleBase) Move() {
	fmt.Println("I am moving...")
}

func (b *VehicleBase) GetInfo() string {
	return fmt.Sprintf("I am vehicle with serial number: %v.", b.serialNumber)
}

func (b *VehicleBase) Repair() {
	b.doRepair(b)
}

func (c *VehicleBase) FillFuel() {
	fmt.Println("Fill with " + c.doFillFuel())
}
