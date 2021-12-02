package mq

import (
	"fmt"
	"reflect"
	"sync"
	"testing"
)

func Test_initialization(t *testing.T) {
	tests := []struct {
		name string
		want iAlignment
	}{
		{name: `eros`, want: initialization()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := initialization(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("initialization() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_multipleAlignment_DelTopic(t *testing.T) {
	type fields struct {
		action    iAction
		topics    map[string][]chan Message
		callbacks map[string][]CallbackFunc
		RWMutex   *sync.RWMutex
	}
	type args struct {
		topic string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := multipleAlignment{
				action:    tt.fields.action,
				topics:    tt.fields.topics,
				callbacks: tt.fields.callbacks,
				RWMutex:   tt.fields.RWMutex,
			}
			fmt.Println(i)
		})
	}
}

func Test_multipleAlignment_Publish(t *testing.T) {
	alignment := initialization()
	subscribe := alignment.Subscribe(`eros`, 10)
	fmt.Println(subscribe)
}

func Test_multipleAlignment_Subscribe(t *testing.T) {
	type fields struct {
		action    iAction
		topics    map[string][]chan Message
		callbacks map[string][]CallbackFunc
		RWMutex   *sync.RWMutex
	}
	type args struct {
		topic  string
		buffer int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   <-chan Message
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := multipleAlignment{
				action:    tt.fields.action,
				topics:    tt.fields.topics,
				callbacks: tt.fields.callbacks,
				RWMutex:   tt.fields.RWMutex,
			}
			if got := i.Subscribe(tt.args.topic, tt.args.buffer); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Subscribe() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_multipleAlignment_SubscribeCallback(t *testing.T) {
	type fields struct {
		action    iAction
		topics    map[string][]chan Message
		callbacks map[string][]CallbackFunc
		RWMutex   *sync.RWMutex
	}
	type args struct {
		topic    string
		callback CallbackFunc
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := multipleAlignment{
				action:    tt.fields.action,
				topics:    tt.fields.topics,
				callbacks: tt.fields.callbacks,
				RWMutex:   tt.fields.RWMutex,
			}
			fmt.Println(i)
		})
	}
}

func Test_multipleAlignment_Unsubscribe(t *testing.T) {
	type fields struct {
		action    iAction
		topics    map[string][]chan Message
		callbacks map[string][]CallbackFunc
		RWMutex   *sync.RWMutex
	}
	type args struct {
		topic string
		msg   <-chan Message
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := multipleAlignment{
				action:    tt.fields.action,
				topics:    tt.fields.topics,
				callbacks: tt.fields.callbacks,
				RWMutex:   tt.fields.RWMutex,
			}
			fmt.Println(i)
		})
	}
}

func Test_multipleAlignment__subscribe(t *testing.T) {
	type fields struct {
		action    iAction
		topics    map[string][]chan Message
		callbacks map[string][]CallbackFunc
		RWMutex   *sync.RWMutex
	}
	type args struct {
		subs []chan Message
		msg  Message
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := multipleAlignment{
				action:    tt.fields.action,
				topics:    tt.fields.topics,
				callbacks: tt.fields.callbacks,
				RWMutex:   tt.fields.RWMutex,
			}
			fmt.Println(i)
		})
	}
}

func Test_multipleAlignment_ariaNotify(t *testing.T) {
	type fields struct {
		action    iAction
		topics    map[string][]chan Message
		callbacks map[string][]CallbackFunc
		RWMutex   *sync.RWMutex
	}
	type args struct {
		events []event
		status int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &multipleAlignment{
				action:    tt.fields.action,
				topics:    tt.fields.topics,
				callbacks: tt.fields.callbacks,
				RWMutex:   tt.fields.RWMutex,
			}
			fmt.Println(i)
		})
	}
}
