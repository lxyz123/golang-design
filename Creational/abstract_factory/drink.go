package main

import "fmt"

type Drink interface {
	Choose()
}

type ChineseTea struct {
}

func (ct ChineseTea) Choose() {
	fmt.Println("choose chinese tea")
}

type ChineseCoffee struct {
}

func (cf ChineseCoffee) Choose() {
	fmt.Println("choose chinese coffee")
}

type DrinkFactory interface {
	ChooseDrink(kind string) Drink
}

type ChineseDrinkFactory struct {
}

func (cd ChineseDrinkFactory) ChooseDrink(kind string) Drink {
	if kind == "tea" {
		return &ChineseTea{}
	} else if kind == "coffee" {
		return &ChineseCoffee{}
	} else {
		return nil
	}
}

type AmericaTea struct {
}

func (at AmericaTea) Choose() {
	fmt.Println("choose america tea")
}

type AmericaCoffee struct {
}

func (ac AmericaCoffee) Choose() {
	fmt.Println("choose america coffee")
}

type AmericaDrink interface {
	ChooseDrink(kind string) Drink
}

type AmericaDrinkFactory struct {
}

func (cd AmericaDrinkFactory) ChooseDrink(kind string) Drink {
	if kind == "tea" {
		return &AmericaTea{}
	} else if kind == "coffee" {
		return &AmericaCoffee{}
	} else {
		return nil
	}
}

type DrinkFactoryStore struct {
	drinkFactory DrinkFactory
}

func (df DrinkFactoryStore) chooseDrink(kind string) Drink {
	return df.drinkFactory.ChooseDrink(kind)
}

func main() {
	factory := new(DrinkFactoryStore)
	factory.drinkFactory = new(AmericaDrinkFactory)
	americaDrink := factory.chooseDrink("tea")
	americaDrink.Choose()

	factory.drinkFactory = new(ChineseDrinkFactory)
	chineseDrink := factory.chooseDrink("coffee")
	chineseDrink.Choose()
}
