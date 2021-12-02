package model

type IComponent interface {
	Calc() int
}

type ConcreteComponent struct{}

func (*ConcreteComponent) Calc() int {
	return 0
}

type MulDecorator struct {
	IComponent
	num int
}

func WarpMulDecorator(c IComponent, num int) IComponent {
	return &MulDecorator{
		IComponent: c,
		num:        num,
	}
}

func (d *MulDecorator) Calc() int {
	return d.IComponent.Calc() * d.num
}

type AddDecorator struct {
	IComponent
	num int
}

func WarpAddDecorator(c IComponent, num int) IComponent {
	return &AddDecorator{
		IComponent: c,
		num:        num,
	}
}

func (d *AddDecorator) Calc() int {
	return d.IComponent.Calc() + d.num
}
