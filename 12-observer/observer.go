package main

import (
	"fmt"
	"time"
)

// 事件类型

type Event struct {
	Data string
}

type Observer interface {
	Update(event *Event)
}

type Subject interface {
	Register(Observer)
	Deregister(Observer)
	Notify(*Event)
}

type CreateObserver struct {
	Id int
}

type CreateSubject struct {
	Observers map[Observer]struct{}
}

func (co *CreateObserver) Update(e *Event) {
	fmt.Printf("observer %d receive %s msg\n", co.Id, e)
}

func (cs *CreateSubject) Register(ob Observer) {
	cs.Observers[ob] = struct{}{}
}

func (cs *CreateSubject) Deregister(ob Observer) {
	delete(cs.Observers, ob)
}

func (cs *CreateSubject) Notify(event *Event) {
	for ob, _ := range cs.Observers {
		ob.Update(event)
	}
}

func main() {
	cs := &CreateSubject{Observers: make(map[Observer]struct{})}
	cObserver1 := &CreateObserver{1}
	cObserver2 := &CreateObserver{2}
	cs.Register(cObserver1)
	cs.Register(cObserver2)

	for i := 0; i < 5; i++ {
		e := &Event{fmt.Sprintf("msg %d", i)}
		cs.Notify(e)
		time.Sleep(time.Second)
	}
}
