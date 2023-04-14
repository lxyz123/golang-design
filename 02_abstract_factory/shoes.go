package main

import "fmt"

type AbstractNike interface {
	ShowNike()
}

type AbstractAdidas interface {
	ShowAdidas()
}

type AbstractFila interface {
	ShowFila()
}

type AbstractFactory interface {
	BuyNike() AbstractNike
	BuyAdidas() AbstractAdidas
	BuyFila() AbstractFila
}

// made in Thailand

type ThailandNike struct {
}

func (c ThailandNike) ShowNike() {
	fmt.Println("nike made in Thailand")
}

type ThailandAdidas struct{}

func (c ThailandAdidas) ShowAdidas() {
	fmt.Println("adidas made in Thailand")
}

type ThailandFila struct {
}

func (c ThailandFila) ShowFila() {
	fmt.Println("fila made in Thailand")
}

type ThailandFactory struct {
}

func (thai ThailandFactory) BuyNike() AbstractNike {
	var nike AbstractNike
	nike = new(ThailandNike)
	return nike
}

func (thai ThailandFactory) BuyAdidas() AbstractAdidas {
	var adidas AbstractAdidas
	adidas = new(ThailandAdidas)
	return adidas
}
func (thai ThailandFactory) BuyFila() AbstractFila {
	var fila AbstractFila
	fila = new(ThailandFila)
	return fila
}

// made in America

type AmericaNike struct {
}

func (c AmericaNike) ShowNike() {
	fmt.Println("nike made in America")
}

type AmericaAdidas struct{}

func (c AmericaAdidas) ShowAdidas() {
	fmt.Println("adidas made in America")
}

type AmericaFila struct {
}

func (c AmericaFila) ShowFila() {
	fmt.Println("fila made in America")
}

type AmericaFactory struct {
}

func (am AmericaFactory) BuyNike() AbstractNike {
	var nike AbstractNike
	nike = new(AmericaNike)
	return nike
}

func (am AmericaFactory) BuyAdidas() AbstractAdidas {
	var adidas AbstractAdidas
	adidas = new(AmericaAdidas)
	return adidas
}
func (am AmericaFactory) BuyFila() AbstractFila {
	var fila AbstractFila
	fila = new(AmericaFila)
	return fila
}

func main() {
	// buy America product
	var americaFactory AbstractFactory
	americaFactory = new(AmericaFactory)

	// buy America nike
	var aNike AbstractNike
	aNike = americaFactory.BuyNike()
	aNike.ShowNike()

	var aAdidas AbstractAdidas
	aAdidas = americaFactory.BuyAdidas()
	aAdidas.ShowAdidas()

	// buy Thailand product
	var thaiFactory AbstractFactory
	thaiFactory = new(ThailandFactory)

	var tFila AbstractFila
	tFila = thaiFactory.BuyFila()
	tFila.ShowFila()
}
