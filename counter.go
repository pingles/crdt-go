package crdt

type Counter struct {
	processIdentity string

	incValues map[string]int64
	decValues map[string]int64
}

// retrieves the current value of the distributed counter
func (c *Counter) Value() int64 {
	x := int64(0)
	for _, value := range c.incValues {
		x = x + value
	}
	for _, value := range c.decValues {
		x = x - value
	}
	return x
}

// called when we want to merge our counter with the state of the same
// counter from another node.
func (c *Counter) Merge(other *Counter) {
	for key, value := range other.incValues {
		ourValue := c.incValues[key]
		if value > ourValue {
			c.incValues[key] = value
		}
	}

	for key, value := range other.decValues {
		ourValue := c.decValues[key]
		if value > ourValue {
			c.decValues[key] = value
		}
	}
}

func (c *Counter) Increment() {
	c.incValues[c.processIdentity] += 1
}

func (c *Counter) Decrement() {
	c.decValues[c.processIdentity] += 1
}

// creates a new counter, each node that has a replica of the counter
// state must have a unique identity
func NewCounter(processIdentity string) *Counter {
	return &Counter{processIdentity, make(map[string]int64), make(map[string]int64)}
}
