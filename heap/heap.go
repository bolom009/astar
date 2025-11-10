package heap

import "sort"

type Interface[T comparable] interface {
	sort.Interface
	Push(x T)
	Pop() T
}

func Init[T comparable](h Interface[T]) {
	n := h.Len()
	for i := n/2 - 1; i >= 0; i-- {
		down[T](h, i, n)
	}
}

func Push[T comparable](h Interface[T], x T) {
	h.Push(x)
	up(h, h.Len()-1)
}

func Pop[T comparable](h Interface[T]) T {
	n := h.Len() - 1
	h.Swap(0, n)
	down(h, 0, n)
	return h.Pop()
}

func up[T comparable](h Interface[T], j int) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || !h.Less(j, i) {
			break
		}
		h.Swap(i, j)
		j = i
	}
}

func down[T comparable](h Interface[T], i0, n int) bool {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 {
			break
		}
		j := j1
		if j2 := j1 + 1; j2 < n && h.Less(j2, j1) {
			j = j2
		}
		if !h.Less(j, i) {
			break
		}
		h.Swap(i, j)
		i = j
	}
	return i > i0
}
