package main

import (
	"GoPractice/Concurrency/ProcessStreamInOrder/queue"
	"container/heap"
	"fmt"
	"math/rand"
)

type Message struct {
	Value          string
	SequenceNumber int
}

type StreamConsumer struct {
	streamIndex int
	pq          interface {
		heap.Interface
		Peek() *queue.Item
	}
}

func (s *StreamConsumer) Process(msg Message) {
	// if currently submitted msg matches our expected place in the stream
	// we process it, and increment the index
	if msg.SequenceNumber == s.streamIndex {
		fmt.Print(msg.Value, " ")
		s.streamIndex++
		for {
			head := s.pq.Peek()
			if head != nil && head.GetPriority() == s.streamIndex {
				fmt.Print(head.GetValue(), " ")
				heap.Pop(s.pq)
				s.streamIndex++
			} else {
				return
			}
		}
	} else {
		heap.Push(s.pq, queue.NewItem(msg.Value, msg.SequenceNumber))
	}
}

func main() {
	var (
		messages = []Message{
			{"hello", 1},
			{"how", 2},
			{"are", 3},
			{"you", 4},
			{"doing", 5},
			{"today", 6},
			{"?", 7},
		}

		s = StreamConsumer{
			streamIndex: 1,
			pq:          &queue.PriorityQueue{},
		}
	)

	rand.Shuffle(len(messages), func(i, j int) {
		messages[i], messages[j] = messages[j], messages[i]
	})

	for _, val := range messages {
		s.Process(val)
	}
	fmt.Print("\n")
}
