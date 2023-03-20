package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func bufInit() {
	sc.Buffer(make([]byte, 128), 500000)
	sc.Split(bufio.ScanWords)
}

type Queue struct {
	data []int
}

func NewQueue(n int) *Queue {
	return &Queue{
		data: make([]int, 0, n),
	}
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

func main() {
	bufInit()
	n := scanInt()
	q := scanInt()

	ans := 1
	m := make(map[int]bool, n)
	for i := 0; i < q; i++ {
		t := scanInt()
		if t == 1 {

		} else if t == 2 {
			x := scanInt()
			m[x] = true
		} else if t == 3 {
			for m[ans] {
				ans++
			}
			fmt.Printf("%d\n", ans)
		} else {
			log.Fatal("t error")
		}
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
