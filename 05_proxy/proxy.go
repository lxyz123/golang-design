package main

import "fmt"

// 抽象主题
type WatchMessi interface {
	//买机票
	BuyAirTicket()
	//买球场门票
	BuyStadiumTicket()
}

// 具体主题
type FansFromChina struct {
}

func (f *FansFromChina) BuyAirTicket() {
	fmt.Println("用黄牛买的机票出国")
}

func (f *FansFromChina) BuyStadiumTicket() {
	fmt.Println("用黄牛代买的门票进场")
}

type HuangNiu struct {
	fans WatchMessi
}

func NewProxy(fans WatchMessi) WatchMessi {
	return &HuangNiu{fans}
}

func (h HuangNiu) BuyAirTicket() {
	fmt.Println("黄牛帮买机票")
	h.fans.BuyAirTicket()
}

func (h HuangNiu) BuyStadiumTicket() {
	fmt.Println("黄牛帮买门票")
	h.fans.BuyStadiumTicket()
}

func main() {
	// 球迷想看梅西踢球，找黄牛帮忙买票
	huangniu := NewProxy(new(FansFromChina))
	huangniu.BuyAirTicket()
	huangniu.BuyStadiumTicket()
}
