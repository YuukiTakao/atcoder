package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Point struct {
	x int
	y int
}

func NewPoint(x, y int) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func cutRectangle(ss [][]byte) [][]byte {
	top := -1
	bottom := -1
	left := -1
	right := -1
	for i, s := range ss {
		for j, c := range s {
			if c == '#' {
				if top == -1 {
					top = i
				}
				bottom = i
				if left == -1 || left > j {
					left = j
				}
				if right == -1 || right < j {
					right = j
				}
			}
		}
	}
	x := right - left + 1
	y := bottom - top + 1
	newSS := make([][]byte, y)
	for i := top; i <= bottom; i++ {
		newSS[i-top] = make([]byte, x)
		for j := left; j <= right; j++ {
			newSS[i-top][j-left] = ss[i][j]
		}
	}
	return newSS
}

func main() {
	bufInit()
	defer wr.Flush()
	ha := scanInt()
	wa := scanInt()

	a := make([][]byte, ha)
	for i := 0; i < ha; i++ {
		a[i] = make([]byte, wa)
		a[i] = []byte(scanText())
	}

	hb := scanInt()
	wb := scanInt()
	b := make([][]byte, hb)
	for i := 0; i < hb; i++ {
		b[i] = make([]byte, wb)
		b[i] = []byte(scanText())
	}

	hx := scanInt()
	wx := scanInt()
	x := make([][]byte, hx)
	for i := 0; i < hx; i++ {
		x[i] = make([]byte, wx)
		x[i] = []byte(scanText())

	}

	cutA := cutRectangle(a)
	// for i := 0; i < len(cutA); i++ {
	// 	fprintf("%s\n", cutA[i])
	// }
	cutB := cutRectangle(b)

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
