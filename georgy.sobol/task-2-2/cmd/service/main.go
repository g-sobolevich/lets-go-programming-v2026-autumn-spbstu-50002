package main

import (
	"container/heap"
	"fmt"
)

type DishesHeap []int

func (h DishesHeap) Len() int           { return len(h) }
func (h DishesHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h DishesHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *DishesHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *DishesHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	var (
		n      int
		k      int
		result int
	)
	dishesHeap := &DishesHeap{}
	heap.Init(dishesHeap)
	if _, err := fmt.Scan(&n); err != nil || (n < 1 || n > 10000) {
		fmt.Println("Invalid dishes number")
		return
	}

	for range n {
		var dish int
		if _, err := fmt.Scan(&dish); err != nil || (n < -10000 || n > 10000) {
			fmt.Println("Invalid dish")
			return
		}
		heap.Push(dishesHeap, dish)
	}

	if _, err := fmt.Scan(&k); err != nil || (n < 1 || n > 10000) {
		fmt.Println("Invalid priority number")
		return
	}

	for range k {
		result = heap.Pop(dishesHeap).(int)
	}

	fmt.Println(result)
}
