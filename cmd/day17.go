//go:build ignore

package main

import (
	"AOC/lib"
	"container/heap"
	"fmt"
)

func main() {
	lib.AssertArgs()
	g := grid(lib.AssertInputSingleDigitGrid())

	fmt.Println(g.FindPath(lib.Point{Y: len(g) - 1, X: len(g[0]) - 1}, 1, 3))
	fmt.Println(g.FindPath(lib.Point{Y: len(g) - 1, X: len(g[0]) - 1}, 4, 10))
}

func (m grid) FindPath(end lib.Point, minStraight, maxStraight int8) int {
	pq := priorityQueue{}

	// for part 1, doesn't affect part 2
	heap.Push(&pq, &pqItem{step{
		Point:         lib.Point{X: 0, Y: 0},
		Direction:     lib.RIGHT,
		StraightCount: 1,
	}, 0})

	// for part 2
	heap.Push(&pq, &pqItem{step{
		Point:         lib.Point{X: 0, Y: 0},
		Direction:     lib.DOWN,
		StraightCount: 1,
	}, 0})

	bestTotalHeatLoss := make([][][][4]int, len(m))
	for y := range m {
		bestTotalHeatLoss[y] = make([][][4]int, len(m[0]))
		for x := range m[0] {
			bestTotalHeatLoss[y][x] = make([][4]int, maxStraight)
		}
	}

	for pq.Len() > 0 {
		queued, _ := heap.Pop(&pq).(*pqItem)
		step := queued.Step
		point := step.Point
		if point.Y < 0 || point.Y > end.Y || point.X < 0 || point.X > end.X {
			continue
		}

		currentHeatLoss := int(m[point.Y][point.X]) + queued.TotalHeatLoss
		if point == end {
			if step.StraightCount < minStraight {
				continue
			}
			return currentHeatLoss - int(m[0][0])
		}

		if bestTotalHeatLoss[point.Y][point.X][step.StraightCount-1][step.Direction] != 0 && bestTotalHeatLoss[point.Y][point.X][step.StraightCount-1][step.Direction] <= currentHeatLoss {
			continue
		}
		bestTotalHeatLoss[point.Y][point.X][step.StraightCount-1][step.Direction] = currentHeatLoss

		if step.StraightCount < maxStraight {
			step.addStraight(&pq, currentHeatLoss)
		}

		if step.StraightCount >= minStraight {
			step.addTurns(&pq, currentHeatLoss)
		}
	}

	return -1
}

func (s step) addStraight(pq *priorityQueue, currentHeatLoss int) {
	switch s.Direction {
	case lib.UP:
		heap.Push(pq, &pqItem{step{Point: lib.Point{X: s.Point.X, Y: s.Point.Y - 1}, Direction: lib.UP, StraightCount: s.StraightCount + 1}, currentHeatLoss})
	case lib.DOWN:
		heap.Push(pq, &pqItem{step{Point: lib.Point{X: s.Point.X, Y: s.Point.Y + 1}, Direction: lib.DOWN, StraightCount: s.StraightCount + 1}, currentHeatLoss})
	case lib.LEFT:
		heap.Push(pq, &pqItem{step{Point: lib.Point{X: s.Point.X - 1, Y: s.Point.Y}, Direction: lib.LEFT, StraightCount: s.StraightCount + 1}, currentHeatLoss})
	case lib.RIGHT:
		heap.Push(pq, &pqItem{step{Point: lib.Point{X: s.Point.X + 1, Y: s.Point.Y}, Direction: lib.RIGHT, StraightCount: s.StraightCount + 1}, currentHeatLoss})
	}
}

func (s step) addTurns(pq *priorityQueue, currentHeatLoss int) {
	switch s.Direction {
	case lib.UP, lib.DOWN:
		heap.Push(pq, &pqItem{step{Point: lib.Point{X: s.Point.X - 1, Y: s.Point.Y}, Direction: lib.LEFT, StraightCount: 1}, currentHeatLoss})
		heap.Push(pq, &pqItem{step{Point: lib.Point{X: s.Point.X + 1, Y: s.Point.Y}, Direction: lib.RIGHT, StraightCount: 1}, currentHeatLoss})
	case lib.LEFT, lib.RIGHT:
		heap.Push(pq, &pqItem{step{Point: lib.Point{X: s.Point.X, Y: s.Point.Y - 1}, Direction: lib.UP, StraightCount: 1}, currentHeatLoss})
		heap.Push(pq, &pqItem{step{Point: lib.Point{X: s.Point.X, Y: s.Point.Y + 1}, Direction: lib.DOWN, StraightCount: 1}, currentHeatLoss})
	}
}

// satisfies heap.Interface and holds pqItems... see 2021 day 15
// do not use pq.Push/Pop directly -- use heap.Push/Pop instead
func (pq priorityQueue) Len() int { return len(pq) }

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].TotalHeatLoss < pq[j].TotalHeatLoss
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *priorityQueue) Push(x interface{}) {
	item := x.(*pqItem)
	*pq = append(*pq, item)
}

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}

type priorityQueue []*pqItem

type grid [][]int

type step struct {
	Point         lib.Point
	Direction     lib.OrthogonalDirection
	StraightCount int8
}

type pqItem struct {
	Step          step
	TotalHeatLoss int
}
