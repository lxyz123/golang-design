package main

import (
	"fmt"
	"testing"
)

type Rainy struct {
}

func (r *Rainy) RainyClothes() {
	fmt.Println("rainy clothes")
}

type Sunny struct {
}

func (s *Sunny) SunnyClothes() {
	fmt.Println("sunny clothes")
}

func Test(t *testing.T) {
	rc := new(Rainy)
	rc.RainyClothes()
	sc := new(Sunny)
	sc.SunnyClothes()
}
