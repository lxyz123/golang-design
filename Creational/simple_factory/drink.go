package main

import "fmt"

type Drink interface {
	Choose()
}

type Cola struct{}

func (c Cola) Choose() {
	fmt.Println("I choose cola")
}

type Coffee struct{}

func (c Coffee) Choose() {
	fmt.Println("I choose coffee")
}

type Tea struct{}

func (t Tea) Choose() {
	fmt.Println("I choose tea")
}

type ChooseDrinkFactory struct{}

func (c ChooseDrinkFactory) NewChooseDrinkFactory(kind string) Drink {
	var drink Drink
	if kind == "tea" {
		drink = new(Tea)
	} else if kind == "coffee" {
		drink = new(Coffee)
	} else if kind == "cola" {
		drink = new(Cola)
	} else {
		fmt.Println("unknown drink")
	}
	return drink
}

func main() {
	factory := new(ChooseDrinkFactory)

	coffee := factory.NewChooseDrinkFactory("coffee")
	coffee.Choose()

	tea := factory.NewChooseDrinkFactory("tea")
	tea.Choose()

	cola := factory.NewChooseDrinkFactory("cola")
	cola.Choose()
}
