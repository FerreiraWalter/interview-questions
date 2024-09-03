package main

import (
	"container/heap"
	"fmt"
)

type Location struct {
	x, y int
}

func (l Location) distFromUser(user Location) int {
	return (l.x-user.x)*(l.x-user.x) + (l.y-user.y)*(l.y-user.y)
}

type MaxHeap []Location

func (h MaxHeap) Len() int { return len(h) }
func (h MaxHeap) Less(i, j int) bool {
	return h[i].distFromUser(Location{0, 0}) > h[j].distFromUser(Location{0, 0})
}
func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h MaxHeap) Top() Location { return h[0] }
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(Location))
}
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (h MaxHeap) Empty() bool { return len(h) == 0 }

func newHeap() *MaxHeap {
	max := &MaxHeap{}
	heap.Init(max)
	return max
}

func findClosestDrivers(user Location, locations []Location, k int) []Location {
	maxHeap := newHeap()
	var res []Location

	for i := 0; i < k; i++ {
		heap.Push(maxHeap, locations[i])
	}

	for i := k; i < len(locations); i++ {
		top := maxHeap.Top()
		if locations[i].distFromUser(user) < top.distFromUser(user) {
			maxHeap.Pop()
			heap.Push(maxHeap, locations[i])
		}
	}

	for !maxHeap.Empty() {
		res = append(res, maxHeap.Top())
		heap.Pop(maxHeap)
	}

	return res
}

func main() {
	user := Location{5, 6}
	locations := []Location{
		{1, 3},
		{3, 4},
		{2, -1},
		{5, 6},
		{-2, -3},
	}
	result := findClosestDrivers(user, locations, 2)
	fmt.Printf("drivers locations closest to the user: ")
	for _, p := range result {
		fmt.Printf("[%v, %v] ", p.x, p.y)
	}
}
