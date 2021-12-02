package model

type IOperator interface {
	set_A(int)
	set_B(int)
	get() int
	action(func())
}

type IOperatorFactory interface {
	create() IOperator
}

func (o *Operator) set_A(i int) {
	o.a = i
}
func (o *Operator) set_B(i int) {
	o.b = i
}

func (o *Operator) get() int {
	return o.a + o.b
}

func (o Operator) action(f func()) {
	go f()
}
