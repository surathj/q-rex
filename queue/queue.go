package queue

import (
	"fmt"
	"github.com/surathj/q-rex/builder"
)

type Queue interface {
	Enqueue(interface{})
	Dequeue()
	GetSize() int
}

// Front stores the current leading index
// Rear Stores the current last occupied index
// FreePos stores the next available index
type QueueObject struct {
	Front int
	Rear int
	Size int
	FreePos int
	Type builder.QueueType
	Items []interface{}
}

// Queue functionality Implementation
func (q *QueueObject) Enqueue(item interface{}) {
	if q.FreePos == len(q.Items) {
		fmt.Println("Queue Full")
		return
	}
	temp := q.FreePos
	q.Items[q.FreePos] = item  
	q.FreePos = temp + 1
	fmt.Println("FreePos: ", q.FreePos)
}

func (q *QueueObject) Dequeue() {
	if q.Type == builder.Simple {
		temp := q.Items[1:]
		fmt.Println("Dq temp: ", temp)
		q.Items = make([]interface{}, len(temp)+1)
		for i := range temp {
			if temp[i] == nil {
				q.FreePos = i
				break
			}
			q.Items[i] = temp[i]
		}
		fmt.Println("q after dq: ", q.Items)
		fmt.Println("Dq FreePos: ", q.FreePos)
	}
	if q.Type == builder.Circular {
		// Not implemented
	}
}

func (q *QueueObject) GetSize() int {
	return q.Size
}