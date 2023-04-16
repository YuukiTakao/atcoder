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
	n := scanInt()
	a := make([][]int, n)
	a90 := make([][]int, n)
	a180 := make([][]int, n)
	a270 := make([][]int, n)
	b := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, n)
		a90[i] = make([]int, n)
		a180[i] = make([]int, n)
		a270[i] = make([]int, n)
		b[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			a[i][j] = scanInt()
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			b[i][j] = scanInt()
		}
	}

	is0 := true
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if a[i][j] == 1 {
				if b[i][j] == 0 {
					is0 = false
				}
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			a90[i][j] = a[n-1-j][i]
		}
	}
	is90 := true
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if a90[i][j] == 1 {
				if b[i][j] == 0 {
					is90 = false
				}
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			a180[i][j] = a[n-1-i][n-1-j]
		}
	}
	is180 := true
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if a180[i][j] == 1 {
				if b[i][j] == 0 {
					is180 = false
				}
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			a270[i][j] = a[j][n-1-i]
		}
	}
	is270 := true
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if a270[i][j] == 1 {
				if b[i][j] == 0 {
					is270 = false
				}
			}
		}
	}

	if is0 || is90 || is180 || is270 {
		fprintln("Yes")

	} else {
		fprintln("No")
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
