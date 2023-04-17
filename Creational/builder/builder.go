package main

import "fmt"

type Car interface {
	PrepareAccessories()
	Assemble()
	TestPerformance()
}

type Creator struct {
	car Car
}

func (c *Creator) Construct() *Car {
	c.car.PrepareAccessories()
	c.car.Assemble()
	c.car.TestPerformance()
	return &c.car
}

type BenzCar struct {
	Accessory           string
	AssembleProduction  string
	PerformanceStandard string
}

func (b *BenzCar) PrepareAccessories() {
	fmt.Println("prepare benz accessories")
	b.Accessory = "benz Accessory"
}

func (b *BenzCar) Assemble() {
	fmt.Println("assemble benz")
	b.AssembleProduction = "benz introduce"
}

func (b *BenzCar) TestPerformance() {
	fmt.Println("test benz performance")
	b.PerformanceStandard = "benz performance"
}

func main() {
	b := BenzCar{}
	d := &Creator{&b}
	car := d.Construct()
	fmt.Printf("%+v", *car)
}
