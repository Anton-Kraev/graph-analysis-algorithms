package AStar

import (
	pq "gopkg.in/dnaeon/go-priorityqueue.v1"
)

func AStar[T Number](s, t int, adj [][]Edge[T], h func(v, t int) T) T {
	maxT := MaxValue(T(0))
	dist := make([]T, len(adj))
	for i := 0; i < len(dist); i++ {
		dist[i] = maxT
	}
	dist[s] = 0

	q := pq.New[int, T](pq.MinHeap)
	q.Put(s, h(s, t))

	for !q.IsEmpty() {
		item := q.Get()
		v, d := item.Value, item.Priority

		if v == t {
			return d
		}
		if d != dist[v]+h(v, t) {
			continue
		}

		for _, e := range adj[v] {
			if dist[e.to] > dist[v]+e.c {
				dist[e.to] = dist[v] + e.c
				q.Put(e.to, dist[e.to]+h(e.to, t))
			}
		}
	}

	return maxT
}
