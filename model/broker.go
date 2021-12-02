package model

type IBroker interface {
	action(broker *Broker)
}

/**
 * NewBroker
 * @Description:
 * @return *Broker
 */
func NewBroker() *Broker {
	return &Broker{obs: make([]IBroker, 0)}
}

/**
 * notify
 * @Description:
 * @receiver b
 */
func (b *Broker) notify() {
	for _, o := range b.obs {
		o.action(b)
	}
}

/**
 * attach
 * @Description:
 * @receiver b
 * @param broker
 */
func (b *Broker) attach(broker IBroker) {
	b.obs = append(b.obs, broker)
}

/**
 * registry
 * @Description:
 * @receiver b
 * @param c
 */
func (b *Broker) registry(c chan interface{}) {
	b.call = c
	b.notify()
}
