package main

import "fmt"

// ======== 抽象层 ============

// 水果类（抽象接口）

type Fruit interface {
	Show()
}

// ========= 基础类模块 ================

type Apple struct {
	Fruit
}

func (apple *Apple) Show() {
	fmt.Println("i am a apple")
}

type Banana struct {
	Fruit
}

func (banana *Banana) Show() {
	fmt.Println("i am a banana")
}

type Pear struct {
	Fruit
}

func (pear *Pear) Show() {
	fmt.Println("i am a pear")
}

// =============== 工厂模块 ===============

type Factory struct {
}

func (fac *Factory) CreateFruit(kind string) Fruit {
	var fruit Fruit

	if kind == "apple" {
		fruit = new(Apple)
	} else if kind == "banana" {
		fruit = new(Banana)
	} else if kind == "pear" {
		fruit = new(Pear)
	}

	return fruit
}

// =========== 业务逻辑层 ===============

func main() {
	factory := new(Factory)

	apple := factory.CreateFruit("apple")
	apple.Show()

	banana := factory.CreateFruit("banana")
	banana.Show()

	pear := factory.CreateFruit("pear")
	pear.Show()
}
