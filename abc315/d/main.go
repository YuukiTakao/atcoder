package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

type Point struct {
	x int
	y int
}

func main() {
	bufInit()
	defer wr.Flush()
	h := scanInt()
	w := scanInt()
	c := make([][]byte, h)
	for i := 0; i < h; i++ {
		c[i] = make([]byte, w)
		c[i] = []byte(scanText())
	}
	cntW := make([][]int, h)
	for i := 0; i < h; i++ {
		cntW[i] = make([]int, 26)
	}
	cntH := make([][]int, w)
	for i := 0; i < w; i++ {
		cntH[i] = make([]int, 26)
	}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			cntW[i][c[i][j]-'a']++
			cntH[j][c[i][j]-'a']++
		}
	}
	// fprintf("cntW=%v\n", cntW)
	// fprintf("cntH=%v\n", cntH)
	hc := h
	wc := w
	fx := make([]int, h)
	fy := make([]int, w)
	for i := 0; i < h+w; i++ {
		ux := make([]Point, 0)
		uy := make([]Point, 0)
		for i := 0; i < h; i++ {
			if fx[i] > 0 {
				continue
			}
			for j := 0; j < 26; j++ {
				if cntW[i][j] == wc && wc >= 2 {
					ux = append(ux, Point{i, j})
				}
			}
		}
		for i := 0; i < w; i++ {
			if fy[i] > 0 {
				continue
			}
			for j := 0; j < 26; j++ {
				if cntH[i][j] == hc && hc >= 2 {
					uy = append(uy, Point{i, j})
				}
			}
		}
		for _, v := range ux {
			fx[v.x] = 1
			hc--
			for j := 0; j < w; j++ {
				cntH[j][v.y]--
			}
		}
		for _, v := range uy {
			fy[v.x] = 1
			wc--
			for j := 0; j < h; j++ {
				cntW[j][v.y]--
			}
		}
	}
	fprintf("%d\n", hc*wc)
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
