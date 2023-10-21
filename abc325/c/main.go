package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	bufInit()
	defer wr.Flush()
	h := scanInt()
	w := scanInt()

	g := make(map[int][]byte, h)
	for i := 0; i < h; i++ {
		g[i] = make([]byte, w)
		g[i] = []byte(scanText())
		// fprintf("g[%d]=%s\n", i, g[i])
	}

	visited := make([][]bool, h)
	for i := 0; i < h; i++ {
		visited[i] = make([]bool, w)
	}
	var dfsGrid func(y, x int)
	dfsGrid = func(y, x int) {
		visited[y][x] = true

		for dh := -1; dh <= 1; dh++ {
			for dw := -1; dw <= 1; dw++ {
				nh := y + dh
				nw := x + dw
				if nh < 0 || nh >= h || nw < 0 || nw >= w {
					continue
				}
				if visited[nh][nw] {
					continue
				}
				if g[nh][nw] != '#' {
					continue
				}
				dfsGrid(nh, nw)
			}
		}
	}
	ans := 0
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if g[i][j] == '#' && !visited[i][j] {
				dfsGrid(i, j)
				// fprintf("i=%d j=%d ans=%v\n", i, j, ans)
				ans += 1
			}
		}
	}
	fprintf("%d\n", ans)
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
	const startBufSize = 4096
	const maxScanTokenSize = math.MaxInt64
	sc.Buffer(make([]byte, startBufSize), maxScanTokenSize)
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
