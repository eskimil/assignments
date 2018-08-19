// +build !solution

// Leave an empty line above this comment.
package lab5

type CondQueue struct {
	queue *FIFOQueue
	// TODO(student): Add needed field(s)
}

func NewCondQueue(size int) *CondQueue {
	return &CondQueue{}
}

func (cq *CondQueue) Enqueue(value interface{}) {

}

func (cq *CondQueue) Dequeue() interface{} {
	return nil
}

func (cq *CondQueue) Flush() {

}

func (cq *CondQueue) Empty() bool {
	return true
}

func (cq *CondQueue) Len() int {
	return -1
}
