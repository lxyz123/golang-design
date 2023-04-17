package main

type Component interface {
	Calc() int
}

type ConcreteComponent struct{}

func (*ConcreteComponent) Calc() int {
	return 0
}

type MulDecorator struct {
	Component
	num int
}

func WrapMulDecorator(c Component, num int) Component {
	return &MulDecorator{
		Component: c,
		num:       num,
	}
}

func (d *MulDecorator) Calc() int {
	return d.Component.Calc() * d.num
}

type AddDecorator struct {
	Component
	num int
}

func WrapAddDecorator(c Component, num int) Component {
	return &AddDecorator{
		Component: c,
		num:       num,
	}
}

func (d *AddDecorator) Calc() int {
	return d.Component.Calc() + d.num
}

type DivDecorator struct {
	Component
	num int
}

func WrapDivDecorator(c Component, num int) Component {
	return &DivDecorator{
		Component: c,
		num:       num,
	}
}

func (d *DivDecorator) Calc() int {
	return d.Component.Calc() / d.num
}

type SubDecorator struct {
	Component
	num int
}

func WrapSubDecorator(c Component, num int) Component {
	return &SubDecorator{
		Component: c,
		num:       num,
	}
}

func (s *SubDecorator) Calc() int {
	return s.Component.Calc() - s.num
}
