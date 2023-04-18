package main

import (
	"testing"
)

var manager *ProtoManager

type Type1 struct {
	name string
}

func (t *Type1) clone() CloneAble {
	tc := *t
	return &tc
}

type Type2 struct {
	name string
}

func (t *Type2) clone() CloneAble {
	tc := *t
	return &tc
}

func TestClone(t *testing.T) {
	t1 := manager.Get("t1")

	t2 := t1.clone()

	if t1 == t2 {
		t.Fatal("error! get clone not working")
	}
}

func TestCloneFromManager(t *testing.T) {
	c := manager.Get("t1").clone()

	t1 := c.(*Type1)

	if t1.name != "type1" {
		t.Fatal("error")
	}

	d := manager.Get("t1").clone()
	d1 := d.(*Type1)
	if d1.name != "type1" {
		t.Fatal("error")
	}

	t2 := &Type1{"type2"}
	manager.Set("t1", t2)

	if t1.name != "type1" || d1.name != "type1" {
		t.Fatal("error")
	}

	e := manager.Get("t1").clone()
	e1 := e.(*Type1)

	if e1.name != "type2" {
		t.Fatal("error")
	}
}

func init() {
	manager = NewProtoTypeManager()
	t1 := &Type1{name: "type1"}
	manager.Set("t1", t1)
}
