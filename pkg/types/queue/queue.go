package queue

type Queue struct {
	First  *Node
	Last   *Node
	Length int
}

// Enqueue puts element in the end of the queue
func (q *Queue) Enqueue(v interface{}) {
	defer func() {
		q.Length++
	}()

	newNode := &Node{Val: v}
	if q.First == nil {
		q.First = newNode
		q.Last = newNode
		return
	}
	q.Last.Next = newNode
	q.Last = newNode
}

// Dequeue returns first element of a queue
func (q *Queue) Dequeue() interface{} {
	defer func() {
		q.Length--
	}()

	if q.First == q.Last {
		q.First = nil
		q.Last = nil

		return q.First
	}

	holding := q.First
	q.First = q.First.Next

	return holding.Val
}

func (q *Queue) Unshift() {
	q.First = q.First.Next
	q.Length--
}

//Discord
//Udemy
//google
