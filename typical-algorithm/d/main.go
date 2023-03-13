package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
	"strconv"
)

func bufInit() {
	sc.Buffer(make([]byte, 128), 500000)
	sc.Split(bufio.ScanWords)
}

type AdjacencyList struct {
	all   map[int][]AlType
	fixed []bool
}

func NewAdlist(v_count int) AdjacencyList {
	al := AdjacencyList{
		all:   make(map[int][]AlType, v_count),
		fixed: make([]bool, v_count+1), // 1 indexed
	}
	return al
}
func (al AdjacencyList) Push(key int, v AlType) {
	al.all[key] = append(al.all[key], v)
}

type AlType struct {
	// pq
	value    int
	priority int
}
type PriorityQueue []*AlType      // Item
func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*AlType)
	*pq = append(*pq, item)
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}
func (g AdjacencyList) Dijkstra(n int) []int {
	distances := make([]int, n)
	inf := int(math.Pow10(18))
	for i := 0; i < n; i++ {
		distances[i] = inf
	}
	pq := new(PriorityQueue)
	heap.Init(pq)
	heap.Push(pq, &AlType{0, 0})
	for pq.Len() > 0 {
		node := heap.Pop(pq).(*AlType)
		pos := node.value
		cost := node.priority
		if distances[pos] <= cost {
			continue
		}
		distances[pos] = cost
		for _, next := range g.all[pos] {
			heap.Push(pq, &AlType{
				value:    next.value,
				priority: next.priority + cost,
			})
		}
	}
	return distances
}
func main() {
	bufInit()
	n := scanInt()
	m := scanInt()
	g := NewAdlist(n)
	for i := 0; i < m; i++ {
		u := scanInt()
		v := scanInt()
		cost := scanInt()
		g.Push(u, AlType{value: v, priority: cost})
	}
	distances := g.Dijkstra(n)
	fmt.Printf("%d\n", distances[n-1])
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
