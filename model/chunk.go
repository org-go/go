package model

type (
	IChunk interface {
		chunk(fn func(i int, size int) (int, func()))
	}
)

/**
 * chunk
 * @Description: callback function after slice handler data
 * @receiver c
 * @param fn	callback
 */
func (c *Chunk) chunk(fn func(i int, size int) (int, func())) {
	for i := 0; ; {
		size := 0
		var callback func()
		size, callback = fn(i, size)
		if size == 0 {
			break
		} else {
			callback()
			i++
		}
	}
}
