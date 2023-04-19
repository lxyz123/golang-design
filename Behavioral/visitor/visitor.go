package main

import "fmt"

type Visitor interface {
	VisitConcreteElementA(a *ConcreteElementA)
	VisitConcreteElementB(b *ConcreteElementB)
}

type ConcreteVisitor1 struct {
}

func (*ConcreteVisitor1) VisitConcreteElementA(a *ConcreteElementA) {
	fmt.Println("concreteVisitor1 visitConcreteElementA")
}

func (*ConcreteVisitor1) VisitConcreteElementB(b *ConcreteElementB) {
	fmt.Println("concreteVisitor2 visitConcreteElementB")
}

type ConcreteVisitor2 struct {
}

func (*ConcreteVisitor2) VisitConcreteElementA(a *ConcreteElementA) {
	fmt.Println("concreteVisitor2 visitConcreteElementA")
}

func (*ConcreteVisitor2) VisitConcreteElementB(b *ConcreteElementB) {
	fmt.Println("concreteVisitor2 visitConcreteElementB")
}

type Element interface {
	Accept(visitor Visitor)
}

type ConcreteElementA struct {
}

func (cea *ConcreteElementA) Accept(visitor Visitor) {
	visitor.VisitConcreteElementA(cea)
}

type ConcreteElementB struct {
}

func (ceb *ConcreteElementB) Accept(visitor Visitor) {
	visitor.VisitConcreteElementB(ceb)
}

func main() {
	var elements []Element
	elements = append(elements, &ConcreteElementA{})
	elements = append(elements, &ConcreteElementB{})
	for _, item := range elements {
		cv1 := &ConcreteVisitor1{}
		cv2 := &ConcreteVisitor2{}
		item.Accept(cv1)
		item.Accept(cv2)
	}
}
