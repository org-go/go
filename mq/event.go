package mq

import (
	"fmt"
)

type iAction interface {
	OnLog([]event)
	OnAction([]event)
}
type event struct {
	Gid string `json:"gid"`
}

type action struct{}

func (d action) OnLog(events []event) {
	fmt.Printf("%s log.", events)
}

func (d action) OnAction(events []event) {
	fmt.Printf("%s action.", events)
}
