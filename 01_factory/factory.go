package main

import "fmt"

// ========= 抽象层 ====================

// 水果类（抽象接口）

type Fruit interface {
	Show() // 接口的某方法
}

// 工厂类（抽象接口）

type AbstractFactory interface {
	CreateFruit() Fruit // 生产水果类（抽象）的生产器方法
}

// ========= 基础类模块 =============

type Apple struct {
}

func (apple *Apple) Show() {
	fmt.Println("i am a apple")
}

type Banana struct {
}

func (banana *Banana) Show() {
	fmt.Println("i am a banana")
}

type Pear struct {
}

func (pear *Pear) Show() {
	fmt.Println("i am a pear")
}

type JapanApple struct {
}

func (jp *JapanApple) Show() {
	fmt.Println("i am a japan apple")
}

// ======== 工厂模块 ===============

// 具体的 apple 模块

type AppleFactory struct {
	AbstractFactory
}

func (fac *AppleFactory) CreateFruit() Fruit {
	var fruit Fruit

	// 生产一个具体的apple
	fruit = new(Apple)
	return fruit
}

type BananaFactory struct {
	AbstractFactory
}

func (fac *BananaFactory) CreateFruit() Fruit {
	var fruit Fruit

	// 生产一个具体的banana
	fruit = new(Banana)
	return fruit
}

type PearFactory struct {
	AbstractFactory
}

func (fac *PearFactory) CreateFruit() Fruit {
	var fruit Fruit

	// 生产一个具体的pear
	fruit = new(Pear)
	return fruit
}

type JapanAppleFactory struct {
	AbstractFactory
}

func (fac *JapanAppleFactory) CreateFruit() Fruit {
	var fruit Fruit

	// 生产一个具体的japanApple
	fruit = new(JapanApple)
	return fruit
}

// ============= 业务逻辑层 ====================

func main() {
	// 需求1：需要一个具体的苹果对象
	// 1-先要一个具体的苹果工厂
	var appleFac AbstractFactory
	appleFac = new(AppleFactory)
	// 2-生产相对应的具体水果
	var apple Fruit
	apple = appleFac.CreateFruit()
	apple.Show()

	// 需求1：需要一个具体的香蕉对象
	// 1-先要一个具体的香蕉工厂
	var bananaFac AbstractFactory
	bananaFac = new(BananaFactory)
	// 2-生产相对应的具体水果
	var banana Fruit
	banana = bananaFac.CreateFruit()
	banana.Show()

	// 需求3：需要一个具体的梨对象
	// 1-先要一个具体的梨工厂
	var pearFac AbstractFactory
	pearFac = new(PearFactory)
	// 2-生产相对应的具体水果
	var pear Fruit
	pear = pearFac.CreateFruit()
	pear.Show()

	// 需求4：需要一个日本的苹果？？
	var japanAppleFac AbstractFactory
	japanAppleFac = new(JapanAppleFactory)
	var japanApple Fruit
	japanApple = japanAppleFac.CreateFruit()
	japanApple.Show()
}
