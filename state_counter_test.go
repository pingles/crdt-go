package crdt

import (
	"testing"
)

func TestStateCounterIncrement(t *testing.T) {
	c := NewStateCounter("node1")
	if c.Value() != 0 {
		t.Error("should have 0 state, was", c.Value())
	}
	c.Increment()
	if c.Value() != 1 {
		t.Error("should have incremented to 1, was:", c.Value())
	}
}

func TestStateCounterDecrement(t *testing.T) {
	c1 := NewStateCounter("node1")
	c1.Increment()
	c1.Decrement()

	if c1.Value() != 0 {
		t.Error("should have 0 value, was", c1.Value())
	}

	c2 := NewStateCounter("node2")
	c2.Increment()
	c2.Increment()
	c2.Decrement()

	c1.Merge(c2)
	if c1.Value() != 1 {
		t.Error("should have converged to value of 1, was", c1.Value())
	}
}

func TestStateCounterMerge(t *testing.T) {
	c1 := NewStateCounter("node1")
	c2 := NewStateCounter("node2")

	c1.Increment()
	c2.Increment()

	c1.Merge(c2)

	if c1.Value() != 2 {
		t.Error("should have merged two incs, value was", c1.Value())
	}
}
