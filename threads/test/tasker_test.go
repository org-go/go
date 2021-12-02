package test

import (
	"acs-sdk-go/threads"
	"fmt"

	"testing"
)

func TestTasker(t *testing.T) {

	threads.NewTask(3).Schedule(func() {
		for i := 0; i < 4; i++ {
			fmt.Println(i)
		}
	})
}
