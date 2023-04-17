package main

import "fmt"

// a messi fans want to go abroad to look messi

type WatchMessi interface {
	BuyAirTicket()
	BuyStadiumTicket()
}

type FansFromChina struct {
}

func (f *FansFromChina) BuyAirTicket() {
	fmt.Println("air plane")
}

func (f *FansFromChina) BuyStadiumTicket() {
	fmt.Println("stadium ticket")
}

type SellTicketProxy struct {
	fans WatchMessi
}

func NewProxy(fans WatchMessi) WatchMessi {
	return &SellTicketProxy{fans}
}

func (h SellTicketProxy) BuyAirTicket() {
	fmt.Println("help buy air ticket")
	h.fans.BuyAirTicket()
}

func (h SellTicketProxy) BuyStadiumTicket() {
	fmt.Println("help buy stadium ticket")
	h.fans.BuyStadiumTicket()
}

func main() {
	// without tickets, ask for proxy to help buy tickets
	sell := NewProxy(new(FansFromChina))
	sell.BuyAirTicket()
	sell.BuyStadiumTicket()
}
