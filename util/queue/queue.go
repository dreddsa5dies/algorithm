package queue

type (
	// Queue - очередь
	Queue struct {
		start, end *node
		length     int
	}
	node struct {
		value interface{}
		next  *node
	}
)

// New - Create a new queue
func New() *Queue {
	return &Queue{nil, nil, 0}
}

// Dequeue - Take the next item off the front of the queue
func (que *Queue) Dequeue() interface{} {
	if que.length == 0 {
		return nil
	}
	n := que.start
	if que.length == 1 {
		que.start = nil
		que.end = nil
	} else {
		que.start = que.start.next
	}
	que.length--
	return n.value
}

// Enqueue - Put an item on the end of a queue
func (que *Queue) Enqueue(value interface{}) {
	n := &node{value, nil}
	if que.length == 0 {
		que.start = n
		que.end = n
	} else {
		que.end.next = n
		que.end = n
	}
	que.length++
}

// Len - Return the number of items in the queue
func (que *Queue) Len() int {
	return que.length
}

// Peek - Return the first item in the queue without removing it
func (que *Queue) Peek() interface{} {
	if que.length == 0 {
		return nil
	}
	return que.start.value
}
