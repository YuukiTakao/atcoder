package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
	"strconv"
)

type Item struct {
	value    int
	priority int

	index int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}
func (pq *PriorityQueue) update(item *Item, value int, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}
func maxOf(vars ...int) int {
	max := int(math.Pow10(17)) * -1
	for _, v := range vars {
		if max < v {
			max = v
		}
	}
	return max
}
func main() {
	sc.Buffer(make([]byte, 128), 500000)
	sc.Split(bufio.ScanWords)
	n := scanInt()
	k := scanInt()
	// fmt.Printf("%d %d\n", n, k)

	p := make([]int, n)
	pq := make(PriorityQueue, k)
	for i := 0; i < n; i++ {
		p[i] = scanInt()
		if i < k {
			pq[i] = &Item{
				value:    i,
				priority: p[i],
				index:    i,
			}
		}
	}
	heap.Init(&pq)
	first := heap.Pop(&pq).(*Item)
	fmt.Printf("%d\n", first.priority)
	heap.Push(&pq, first)

	for i := k; i < n; i++ {
		min := heap.Pop(&pq).(*Item).priority
		min = maxOf(min, p[i])
		item := &Item{
			value:    i + 1,
			priority: min,
			index:    i,
		}
		heap.Push(&pq, item)

		ans := heap.Pop(&pq).(*Item)
		fmt.Printf("%d\n", ans.priority)
		heap.Push(&pq, ans)
	}
}

var sc = bufio.NewScanner(os.Stdin)

func scanInts(n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = scanInt()
	}
	return a
}
func scanInt() int {
	sc.Scan()
	return atoi(sc.Text())
}
func atoi(nStr string) int {
	i, err := strconv.Atoi(nStr)
	if err != nil {
		panic(err)
	}
	return i
}
func scanText() string {
	sc.Scan()
	return sc.Text()
}
