package pubsub

import "sync"

type Message struct {
	Data string
}

type PubSub struct {
	subs []chan Message
	mu   sync.Mutex
}

func (ps *PubSub) Subscribe() chan Message {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	ch := make(chan Message, 1)
	ps.subs = append(ps.subs, ch)
	return ch
}

func (ps *PubSub) Publish(msg *Message) {
	ps.mu.Lock()
	defer ps.mu.Unlock()
	for _, sub := range ps.subs {
		sub <- *msg
	}
}
