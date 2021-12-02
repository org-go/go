package hash

import (
	"fmt"
	"hash/crc32"
	"testing"
)

var m *Map

func init() {
	m = New(3, func(hash []byte) uint32 {
		return crc32.ChecksumIEEE(hash)
	})
}

func TestMap_Add(t *testing.T) {

	m.Add(`A`)
	m.Add(`B`)
	m.Add(`D`)

	fmt.Println(m)
}

func TestMap_Get(t *testing.T) {

	TestMap_Add(&testing.T{})
	A := m.Get(`A`)
	B := m.Get(`B`)
	C := m.Get(`C`)

	fmt.Println(A, B, C)

}
