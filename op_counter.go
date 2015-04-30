package crdt

type OpCounter struct {
  processIdentity string
  value           int64
  
  remoteReplicas  []chan<- CounterOperation
}

type CounterOperation interface {
  Perform(*OpCounter)
}
type IncrementOperation struct {
  value int64
}
type DecrementOperation struct {
  value int64
}
func IncrementByOne() *IncrementOperation {
  return &IncrementOperation{1}
}
func (op *IncrementOperation) Perform(counter *OpCounter) {
  counter.inc(op.value)
}
func (c *OpCounter) inc(x int64) {
  c.value += x
}

func (op *DecrementOperation) Perform(counter *OpCounter) {
  counter.dec(op.value)
}

func DecrementByOne() *DecrementOperation {
  return &DecrementOperation{1}
}
func (c *OpCounter) dec(x int64) {
  c.value -= x
}

func NewOpCounter(processIdentity string) *OpCounter {
  replicas := make([]chan<- CounterOperation, 0)
  return &OpCounter{processIdentity: processIdentity, remoteReplicas: replicas}
}

func (c *OpCounter) Value() int64 {
  return c.value
}

func (c *OpCounter) Increment() {
  op := IncrementByOne()
  op.Perform(c)
  c.replicate(op)
}

func (c *OpCounter) Decrement() {
  op := DecrementByOne()
  op.Perform(c)
  c.replicate(op)
}

func (c *OpCounter) replicate(op CounterOperation) {
  for _, replica := range c.remoteReplicas {
    replica <- op
  }
}

func (c *OpCounter) Listen(operations <-chan CounterOperation) {
  go func() {
    for op := range operations {
      op.Perform(c)
    }
  }()
}

func (c *OpCounter) AddReplica(replica chan<- CounterOperation) {
  c.remoteReplicas = append(c.remoteReplicas, replica)
}