package main

import (
	"GoPractice/ProblemSets/Asynchronous/ProcessStreamInOrder/queue"
	"container/heap"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Message struct {
	Value          string
	SequenceNumber int
}

type StreamConsumer struct {

	// represents "where" we are processing the stream
	streamIndex int

	// priority queue acts as a buffer to keep track of messages from stream
	// that been consumed out of order from what our current streamIndex is expecting
	buffer interface {
		heap.Interface
		Peek() *queue.Item
	}

	// lock critical section of priority queue and streamIndex
	mutex *sync.Mutex
}

func (s *StreamConsumer) Process(msg Message) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if msg.SequenceNumber == s.streamIndex {
		fmt.Print(msg.Value, " ")
		s.streamIndex++
		for {
			head := s.buffer.Peek()
			if head != nil && head.GetPriority() == s.streamIndex {
				fmt.Print(head.GetValue(), " ")
				heap.Pop(s.buffer)
				s.streamIndex++
			} else {
				return
			}
		}
	} else {
		heap.Push(s.buffer, queue.NewItem(msg.Value, msg.SequenceNumber))
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
			buffer:      &queue.PriorityQueue{},
			mutex:       &sync.Mutex{},
		}

		wg = &sync.WaitGroup{}
	)

	// expect all messages to be processed
	wg.Add(len(messages))

	// randomize the message stream
	rand.Seed(time.Now().Unix())
	rand.Shuffle(len(messages), func(i, j int) {
		messages[i], messages[j] = messages[j], messages[i]
	})

	// process message stream in a separate go routine each
	// adds more randomness with respect to order
	for _, val := range messages {
		msg := val
		go func() {
			defer wg.Done()
			s.Process(msg)
		}()
	}

	// block until all go routines return
	wg.Wait()
	fmt.Print("\n")
}
