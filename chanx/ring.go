package chanx

type T interface{}

type RingChan struct {
	In     chan<- T
	Out    <-chan T
	buffer *RingBuffer
}

func (r RingChan) Len() int {
	return len(r.In) + len(r.Out) + r.buffer.Len()
}

// BuffLen return len of the buffer
func (r RingChan) BuffLen() int {
	return r.buffer.Len()
}

// NewRingChan return self
func NewRingChan(initCapacity int) RingChan {
	return NewRingChanSize(initCapacity, initCapacity, initCapacity)
}

func NewRingChanSize(initIncapacity, initOutCapacity, initBuffCapacity int) RingChan {
	in := make(chan T, initIncapacity)
	out := make(chan T, initOutCapacity)
	ch := RingChan{in, out, NewBuffer(initBuffCapacity)}

	go process(in, out, ch)

	return ch
}

func process(in, out chan T, ch RingChan) {
	defer close(out)
loop:
	for {
		val, ok := <-in
		if !ok {
			break loop
		}
		select {
		case out <- val:
			continue
		default:
		}
		ch.buffer.Write(val)
		for !ch.buffer.IsEmpty() {
			select {
			case val, ok := <-in:
				if !ok {
					break loop
				}
				ch.buffer.Write(val)
			case out <- ch.buffer.Peek():
				ch.buffer.Pop()
				if ch.buffer.IsEmpty() && ch.buffer.size > ch.buffer.initialSize {
					ch.buffer.Reset()
				}
			}
		}
	}
	for !ch.buffer.IsEmpty() {
		out <- ch.buffer.Pop()
	}
	ch.buffer.Reset()
}
