package main

import (
	"fmt"
)

type Handler interface {
	Handle()
	next(handler Handler)
}

type IDCheck struct {
	handler Handler
}

func (id *IDCheck) Handle() {
	id.next(id.handler)
	fmt.Println("id check pass")
}

func (id *IDCheck) next(handler Handler) {
	if id.handler != nil {
		id.handler.Handle()
	}
}

type PackageCheck struct {
	handler Handler
}

func (pc *PackageCheck) Handle() {
	pc.next(pc.handler)
	fmt.Println("package check pass")
}

func (pc *PackageCheck) next(handler Handler) {
	if pc.handler != nil {
		pc.handler.Handle()
	}
}

type BodyCheck struct {
	handler Handler
}

func (bc *BodyCheck) Handle() {
	bc.next(bc.handler)
	fmt.Println("body check pass")
}

func (bc *BodyCheck) next(handler Handler) {
	if bc.handler != nil {
		bc.handler.Handle()
	}
}

func main() {
	idCheck := &IDCheck{}
	packageCheck := &PackageCheck{}
	bodyCheck := &BodyCheck{}
	idCheck.handler = packageCheck
	packageCheck.handler = bodyCheck
	idCheck.Handle()
}
