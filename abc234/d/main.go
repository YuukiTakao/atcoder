package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
	"strconv"
)

type IntHeap []int

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func maxOf(vars ...int) int {
	max := int(math.Pow10(18)) * -1
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
	h := new(IntHeap)
	for i := 0; i < n; i++ {
		p[i] = scanInt()
		if i < k {
			heap.Push(h, p[i])
		}
	}

	heap.Init(h)
	fmt.Printf("%d\n", (*h)[0])
	for i := k; i < n; i++ {
		min := heap.Pop(h).(int)
		min = maxOf(min, p[i])
		heap.Push(h, min)
		ans := heap.Pop(h)
		fmt.Printf("%d\n", ans)
		heap.Push(h, ans)
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
