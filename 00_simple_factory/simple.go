package main

import "fmt"

// API is interface
type API interface {
	Say(name string) string
}

// NewAPI return Api instance by type
func NewAPI(t int) API {
	if t == 1 {
		return &hiAPI{}
	} else if t == 2 {
		return &helloAPI{}
	}
	return nil
}

// hiAPI is one of API implement
type hiAPI struct{}

// Say hi to name
func (*hiAPI) Say(name string) string {
	return fmt.Sprintf("Hi, %s", name)
}

// HelloAPI is another API implement
type helloAPI struct{}

// Say hello to name
func (*helloAPI) Say(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}

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

type Factory struct{}

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
