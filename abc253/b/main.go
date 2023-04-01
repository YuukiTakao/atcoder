package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func GridDistanceManhattan(x1, y1, x2, y2 int) int {
	return abs(x1-x2) + abs(y1-y2)
}
func abs(n int) int {
	if n < 0 {
		n *= -1
	}
	return n
}

type Point struct {
	x int
	y int
}

func main() {
	bufInit()
	defer wr.Flush()
	h := scanInt()
	w := scanInt()
	// fprintf("%d %d\n", h, w)

	s := make([][]byte, h)
	for i := 0; i < h; i++ {
		s[i] = []byte(scanText())
	}

	op := make([]Point, 0, 2)
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if s[i][j] == 'o' {
				// fprintf("s[%d][%d]\n", i, j)
				op = append(op, Point{x: i, y: j})
			}
		}
	}
	// fprintf("%v\n", op[0])
	// fprintf("%v\n", op[1])
	fprintf("%d\n", GridDistanceManhattan(op[0].x, op[0].y, op[1].x, op[1].y))
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
