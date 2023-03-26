package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// 2次元座標のマンハッタン距離
func manhattanDistance(x1, y1, x2, y2 int) int {
	return abs(x1-x2) + abs(y1-y2)
}
func abs(n int) int {
	if n < 0 {
		n *= -1
	}
	return n
}

func isNum(c byte) bool {
	return '0' <= c && c <= '9'
}
func main() {
	bufInit()
	defer wr.Flush()
	r := scanInt()
	c := scanInt()

	grid := make([][]byte, r)
	for i := 0; i < r; i++ {
		grid[i] = []byte(scanText()) //
		// grid[i] = scanBytes() // これだと改行が入ってしまう
	}

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			if !isNum(grid[i][j]) {
				continue
			}
			bomb := int(grid[i][j] - '0')
			for k := 0; k < r; k++ {
				for l := 0; l < c; l++ {
					d := abs(i-k) + abs(j-l)
					if d <= bomb && grid[k][l] == '#' {
						grid[k][l] = '.'
					}
				}
			}
			grid[i][j] = '.'
		}
	}

	for i := 0; i < r; i++ {
		fprintln(string(grid[i]))
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
