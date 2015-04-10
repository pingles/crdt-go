package crdt

type Counter struct {
  nodeIdentity  string
  counterValues map[string]int64
}

// retrieves the current value of the distributed counter
func (c *Counter) Value() int64 {
  x := int64(0)
  for _, value := range c.counterValues {
    x = x + value
  }
  return x
}

// called when we want to merge our counter with the state of the same
// counter from another node.
func (c *Counter) Merge(other *Counter) {
  for key, value := range other.counterValues {
    ourValue := c.counterValues[key]
    if value > ourValue {
      c.counterValues[key] = value
    }
  }
}

func (c *Counter) Increment() {
  c.counterValues[c.nodeIdentity] += 1
}

// creates a new counter, each node that has a replica of the counter
// state must have a unique identity
func NewCounter(nodeIdentity string) *Counter {
  return &Counter{nodeIdentity, make(map[string]int64)}
}