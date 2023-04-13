package main

import "fmt"

// 饿汉式

type singleton struct{}

var instance *singleton = new(singleton)

func GetInstance() *singleton {
	return instance
}

// GetInstance 懒汉式

//type singleton struct {}
//var instance *singleton
//func GetInstance() *singleton {
// 只有首次调用GetInstance()时，才会生成这个单例的实例
//	if instance == nil {
//		instance = new(singleton)
//		return instance
//	}
//	return instance
//}

func (s *singleton) SomeThing() {
	fmt.Println("function of singleton")
}

func main() {
	s := GetInstance()
	s.SomeThing()
}
