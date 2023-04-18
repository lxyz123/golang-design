package main

import "fmt"

type Lego interface {
	Head()
	Body()
	Foot()
}

type MakeLego struct {
	lego Lego
}

func NewMakeLego(lego Lego) *MakeLego {
	return &MakeLego{lego: lego}
}

func (ml *MakeLego) Constructor() {
	ml.lego.Head()
	ml.lego.Body()
	ml.lego.Foot()
}

type MonsterLego struct {
	result string
}

func (ml *MonsterLego) Head() {
	ml.result += "head install "
}

func (ml *MonsterLego) Body() {
	ml.result += "body install "
}

func (ml *MonsterLego) Foot() {
	ml.result += "foot install "
}

func (ml *MonsterLego) GetResult() string {
	return ml.result
}

func main() {
	ml := &MonsterLego{}
	director := NewMakeLego(ml)
	director.Constructor()
	fmt.Printf("%v", *ml)
}
