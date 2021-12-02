package model

type IFacade interface {
	action(func())
}

func (f *Facade) action(fc func()) {
	defer func() {
		go fc()
	}()
	for _, o := range f.obs {
		o.T()
	}

}

func (f *Facade) facade(ios ...IFacadeImplement) IFacade {
	for _, o := range ios {
		f.obs = append(f.obs, o)
	}
	return f
}

// test

type IFacadeImplement interface {
	T()
}
