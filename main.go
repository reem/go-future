package future

import "sync"

type FutureState int

// The consumer hasn't asked for a value, and the producer hasn't given it.
const start FutureState = 0

// The producer has yielded a value, but the consumer hasn't asked for it.
const producerCompleted FutureState = 1

// The consumer has asked for a value, but the producer hasn't given it.
const consumerWaiting FutureState = 3

// The consumer asked for a value and received it.
const complete FutureState = 4

type Future struct {
	state *FutureState
	cond  *sync.Cond
}

type Producer struct {
	state *FutureState
	cond  *sync.Cond
}

func Pair() (producer Producer, consumer Future) {
	var mutex = &sync.Mutex{}
	var cond = sync.NewCond(mutex)
	var state = start

	producer = Producer{&state, cond}
	consumer = Future{&state, cond}
	return
}

func (p *Producer) Complete() {
	p.cond.L.Lock()
	defer p.cond.L.Unlock()

	switch *p.state {
	case start:
		*p.state = producerCompleted
	case consumerWaiting:
		*p.state = complete
		p.cond.Signal()
	default:
		panic("Completed a future more than once.")
	}
}

func (f *Future) Await() {
	f.cond.L.Lock()
	defer f.cond.L.Unlock()

	switch *f.state {
	case start:
		*f.state = consumerWaiting
		f.cond.Wait()
	case producerCompleted:
		*f.state = complete
	default:
		panic("Awaited a future more than once.")
	}
}
