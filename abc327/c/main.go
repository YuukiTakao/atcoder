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
	a := make([][]int, 9)
	for i := 0; i < 9; i++ {
		a[i] = make([]int, 9)
		for j := 0; j < 9; j++ {
			a[i][j] = scanInt()
		}
		// fprintln(a[i])
	}
	for i := 0; i < 9; i++ {
		rowCheck := make(map[int]bool, 9)
		for j := 0; j < 9; j++ {
			rowCheck[a[i][j]] = true
		}
		for k := 1; k <= 9; k++ {
			if !rowCheck[k] {
				fprintln("No")
				return
			}
		}
		colCheck := make(map[int]bool, 9)
		for j := 0; j < 9; j++ {
			colCheck[a[j][i]] = true
		}
		for k := 1; k <= 9; k++ {
			if !colCheck[k] {
				fprintln("No")
				return
			}
		}
	}

	groupCheck := make([]bool, 10)

	var check func(rowStart, colStart int) bool
	check = func(rowStart, colStart int) bool {
		for row := rowStart; row < rowStart+3; row++ {
			for col := rowStart; col < rowStart+3; col++ {
				groupCheck[a[row][col]] = true
			}
		}
		for i := 1; i <= 9; i++ {
			if !groupCheck[i] {
				return false
			}
		}
		return true
	}
	for i := 0; i <= 6; i += 3 {
		groupCheck = make([]bool, 10)
		for j := 0; j <= 6; j += 3 {
			if !check(i, j) {
				fprintln("No")
				return
			}
		}
	}

	fprintln("Yes")
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
