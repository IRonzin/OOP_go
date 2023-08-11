package oop

type IVehicle interface {
	//перекрытие метода
	Move()

	//расширение функциональности родителя
	GetInfo() string

	//полиморфизм через интерфейсы
	Repair()
	repairBy() string

	//полиморфизм в функциональном стиле
	FillFuel()
}
