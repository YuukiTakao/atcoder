package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	bufInit()
	defer wr.Flush()
	h := scanInt()
	w := scanInt()

	g := make([][]byte, h)
	for i := 0; i < h; i++ {
		g[i] = make([]byte, w)
		g[i] = []byte(scanText())
		// fprintf("%s\n", g[i])
	}

	visited := make([][]bool, h+1)
	for i := 0; i <= h; i++ {
		visited[i] = make([]bool, w+1)
	}

	for i := 0; ; {
		for j := 0; ; {
			// fprintf("i=%d j=%d\n", i, j)
			if !visited[i][j] {
				visited[i][j] = true
				if g[i][j] == 'R' {
					if j == w-1 {
						fprintln(i+1, j+1)
						return
					}
					j++
				} else if g[i][j] == 'L' {
					if j == 0 {
						fprintln(i+1, j+1)
						return
					}
					j--
				} else if g[i][j] == 'U' {
					if i == 0 {
						fprintln(i+1, j+1)
						return
					}
					i--
				} else if g[i][j] == 'D' {
					if i == h-1 {
						fprintln(i+1, j+1)
						return
					}
					i++
				}
			} else {
				fprintln(-1)
				return
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
