package hash

import (
	"hash/crc32"
	"sort"
	"strconv"
)

type Hash func(hash []byte) uint32

type Map struct {
	hash     Hash
	replicas int
	keys     []int
	hashMap  map[int]string
}

func New(replicas int, fn Hash) *Map {

	maps := &Map{
		hash:     fn,
		replicas: replicas,
		hashMap:  make(map[int]string),
	}
	if maps.hash == nil {
		maps.hash = crc32.ChecksumIEEE
	}
	return maps

}

func (m *Map) Add(keys ...string) {

	for _, key := range keys {
		for i := 0; i < m.replicas; i++ {
			hash := int(m.hash([]byte(key + strconv.Itoa(i))))
			m.keys = append(m.keys, hash)
			m.hashMap[hash] = key
		}
	}
	sort.Ints(m.keys)

}

func (m *Map) Get(key string) string {

	if m.keys == nil {
		return ``
	}
	hash := int(m.hash([]byte(key)))
	idx := sort.Search(len(m.keys), func(i int) bool { return m.keys[i] >= hash })
	if idx == len(m.keys) {
		idx = 0
	}

	return m.hashMap[m.keys[idx]]

}
