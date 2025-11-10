package astar

type nodeItem[T any] struct {
	item     T
	priority float32
	index    int
}

type priorityQueue[T any] []nodeItem[T]

// Less - reverted ordering for reconstruct path
func (p priorityQueue[T]) Less(i, j int) bool {
	return p[i].priority < p[j].priority
}

func (p priorityQueue[T]) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
	p[i].index = i
	p[j].index = j
}

func (p priorityQueue[T]) Len() int {
	return len(p)
}

func (p *priorityQueue[T]) Push(it nodeItem[T]) {
	it.index = len(*p)
	*p = append(*p, it)
}

func (p *priorityQueue[T]) Pop() nodeItem[T] {
	old := *p
	n := len(old)
	it := old[n-1]
	it.index = -1
	*p = old[:n-1]
	return it
}
