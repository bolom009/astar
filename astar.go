// Package astar implements the A* search algorithm for finding least-cost paths.
package astar

import (
	"github.com/bolom009/astar/heap"
	"github.com/bolom009/astar/intmap"
)

const (
	predictedCapacity = 64
)

type Graph[Node any] interface {
	Neighbours(n Node) []Node
}

type CostFunc[T comparable] func(a, b T) float32
type HasherFunc[T comparable] func(v T) int64

func reconstructPath[T comparable](paths map[T]T, current, start T) []T {
	path := make([]T, 0, 20)
	for current != start {
		path = append(path, current)
		current = paths[current]
	}

	path = append(path, start)
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
	return path
}

func FindPath[T comparable](g Graph[T], start, dest T, hashFn HasherFunc[T], d, h CostFunc[T]) []T {
	open := &priorityQueue[T]{}
	heap.Init(open)
	heap.Push(open, nodeItem[T]{
		item:     start,
		priority: h(start, dest),
	})

	gScore := intmap.New[int64, float32](predictedCapacity)
	gScore.Put(hashFn(start), 0)

	paths := make(map[T]T, predictedCapacity)
	for open.Len() > 0 {
		current := heap.Pop(open).item
		if current == dest {
			return reconstructPath[T](paths, current, start)
		}

		ck := hashFn(current)
		curScore, _ := gScore.Get(ck)
		for _, nb := range g.Neighbours(current) {
			nbk := hashFn(nb)
			tent := curScore + d(current, nb)
			if gs, ok := gScore.Get(nbk); !ok || tent < gs {
				gScore.Put(nbk, tent)
				paths[nb] = current
				f := tent + h(nb, dest)

				heap.Push(open, nodeItem[T]{item: nb, priority: f})
			}
		}
	}

	return nil
}
