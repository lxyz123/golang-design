package facade

import "fmt"

type API interface {
	Test() string
}

type apiImpl struct {
	a AModuleAPI
	b BModuleAPI
}

func NewAPI() API {
	return &apiImpl{
		a: NewAModuleAPI(),
		b: NewBModuleAPI(),
	}
}

func (a *apiImpl) Test() string {
	aRet := a.a.TestA()
	bRet := a.b.TestB()
	return fmt.Sprintf("%s\n%s", aRet, bRet)
}

func NewAModuleAPI() AModuleAPI {
	return &aModuleImpl{}
}

type AModuleAPI interface {
	TestA() string
}

type aModuleImpl struct {
}

func (*aModuleImpl) TestA() string {
	return "A module running"
}

type BModuleAPI interface {
	TestB() string
}

type bModuleImpl struct {
}

func NewBModuleAPI() BModuleAPI {
	return &bModuleImpl{}
}

func (*bModuleImpl) TestB() string {
	return "B module running"
}
