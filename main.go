package main

import (
	"fmt"
	"github.com/surathj/q-rex/queue"
	"github.com/surathj/q-rex/builder"
	
)

func main(){
	builderComplex := &builder.BuilderDirector{}

	simpleQueueBuilder := &queue.SimpleQueueBuilder{} 
	builderComplex.SetQueueBuilder(simpleQueueBuilder)
	builderComplex.Construct()

	simpleQueue := simpleQueueBuilder.Build()
	simpleQueue.Enqueue(1)
	simpleQueue.Enqueue("a")
	simpleQueue.Enqueue(5)
	simpleQueue.Enqueue(90)
	simpleQueue.Enqueue("h")
	fmt.Println("items", simpleQueue.Items)
	simpleQueue.Dequeue()
	simpleQueue.Dequeue()
	simpleQueue.Enqueue("67")
	simpleQueue.Enqueue("69")
	fmt.Println("items", simpleQueue.Items)
}