package fx

import "fmt"

type (
	FilterFunc func(interface{}) bool

	SafeFunc func(filterFunc FilterFunc) interface{}

	TempChan chan interface{}

	TpChanFunc func(chan int) interface{}
)

type Fx struct{}

func (f *Fx) filter(filterFunc FilterFunc) *Fx {
	if filterFunc(fc()) {
	}
	return f
}

func (f *Fx) safe(safeFunc SafeFunc) *Fx {
	i := safeFunc(func(i interface{}) bool {
		if i == fc() {
			return true
		}
		return false
	})
	fmt.Println(i)
	return f
}

func fc() int {
	return 3
}
