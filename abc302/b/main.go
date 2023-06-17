package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Point struct {
	x, y int
}

type Queue struct {
	data []Point
}

func NewQueue(n int) *Queue {
	return &Queue{data: make([]Point, n)}
}
func (q *Queue) Enqueue(n Point) {
	q.data = append(q.data, n)
}
func (q *Queue) Dequeue() Point {
	if len(q.data) == 0 {
	}
	a := q.data[0]
	q.data = q.data[1:]
	return a
}
func (q *Queue) Peek() Point {
	if len(q.data) == 0 {
		return Point{}
	}
	return q.data[0]
}
func (q *Queue) Len() int {
	return len(q.data)
}

func main() {
	bufInit()
	defer wr.Flush()
	h := scanInt()
	w := scanInt()

	ss := make([][]byte, h)
	for i := 0; i < h; i++ {
		ss[i] = []byte(scanText())
		fmt.Printf("%s\n", ss[i])
	}

	// pp := make([]Point, 5)
	seen := make([][]bool, h)
	for i := 0; i < h; i++ {
		seen[i] = make([]bool, w)
	}

	que := NewQueue(h * w)

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if ss[i][j] == 's' {
				que.Enqueue(Point{i, j})

				for k := 0; k < h; k++ {
					for l := 0; l < w; l++ {

					}
				}
			}
		}
	}

}

var wr *bufio.Writer
var sc = bufio.NewScanner(os.Stdin)

func fprintf(f string, a ...interface{}) {
	fmt.Fprintf(wr, f, a...)
}
func fprintln(a ...interface{}) {
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
func scanInts(n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = scanInt()
	}
	return a
}
func scanInts2(n int) ([]int, []int) {
	a, b := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		a[i], b[i] = scanInt(), scanInt()
	}
	return a, b
}
func scanInt() int {
	sc.Scan()
	return atoi(sc.Text())
}
func scanFloat() float64 {
	sc.Scan()
	f, e := strconv.ParseFloat(sc.Text(), 64)
	if e != nil {
		panic(e)
	}
	return f
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
func scanBytes() []byte {
	sc.Scan()
	return sc.Bytes()
}
