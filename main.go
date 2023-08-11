package main

import (
	"fmt"

	"oop/oop"
)

func main() {

	/*//var d int = 10
	er := errors.New("err123")
	var f rune = 123
	changer := calc.NewChanger(3)
	changer.Mul(2)
	fmt.Printf("Changer state is %v \n", changer.GetState())
	switch g := f; g {
	case 124:
		fmt.Println(g)
	default:
		fmt.Println("unknown")
	}

	var b oop.Vehicle
	var in any
	in = "str"
	in = 123

	switch v := in.(type) {
	case int:
		fmt.Println(v + 123)
	case string:
		fmt.Println(v + "345")
	}

	fmt.Println(er.Error(), f)
	fmt.Println(reflect.TypeOf(b))*/

	var b = oop.NewCar("1234", "petrol 95")
	var p = oop.NewPlane("5555", "keros", func() string {
		return "keros 65749369"
	} /*nil*/)
	Drive(b)
	Drive(p)
}

func Drive(v oop.IVehicle) {
	fmt.Println(v.GetInfo())
	v.FillFuel()
	v.Move()
	v.Repair()
}
