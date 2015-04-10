package crdt

import (
  "testing"
)

func TestCounterIncrement(t *testing.T) {
  c := NewCounter("node1")
  if c.Value() != 0 {
    t.Error("should have 0 state, was", c.Value())
  }
  c.Increment()
  if c.Value() != 1 {
    t.Error("should have incremented to 1, was:", c.Value())
  }
}

func TestCounterMerge(t *testing.T) {
  c1 := NewCounter("node1")
  c2 := NewCounter("node2")
  
  c1.Increment()
  c2.Increment()
  
  c1.Merge(c2)
  
  if c1.Value() != 2 {
    t.Error("should have merged two incs, value was", c1.Value())
  }
}