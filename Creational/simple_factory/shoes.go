package main

import "fmt"

type Shoes interface {
	Show()
}

type Nike struct {
	Shoes
}

func (n *Nike) Show() {
	fmt.Println("buy nike shoes")
}

type Adidas struct {
	Shoes
}

func (a *Adidas) Show() {
	fmt.Println("buy adidas shoes")
}

type Fila struct {
	Shoes
}

func (f *Fila) Show() {
	fmt.Println("buy fila shoes")
}

type Factory struct{}

func (fac *Factory) BuyShoesFactory(kind string) Shoes {
	var shoes Shoes

	if kind == "nike" {
		shoes = new(Nike)
	} else if kind == "adidas" {
		shoes = new(Adidas)
	} else if kind == "fila" {
		shoes = new(Fila)
	}

	return shoes
}

func main() {
	factory := new(Factory)

	nike := factory.BuyShoesFactory("nike")
	nike.Show()

	adidas := factory.BuyShoesFactory("adidas")
	adidas.Show()

	fila := factory.BuyShoesFactory("fila")
	fila.Show()
}
