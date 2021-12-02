package model

import (
	"fmt"
	"testing"
)

const chunk = 1000

// broker notify.
func TestBroker_run(t *testing.T) {

	b := NewBroker()
	b.attach(Reader{name: `1`})
	b.attach(Reader{name: `2`})
	c := make(chan interface{}, 1)
	defer close(c)
	c <- 1
	b.registry(c)

}

// prototype
func TestCloneable_run(t *testing.T) {
	p := NewCloneable()
	p.set(`eros1`, &Reader{name: `s1`})
	p.set(`eros2`, &Reader{name: `s2`})
	clone := p.get(`eros1`).clone()
	reader := clone.(*Reader)
	fmt.Println(fmt.Sprintf(`%p - %p - %p`, p, reader, p.get(`eros1`).(*Reader)))
}

// chunk
func TestChunk_run(t *testing.T) {
	c := Chunk{}
	data := ``
	c.chunk(func(i int, size int) (int, func()) {
		data, size = new(Reader).chunk(chunk*i+1, chunk)
		return size, func() {
			if size > 0 {
				fmt.Println(data, i, size)
			}
		}
	})
}

// factory
func TestChunk_factory(t *testing.T) {
	f := new(Reader)
	factory_compose(f, 1, 2)
}

func factory_compose(operator IOperatorFactory, a, b int) {
	ops := operator.create()
	ops.set_A(a)
	ops.set_B(b)
	get := ops.get()
	fmt.Println(get, ops)
}

type selfFacade Facade

// facade
func TestFacade_run(t *testing.T) {
	f := new(Facade)
	f.facade(&Reader{}).action(func() {
		for _, o := range f.obs {
			o.T()
		}
	})

}

type Reader struct {
	name string
	o    Operator
}

func (r Reader) create() IOperator {
	return &r.o
}

func (r *Reader) chunk(offset, chunk int) (string, int) {
	return `data`, 100
}

func (r *Reader) clone() ICloneable {
	rs := *r
	return &rs

}

func (r Reader) action(b *Broker) {
	fmt.Println(fmt.Sprintf(`%s receive %v`, r.name, b.call))
}

func (r Reader) T() {
	fmt.Println(`----- facade .... implement....`)
}

func TestComposite_run(t *testing.T) {
	root := NewComponent(CompositeNode, `root`)
	c1 := NewComponent(CompositeNode, `c1`)
	c2 := NewComponent(CompositeNode, `c2`)
	c3 := NewComponent(CompositeNode, `c3`)
	l1 := NewComponent(LeafNode, `l1`)
	l2 := NewComponent(LeafNode, `l2`)
	l3 := NewComponent(LeafNode, `l3`)
	root.AddChild(c1)
	root.AddChild(c2)
	c1.AddChild(c3)
	c1.AddChild(l1)
	c2.AddChild(l2)
	c2.AddChild(l3)
	root.Print(``)
	fmt.Println(root)
}

func TestFlyWeight_run(t *testing.T) {
	viewer1 := NewImageViewer(`eros.png`)
	viewer2 := NewImageViewer(`eros.png`)
	viewer1.Display()
	viewer2.Display()
	fmt.Println(viewer1, viewer2)
}

func TestBridge_run(t *testing.T) {

	message := NewCommonMessage(&MessageSMS{})
	message.SendMessage(`eros`, `smoke`)
	fmt.Println(`----------`, message.method)

	commonMessage := NewCommonMessage(ViaEmail())
	commonMessage.SendMessage(`eros1`, `smoke1`)
	fmt.Println(`----------`, commonMessage.method)

	urgencyMessage := NewUrgencyMessage(&MessageEmail{})
	urgencyMessage.SendMessage(`eros2`, `smoke2`)
	fmt.Println(`-----------`, urgencyMessage.method)

	urgencyMessage2 := NewUrgencyMessage(ViaEmail())
	urgencyMessage2.SendMessage(`eros3`, `smoke3`)
	fmt.Println(`-----------`, urgencyMessage2.method)
}

func TestStrategy_run(t *testing.T) {

	payment := NewPayment(`eros`, `1001010`, 10, &Cash{})
	payment.Pay()

	newPayment := NewPayment(`eros`, `1001010`, 10, &Bank{})
	newPayment.Pay()

}

func TestMemento_run(t *testing.T) {

	game := Game{
		hp: 10,
		mp: 10,
	}
	game.Status()
	save := game.Save()
	game.Play(-2, -3)
	game.Status()
	game.Load(save)
	game.Status()
	fmt.Println(save)
}

func TestInterpreter_run(t *testing.T) {
	parser := &Parser{}
	parser.Parse(`1 + 2 + 3 - 4 + 5 - 6`)
	res := parser.Result().Interpret()
	fmt.Println(res)
}

// 多重
func TestChain_run(t *testing.T) {
	pro := NewProjectManagerChain()
	dep := NewDepManagerChain()
	gen := NewGeneralManagerChain()
	fmt.Println(pro, dep, gen)
	pro.SetSuccessor(dep)
	dep.SetSuccessor(gen)
	fmt.Println(pro, dep, gen)
	var c Manager = pro
	c.HandleFeeRequest(`bob`, 400)
	c.HandleFeeRequest(`tom`, 1400)
	c.HandleFeeRequest(`ada`, 10000)
	c.HandleFeeRequest(`floar`, 400)
}

//
func TestVisitor_run(t *testing.T) {
	c := &CustomerCol{}
	c.Add(NewEnterpriseCustomer(`a eros`))
	c.Add(NewEnterpriseCustomer(`b eros`))
	c.Add(NewIndividualCustomer(`1 eros`))
	c.Accept(&ServiceRequestVisitor{})
}
