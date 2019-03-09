package queue

import (
	"github.com/surathj/q-rex/builder"
)

type SimpleQueueBuilder struct{
	queue QueueObject
}

func (s *SimpleQueueBuilder) SetSize(size int) builder.QueueBuildProcess {
	s.queue.Size = size
	s.queue.Items = make([]interface{}, size)
	s.queue.Front = 0
	s.queue.FreePos = 0
	s.queue.Rear = size - 1
	return s
}

func (s *SimpleQueueBuilder) SetType(qtype builder.QueueType) builder.QueueBuildProcess {
	s.queue.Type = builder.Simple
	return s
}

func (s *SimpleQueueBuilder) Build() QueueObject {
	return s.queue
}

