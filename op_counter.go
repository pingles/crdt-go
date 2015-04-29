package crdt

type OpCounter struct {
  processIdentity string
  value           int64
  
  remoteReplicas  []*OpCounter
}

func NewOpCounter(processIdentity string) *OpCounter {
  return &OpCounter{processIdentity: processIdentity, remoteReplicas: make([]*OpCounter, 0)}
}

func (c *OpCounter) Value() int64 {
  return c.value
}

func (c *OpCounter) downstreamIncrement() {
  c.value += 1
}

func (c *OpCounter) downstreamDecrement() {
  c.value -= 1
}

func (c *OpCounter) replicas() []*OpCounter {
  allReplicas := make([]*OpCounter, len(c.remoteReplicas) + 1)
  allReplicas[0] = c
  for idx, r := range c.remoteReplicas {
    allReplicas[idx+1] = r
  }
  return allReplicas
}

func (c *OpCounter) Increment() {
  for _, r := range c.replicas() {
    r.downstreamIncrement()
  }
}

func (c *OpCounter) Decrement() {
  for _, r := range c.replicas() {
    r.downstreamDecrement()
  }
}

func (local *OpCounter) AddReplica(remote *OpCounter) {
  local.remoteReplicas = append(local.remoteReplicas, remote)
}