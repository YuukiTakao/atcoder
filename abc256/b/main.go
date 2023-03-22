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

func NewQueue(n int) *Queue {
	return &Queue{
		data: make([]int, 0),
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

	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = scanInt()
	}
	p := NewQueue(n)

}

var wr *bufio.Writer
var sc = bufio.NewScanner(os.Stdin)

func fprintf(f string, a ...interface{}) {
	fmt.Fprintf(wr, f, a...)
}
func fprintln(f string, a ...interface{}) {
	fmt.Fprintln(wr, a...)
}
func fprint(f string, a ...interface{}) {
	fmt.Fprint(wr, a...)
}
func bufInit() {
	sc.Buffer(make([]byte, 128), 500000)
	sc.Split(bufio.ScanWords)
	wr = bufio.NewWriter(os.Stdout)
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
