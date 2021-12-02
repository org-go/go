package model

import "fmt"

type adapter struct{}

func newAdapter() IFacadeImplement {
	return &adapter{}
}

func (adapter) T() {
	fmt.Println(`000`)
}

func (adapter) Adapter() {
	fmt.Println(`100`)
}
