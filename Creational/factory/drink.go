package main

import "fmt"

type Drink interface {
	Choose()
}

type ChooseDrink interface {
	OneDrink() Drink
}

type Coffee struct {
}

func (c Coffee) Choose() {
	fmt.Println("Choose coffee")
}

type Cola struct {
}

func (c Cola) Choose() {
	fmt.Println("Choose cola")
}

type Tea struct {
}

func (t Tea) Choose() {
	fmt.Println("Choose tea")
}

type CoffeeFactory struct {
	ChooseDrink
}

func (cf CoffeeFactory) OneDrink() Drink {
	var drink Drink
	drink = new(Coffee)
	return drink
}

type TeaFactory struct {
	ChooseDrink
}

func (tf TeaFactory) OneDrink() Drink {
	var drink Drink
	drink = new(Tea)
	return drink
}

type ColaFactory struct {
	ChooseDrink
}

func (cf ColaFactory) OneDrink() Drink {
	var drink Drink
	drink = new(Cola)
	return drink
}

func main() {
	var coffeeFactory ChooseDrink
	coffeeFactory = new(CoffeeFactory)
	coffeeFactory.OneDrink().Choose()

	var teaFactory ChooseDrink
	teaFactory = new(TeaFactory)
	teaFactory.OneDrink().Choose()

	var colaFactory ChooseDrink
	colaFactory = new(ColaFactory)
	colaFactory.OneDrink().Choose()

	// when need a new drink, such as Sprite
	var spriteFactory ChooseDrink
	spriteFactory = new(SpriteFactory)
	spriteFactory.OneDrink().Choose()
}

type Sprite struct {
}

func (s Sprite) Choose() {
	fmt.Println("choose sprite")
}

type SpriteFactory struct {
	ChooseDrink
}

func (sf SpriteFactory) OneDrink() Drink {
	var drink Drink
	drink = new(Sprite)
	return drink
}
