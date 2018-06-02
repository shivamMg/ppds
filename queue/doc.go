/*
Package queue provides implementation to get and print formatted queue.

Example:

	import "github.com/shivammg/ppds/queue"

	// a queue implementation.
	type myQueue struct {
		elems []int
	}

	func (q *myQueue) push(ele int) {
		q.elems = append([]int{ele}, q.elems...)
	}

	func (q *myQueue) pop() (int, bool) {
		if len(q.elems) == 0 {
			return 0, false
		}
		e := q.elems[len(q.elems)-1]
		q.elems = q.elems[:len(q.elems)-1]
		return e, true
	}

	// myQueue implements queue.Queue. Notice that the receiver is of *myQueue
	// type - since Push and Pop are required to modify q.
	func (q *myQueue) Push(ele interface{}) {
		q.push(ele.(int))
	}

	func (q *myQueue) Pop() (interface{}, bool) {
		return q.pop()
	}

	// q := myQueue{}
	// q.push(11)
	// q.push(12)
	// queue.Print(&q) or queue.Sprint(&q)

*/
package queue
