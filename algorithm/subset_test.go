package algorithm

import (
	"strconv"
	"testing"
)

func Test_calcEntropy(t *testing.T) {

	var vals []string
	for i := 0; i < *elements; i++ {
		vals = append(vals, strconv.Itoa(i))
	}

	calcEntropy(vals, complexSubset)

}
