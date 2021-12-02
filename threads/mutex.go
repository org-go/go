package threads

import (
	"acs-sdk-go/threads/safe"
	"sync"
)

type mux struct {
	m sync.Mutex
}

func (g *mux) Run(cf func()) {
	g.m.Lock()
	go func() {
		defer g.m.Unlock()
		cf()
	}()
}

func (g *mux) SafeRun(cf func()) {
	g.m.Lock()
	safe.GoSafe(func() {
		defer g.m.Unlock()
		cf()
	})

}

func (g *mux) Wait() {
	g.m.Lock()
}
