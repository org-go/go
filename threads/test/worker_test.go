package test

import (
	"acs-sdk-go/threads"
	"fmt"
	"testing"
)

func TestWorker(t *testing.T) {
	threads.NewWorker(func() {
		fmt.Println("11111")
	}, 3).Schedule()
}

func TestGroupWorker(t *testing.T) {

	threads.GWorker(func() {
		fmt.Println("1")
	}, func() {
		fmt.Println("2")
	})
}
