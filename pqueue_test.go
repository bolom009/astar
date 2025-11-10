// Copyright 2013 Frederik Zipp. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package astar

import (
	"testing"

	"github.com/bolom009/astar/heap"
)

func TestPushPop(t *testing.T) {
	pq := &priorityQueue[string]{}
	heap.Init(pq)

	want := "cadbe"
	heap.Push(pq, nodeItem[string]{item: "a", priority: 1.2})
	heap.Push(pq, nodeItem[string]{item: "b", priority: 5})
	heap.Push(pq, nodeItem[string]{item: "c", priority: -0.4})
	heap.Push(pq, nodeItem[string]{item: "d", priority: 3.7})
	heap.Push(pq, nodeItem[string]{item: "e", priority: 11})

	s := ""
	for pq.Len() > 0 {
		s += heap.Pop(pq).item
	}

	if s != want {
		t.Errorf("Retrieved item order was %q, want %q.", s, want)
	}
}
