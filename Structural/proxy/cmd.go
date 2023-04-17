package main

import "fmt"

type UseMachine interface {
	Login()
}

type User struct {
}

func (u User) Login() {
	fmt.Println("user login machine success")
}

type JumpServer struct {
	login UseMachine
}

func NewLogin(user UseMachine) UseMachine {
	return &JumpServer{user}
}

func (j JumpServer) Login() {
	fmt.Println("jump server help login target machine")
	j.login.Login()
}

func main() {
	jump := NewLogin(new(User))
	jump.Login()
}
