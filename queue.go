package bitradix

type node32[T any] struct {
	*Radix32[T]
	branch int // -1 root, 0 left branch, 1 right branch
}

type queue32[T any] []*node32[T]

type node64[T any] struct {
	*Radix64[T]
	branch int
}

type queue64[T any] []*node64[T]

// Push adds a node32 to the queue.
func (q *queue32[T]) Push(n *node32[T]) {
	*q = append(*q, n)
}

// Pop removes and returns a node from the queue in first to last order.
func (q *queue32[T]) Pop() *node32[T] {
	lq := len(*q)
	if lq == 0 {
		return nil
	}

	n := (*q)[0]
	switch lq {
	case 1:
		*q = (*q)[:0]
	default:
		*q = (*q)[1:lq]
	}

	return n
}

func (q *queue64[T]) Push(n *node64[T]) {
	*q = append(*q, n)
}

func (q *queue64[T]) Pop() *node64[T] {
	lq := len(*q)
	if lq == 0 {
		return nil
	}

	n := (*q)[0]
	switch lq {
	case 1:
		*q = (*q)[:0]
	default:
		*q = (*q)[1:lq]
	}

	return n
}
