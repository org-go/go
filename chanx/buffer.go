package chanx

import "errors"

var ErrIsEmpty = errors.New(`ring buff is empty`)

type RingBuffer struct {
	buff        []T
	initialSize int
	size        int
	r           int
	w           int
}

// NewBuffer return self
func NewBuffer(initialSize int) *RingBuffer {
	if initialSize <= 0 {
		panic(`initialSize must be great than zero`)
	}
	if initialSize == 1 {
		initialSize = 2
	}
	return &RingBuffer{
		buff:        make([]T, initialSize),
		initialSize: initialSize,
		size:        initialSize,
	}
}

func (r *RingBuffer) Read() (T, error) {

	if r.r == r.w {
		return nil, ErrIsEmpty
	}
	v := r.buff[r.r]
	r.r++
	if r.r == r.size {
		r.r = 0
	}
	return v, nil

}

func (r *RingBuffer) Pop() T {

	val, err := r.Read()
	if err == ErrIsEmpty {
		panic(ErrIsEmpty.Error())
	}
	return val

}

// Peek return T
func (r *RingBuffer) Peek() T {

	if r.r == r.w {
		panic(ErrIsEmpty.Error())
	}
	return r.buff[r.r]

}

func (r *RingBuffer) Write(v T) {
	r.buff[r.w] = v
	r.w++
	if r.w == r.size {
		r.w = 0
	}
	if r.w == r.r {
		r.grow()
	}
}

func (r *RingBuffer) grow() {

	var size int
	if r.size < 1024 {
		size = r.size * 2
	} else {
		size = r.size + r.size/4
	}
	buff := make([]T, size)
	copy(buff[0:], r.buff[r.r:])
	copy(buff[r.size-r.r:], r.buff[0:r.r])
	r.r = 0
	r.w = r.size
	r.size = size
	r.buff = buff

}

func (r *RingBuffer) IsEmpty() bool {
	return r.r == r.w
}

func (r *RingBuffer) Len() int {

	if r.r == r.w {
		return 0
	}
	if r.w > r.r {
		return r.w - r.r
	}
	return r.size - r.r + r.w

}

func (r *RingBuffer) Reset() {
	r.r, r.w = 0, 0
	r.size = r.initialSize
	r.buff = make([]T, r.initialSize)
}
