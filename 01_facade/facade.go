package main

import "fmt"

type Buy struct{}

func (b *Buy) buy() {
	fmt.Println("buy food")
}

type Wash struct{}

func (w *Wash) wash() {
	fmt.Println("wash food")
}

type Cook struct{}

func (w *Cook) cook() {
	fmt.Println("cook food")
}

type lunch struct {
	b *Buy
	w *Wash
	c *Cook
}

// HotPot eat hot pot need to buy、 wash、 cook
func (l *lunch) HotPot() {
	l.b.buy()
	l.w.wash()
	l.c.cook()
}

// Dumplings eat dumplings only need cook, assume have dumplings at home
func (l *lunch) Dumplings() {
	l.c.cook()
}

func (l *lunch) GoOutToEat() {
	l.b.buy()
}

func main() {
	// when I want to eat hot pot but not use facade design
	b := new(Buy)
	b.buy()
	w := new(Wash)
	w.wash()
	c := new(Cook)
	c.cook()

	fmt.Println("----------------------------")

	// use facade design
	l := lunch{
		b: new(Buy),
		w: new(Wash),
		c: new(Cook),
	}
	l.HotPot()
}
