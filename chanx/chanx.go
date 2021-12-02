package chanx

import "sync"

type msg struct {
	v  interface{}
	ok bool
}

type C interface {
	Send(v interface{}) (ok bool)

	Recv() (v interface{}, ok bool)

	Close() (ok bool)

	Wait()
}

type c struct {
	mu     sync.Mutex
	cond   *sync.Cond
	c      chan msg
	closed bool
}

func Make(length int) C {
	c := &c{c: make(chan msg, length)}
	c.cond = sync.NewCond(&c.mu)
	return c
}

func (c *c) Send(v interface{}) (ok bool) {
	defer func() { ok = recover() == nil }()
	c.c <- msg{v, ok}
	return ok
}

func (c *c) Recv() (v interface{}, ok bool) {
	select {
	case msg := <-c.c:
		return msg.v, msg.ok
	}
}

func (c *c) Close() (ok bool) {

	c.mu.Lock()
	defer c.mu.Unlock()
	defer func() { ok = recover() == nil }()
	close(c.c)
	c.closed = true
	c.cond.Broadcast()
	return ok

}

func (c *c) Wait() {

	c.mu.Lock()
	defer c.mu.Unlock()
	for {
		if c.closed {
			return
		}
		c.cond.Wait()
	}
}
