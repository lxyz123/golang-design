package main

import "fmt"

type Shoes interface {
	Show()
}

type AbstractFactory interface {
	CreateShoes() Shoes
}

type Nike struct {
	Shoes
}

func (n Nike) Show() {
	fmt.Println("nike shoes")
}

type Adidas struct {
	Shoes
}

func (a Adidas) Show() {
	fmt.Println("adidas shoes")
}

type Fila struct {
	Shoes
}

func (f Fila) Show() {
	fmt.Println("fila shoes")
}

type Camel struct {
	Shoes
}

func (c Camel) Show() {
	fmt.Println("camel shoes")
}

type NikeFactory struct {
	AbstractFactory
}

func (fac *NikeFactory) CreateShoes() Shoes {
	var shoe Shoes
	shoe = new(Nike)
	return shoe
}

type AdidasFactory struct {
	AbstractFactory
}

func (fac *AdidasFactory) CreateShoes() Shoes {
	var shoe Shoes
	shoe = new(Adidas)
	return shoe
}

type CamelFactory struct {
	AbstractFactory
}

func (fac *CamelFactory) CreateShoes() Shoes {
	var shoe Shoes
	shoe = new(Camel)
	return shoe
}

type FilaFactory struct {
	AbstractFactory
}

func (fac *FilaFactory) CreateShoes() Shoes {
	var shoe Shoes
	shoe = new(Fila)
	return shoe
}

func main() {
	var nikeFactory AbstractFactory
	nikeFactory = new(NikeFactory)
	var nike Shoes
	nike = nikeFactory.CreateShoes()
	nike.Show()

	var adiFactory AbstractFactory
	adiFactory = new(AdidasFactory)
	var adidas Shoes
	adidas = adiFactory.CreateShoes()
	adidas.Show()

	var filaFactory AbstractFactory
	filaFactory = new(FilaFactory)
	var fila Shoes
	fila = filaFactory.CreateShoes()
	fila.Show()

	// when need a new shoes type
	var camelFactory AbstractFactory
	camelFactory = new(CamelFactory)
	var camel Shoes
	camel = camelFactory.CreateShoes()
	camel.Show()
}
