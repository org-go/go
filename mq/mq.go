package mq

import (
	"encoding/gob"
	"strconv"
	"sync"
	"time"
)

type Message struct {
	Trigger string
	Event   string
	Content interface{}
}

type CallbackFunc func(message Message)

type iAlignment interface {
	Publish(topic string, msg Message)
	Subscribe(topic string, buffer int) <-chan Message
	SubscribeCallback(topic string, callback CallbackFunc)
	Unsubscribe(topic string, msg <-chan Message)
	DelTopic(topic string)
}

var Alignment iAlignment

func initialization() iAlignment {
	return &multipleAlignment{
		topics:    make(map[string][]chan Message),
		callbacks: make(map[string][]CallbackFunc),
	}
}

func init() {
	gob.Register(Message{})
	gob.Register(event{})
}

type multipleAlignment struct {
	action    iAction
	topics    map[string][]chan Message
	callbacks map[string][]CallbackFunc
	*sync.RWMutex
}

func (i multipleAlignment) Publish(topic string, msg Message) {
	i.Lock()
	subs, ok := i.topics[topic]
	callbacks, ack := i.callbacks[topic]
	i.Unlock()
	if ok {
		go func(subs []chan Message) { i._subscribe(subs, msg) }(subs)
	}
	if ack {
		for i := 0; i < len(callbacks); i++ {
			go callbacks[i](msg)
		}
	}
}

func (i multipleAlignment) _subscribe(subs []chan Message, msg Message) {
	for i := 0; i < len(subs); i++ {
		select {
		case subs[i] <- msg:
		case <-time.After(1e9):
		}
	}
}

func (i multipleAlignment) Subscribe(topic string, buffer int) <-chan Message {
	ch := make(chan Message, buffer)
	i.Lock()
	i.topics[topic] = append(i.topics[topic], ch)
	i.Unlock()
	return ch
}

func (i multipleAlignment) SubscribeCallback(topic string, callback CallbackFunc) {
	i.Lock()
	i.callbacks[topic] = append(i.callbacks[topic], callback)
	i.Unlock()
}

func (i multipleAlignment) DelTopic(topic string) {
	i.Lock()
	defer i.Unlock()
	if _, ok := i.topics[topic]; ok {
		delete(i.topics, topic)
	}
}

func (i multipleAlignment) Unsubscribe(topic string, msg <-chan Message) {
	i.Lock()
	defer i.Unlock()
	if subs, ok := i.topics[topic]; ok {
		var news []chan Message
		for _, sub := range subs {
			if sub != msg {
				news = append(news, sub)
			}
		}
		i.topics[topic] = news
	}

}

func (i *multipleAlignment) ariaNotify(events []event, status int) {
	for _, event := range events {
		i.Publish(event.Gid, Message{
			Trigger: event.Gid,
			Event:   strconv.FormatInt(int64(status), 10),
			Content: events,
		})
	}
}
