package algorithm

import (
	"flag"
	"math/rand"
)

var (
	rounds   = flag.Int("r", 10000, "rounds")
	subset   = flag.Int("s", 30, "subset")
	elements = flag.Int("e", 1000, "elements")
)

// 平衡
func simpleSubset(set []string, sub int) []string {
	rand.Shuffle(len(set), func(i, j int) {
		set[i], set[j] = set[j], set[i]
	})
	if len(set) <= sub {
		return set
	}

	return set[:sub]
}

func complexSubset(set []string, sub int) []string {
	if len(set) <= sub {
		rand.Shuffle(len(set), func(i, j int) {
			set[i], set[j] = set[j], set[i]
		})
		return set
	}

	// group clients into rounds, each round uses the same shuffled list
	count := uint64(len(set) / sub)

	clientID := rand.Int63()
	round := clientID / int64(count)

	r := rand.New(rand.NewSource(int64(round)))
	r.Shuffle(len(set), func(i, j int) {
		set[i], set[j] = set[j], set[i]
	})

	start := clientID % int64(count) * int64(sub)
	return set[start : start+int64(sub)]
}

func calcEntropy(vals []string, fn func([]string, int) []string) {

	r := make(map[interface{}]int)
	for i := 0; i < *rounds; i++ {
		subs := fn(vals, *subset)
		for _, sub := range subs {
			r[sub]++
		}
	}

}
