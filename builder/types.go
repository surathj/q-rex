package builder

type QueueType int

const (
	Simple QueueType = iota
	Circular
)