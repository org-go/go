package test

import (
	"acs-sdk-go/threads"
	"testing"
	"time"
)

func TestSafe(t *testing.T) {
	threads.ServiceMux.SafeRun(func() {
		println(1, 2, 3)
	})
	time.Sleep(1e9)
}
