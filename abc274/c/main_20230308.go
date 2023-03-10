package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Queue struct {
	data []int
}

func NewQueue() *Queue {
	return &Queue{data: []int{}}
}
func (q *Queue) Enqueue(n int) {
	q.data = append(q.data, n)
}
func (q *Queue) Dequeue() int {
	if len(q.data) == 0 {
		return -1
	}
	a := q.data[0]
	q.data = q.data[1:]
	return a
}
func (q *Queue) Peek() int {
	if len(q.data) == 0 {
		return -1
	}
	return q.data[0]
}
func (q *Queue) Len() int {
	return len(q.data)
}

func repeatDiv2(n int) int {
	count := 0
	for n > 1 {
		n /= 2
		count++
	}
	return count
}

func main() {
	sc.Buffer(make([]byte, 128), 500000)
	sc.Split(bufio.ScanWords)
	n := scanInt()

	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = scanInt()
	}

	ans := make([]int, 2*n+1)
	for i, v := range a {
		ans[2*i+1] = ans[v-1] + 1
		ans[2*i+2] = ans[v-1] + 1
	}

	for _, v := range ans {
		fmt.Printf("%d\n", v)
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
