package crdt

import (
	"testing"
	"time"
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
	c3 := NewOpCounter("node3")

	node2ReplicaCh := make(chan CounterOperation)
	c1.AddReplica(node2ReplicaCh)
	c2.Listen(node2ReplicaCh)

	node3ReplicaCh := make(chan CounterOperation)
	c1.AddReplica(node3ReplicaCh)
	c3.Listen(node3ReplicaCh)

	c1.Increment()

	// hacky for now but the replicas read using go channels
	// so have to wait a little for read to be processed.
	<-time.After(time.Second * 1)

	if c2.Value() != 1 {
		t.Error("should have incremented node2 replica")
	}
	if c3.Value() != 1 {
		t.Error("should have incremented node3 replica")
	}
}
