package main

type CloneAble interface {
	clone() CloneAble
}

type ProtoManager struct {
	protoTypes map[string]CloneAble
}

func NewProtoTypeManager() *ProtoManager {
	return &ProtoManager{map[string]CloneAble{}}
}

func (p *ProtoManager) Set(name string, cloneAble CloneAble) {
	p.protoTypes[name] = cloneAble
}

func (p *ProtoManager) Get(name string) CloneAble {
	return p.protoTypes[name].clone()
}
