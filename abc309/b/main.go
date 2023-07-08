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
	n := scanInt()

	a := make([][]byte, n)
	for i := 0; i < n; i++ {
		a[i] = make([]byte, n)
		a[i] = []byte(scanText())
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			calcI := i
			calcJ := j
			if i > 0 && i < n-1 && j > 0 && j < n-1 {
				fprintf("%c", a[calcI][calcJ])
			} else {
				if i == 0 {
					if j == 0 {
						fprintf("%c", a[calcI+1][calcJ])
					} else if j == n-1 {
						fprintf("%c", a[calcI][calcJ-1])
					} else {
						fprintf("%c", a[calcI][calcJ-1])
					}
				} else if i == n-1 {
					if j == 0 {
						fprintf("%c", a[calcI][calcJ+1])
					} else if j == n-1 {
						fprintf("%c", a[calcI-1][calcJ])
					} else {
						fprintf("%c", a[calcI][calcJ+1])
					}
				} else {
					if j == 0 {
						fprintf("%c", a[calcI+1][calcJ])
					} else if j == n-1 {
						fprintf("%c", a[calcI-1][calcJ])
					}
				}
			}
		}
		fprintln()
	}
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

func minOf(vars ...int) int {
	min := int(math.Pow10(18))
	for _, v := range vars {
		if min > v {
			min = v
		}
	}
	return min
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
