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
	sc.Buffer(make([]byte, 128), 500000)
	sc.Split(bufio.ScanWords)
	n := scanInt()
	// fmt.Printf("%d\n", n)

	p := make([]Point, n)
	for i := 0; i < n; i++ {
		x, y := scanInt(), scanInt()
		p[i] = Point{y: y, x: x}
		// fmt.Printf("%v\n", p[i])
	}

	max := float64(0)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n; j++ {
			// fmt.Printf("%d %d\n", i, j)
			xdis := abs(p[i].x - p[j].x)
			ydis := abs(p[i].y - p[j].y)
			dis := math.Sqrt(float64(xdis*xdis + ydis*ydis))
			if max < dis {
				max = dis
			}
		}
	}
	fmt.Printf("%.10f\n", max)
}
func abs(n int) int {
	if n < 0 {
		n *= -1
	}
	return n
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
