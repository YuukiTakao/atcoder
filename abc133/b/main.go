package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func abs(n int) int {
	if n < 0 {
		n *= -1
	}
	return n
}

func distance(x, y []int, d int) float64 {
	sum := 0
	for i := 0; i < d; i++ {
		sum += (x[i] - y[i]) * (x[i] - y[i])
	}
	return math.Sqrt(float64(sum))
}
func main() {
	sc.Buffer(make([]byte, 128), 500000)
	sc.Split(bufio.ScanWords)
	n := scanInt()
	d := scanInt()

	x := make([][]int, n)
	for i := 0; i < n; i++ {
		x[i] = make([]int, d)
		for j := 0; j < d; j++ {
			x[i][j] = scanInt()
		}
	}

	ans := 0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			distance := distance(x[i], x[j], d)
			if distance == math.Floor(distance) {
				ans++
			}
		}
	}
	fmt.Printf("%d\n", ans)
}

var sc = bufio.NewScanner(os.Stdin)

func scanInts(n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = scanInt()
	}
	return a
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
