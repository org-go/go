package threads

import (
	"acs-sdk-go/threads/safe"
	"sync"
)

type groupx struct {
	group sync.WaitGroup
}

func NewGroup() *groupx {
	return &groupx{group: sync.WaitGroup{}}
}

func (g *groupx) Run(cf func()) {
	g.group.Add(1)
	go func() {
		defer g.group.Done()
		cf()
	}()
}

func (g *groupx) SafeRun(cf func()) {
	g.group.Add(1)
	safe.GoSafe(func() {
		defer g.group.Done()
		cf()
	})
}

func (g *groupx) Wait() {
	g.group.Wait()
}
