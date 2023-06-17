package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type Point struct {
	y int
	x int
}

func main() {
	bufInit()
	defer wr.Flush()
	h := scanInt()
	w := scanInt()

	grid := make([][]byte, h)
	for i := 0; i < h; i++ {
		grid[i] = make([]byte, w)
		grid[i] = []byte(scanText())
	}

	// for i := 0; i < h; i++ {
	// 	for j := 0; j < w; j++ {
	// 		fprintf("%c", grid[i][j])
	// 	}
	// 	fprintln()
	// }
	leftUp := Point{-1, -1}
	rightUp := Point{-1, -1}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if grid[i][j] == '#' {
				if leftUp.x == -1 {
					leftUp.y = i
					leftUp.x = j
				}
				if leftUp.y >= 0 && (rightUp.y == -1 || leftUp.y == i) {
					rightUp.y = i
					rightUp.x = j
				}
			}
		}
	}
	leftDown := Point{-1, -1}
	rightDown := Point{-1, -1}
	for i := h - 1; i >= 0; i-- {
		for j := w - 1; j >= 0; j-- {
			if grid[i][j] == '#' {
				if rightDown.x == -1 {
					rightDown.y = i
					rightDown.x = j
				}
				if rightDown.y >= 0 && (leftDown.y == -1 || rightDown.y == i) {
					leftDown.y = i
					leftDown.x = j
				}
			}
		}
	}
	// fprintln(leftUp, rightUp, rightDown, leftDown)

	calcLeft := minOf(leftUp.x, leftDown.x)
	calcRight := maxOf(rightUp.x, rightDown.x)
	calcUp := minOf(leftUp.y, rightUp.y)
	calcDown := maxOf(leftDown.y, rightDown.y)
	// fprintln(calcLeft, calcRight, calcUp, calcDown)

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if calcLeft <= j && j <= calcRight && calcUp <= i && i <= calcDown {
				if grid[i][j] == '.' {
					fprintln(i+1, j+1)
					return
				}
			}
		}
	}
}

func minOf(vars ...int) int {
	min := int(math.Pow10(18))
	for _, v := range vars {
		if min > v {
			min = v
		}
	}
	return min
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
