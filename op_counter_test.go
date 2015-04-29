package crdt

import (
  "testing"
)

func TestOpCounter(t *testing.T) {
  c := NewOpCounter("node1")
  if c.Value() != 0 {
    t.Error("should have initial value of 0")
  }
  
  c.Increment()
  if c.Value() != 1 {
    t.Error("should have incremented")
  }
  
  c.Decrement()
  if c.Value() != 0 {
    t.Error("should have decremented")
  }
}

func TestOpCounterReplica(t *testing.T) {
  c1 := NewOpCounter("node1")
  c2 := NewOpCounter("node2")
  
  c1.AddReplica(c2)
  c2.AddReplica(c1)
  
  c1.Increment()
  if c2.Value() != 1 {
    t.Error("should have replicated increment op")
  }
}