package main

import (
	"fmt"
	"time"
	"github.com/asynkron/protoactor-go/actor"
)
//defining request msg 
type RequestMsg struct {
	Content string
}
//defning response 
type AcknowledgmentMsg struct {
	Content string
} 

type Actor_A struct{}

type Actor_B struct{}

// Handle RequestMsg and send AcknowledgmentMsg
func (state *Actor_B) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *RequestMsg:
		fmt.Println("Actor B received the message...=>", msg.Content)
		
	}
}


// Handle AcknowledgmentMsg

func (state *Actor_A) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *AcknowledgmentMsg:
		fmt.Println("Actor A received acknowledgment...=>", msg.Content)
	}
}


func main() {


	fmt.Println("Simple Actor Model Communication using Go Language and ProtoActor")






	system := actor.NewActorSystem()

	// Define and spawn Actor_A
	propsA := actor.PropsFromProducer(func() actor.Actor { return &Actor_A{} })
	system.Root.Spawn(propsA)

	// Define and spawn Actor_B
	propsB := actor.PropsFromProducer(func() actor.Actor { return &Actor_B{} })
	pidB := system.Root.Spawn(propsB)

	// Send the RequestMsg to Actor_B
	fmt.Println("Actor A sending message to Actor B...")
	system.Root.Send(pidB, &RequestMsg{Content: "Hello from Actor A!"})

	// Allow time for message exchange
	time.Sleep(1 * time.Second)
}
