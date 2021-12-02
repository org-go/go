package threads

import (
	"fmt"
	"testing"
)

func Test_mux_Run(t *testing.T) {}

func Test_mux_Wait(t *testing.T) {}

func Test_mux_SafeRun(t *testing.T) {

	ServiceMux.SafeRun(func() {
		fmt.Println(`----`)
	})

}
