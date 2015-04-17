package container

type Queue struct {
	left, right Stack
}

func (q *Queue) Size() int {
	return q.left.Size() + q.right.Size()
}

func (q *Queue) Push(v interface{}) {
	q.right.Push(v)
}

func (q *Queue) Pop() (interface{}, bool) {
	if q.left.Size() == 0 {
		for {
			v, ok := q.right.Pop()
			if ok {
				q.left.Push(v)
			} else {
				break
			}
		}
	}
	return q.left.Pop()
}
