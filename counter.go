package crdt

import (
	"sync"
)

type Counter struct {
	Process string

	IncAtoms map[string]int64
	DecAtoms map[string]int64

	lock *sync.RWMutex
}

// retrieves the current value of the distributed counter
func (c *Counter) Value() int64 {
	c.lock.RLock()
	defer c.lock.RUnlock()

	x := int64(0)
	for _, value := range c.IncAtoms {
		x = x + value
	}
	for _, value := range c.DecAtoms {
		x = x - value
	}
	return x
}

// called when we want to merge our counter with the state of the same
// counter from another node.
func (c *Counter) Merge(other *Counter) {
	c.lock.Lock()
	defer c.lock.Unlock()

	for key, value := range other.IncAtoms {
		ourValue := c.IncAtoms[key]
		if value > ourValue {
			c.IncAtoms[key] = value
		}
	}

	for key, value := range other.DecAtoms {
		ourValue := c.DecAtoms[key]
		if value > ourValue {
			c.DecAtoms[key] = value
		}
	}
}

func (c *Counter) Increment() {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.IncAtoms[c.Process] += 1
}

func (c *Counter) Decrement() {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.DecAtoms[c.Process] += 1
}

// creates a new counter, each node that has a replica of the counter
// state must have a unique identity
func NewCounter(process string) *Counter {
	return &Counter{
		Process:  process,
		IncAtoms: make(map[string]int64),
		DecAtoms: make(map[string]int64),
		lock:     &sync.RWMutex{},
	}
}
