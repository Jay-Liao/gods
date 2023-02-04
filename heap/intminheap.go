package heap

import "container/heap"

type IntMinHeap []int

func (h IntMinHeap) Len() int {
	return len(h)
}

func (h IntMinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntMinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntMinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntMinHeap) Pop() any {
	// pop and remove the last item
	last := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return last
}

func NewIntMinHeap() *IntMinHeap {
	return &IntMinHeap{}
}

func (h *IntMinHeap) PushVal(x int) {
	heap.Push(h, x)
}

func (h *IntMinHeap) PopMin() int {
	return heap.Pop(h).(int)
}
