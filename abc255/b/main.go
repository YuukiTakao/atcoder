package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type Point struct {
	x int
	y int
}

// ルート計算はループ外でやるのが早い
func GridDistanceEuclidean(p1, p2 Point) int {
	// return math.Sqrt(math.Pow(float64(p1.x-p2.x), 2) + math.Pow(float64(p1.y-p2.y), 2))
	return (p1.x-p2.x)*(p1.x-p2.x) + (p1.y-p2.y)*(p1.y-p2.y)
}
func main() {
	bufInit()
	defer wr.Flush()
	n := scanInt()
	k := scanInt()
	a := make([]int, k)
	for i := 0; i < k; i++ {
		a[i] = scanInt() - 1
	}
	p := make([]Point, n)
	for i := 0; i < n; i++ {
		p[i].x = scanInt()
		p[i].y = scanInt()
	}
	ans := 0
	for i := 0; i < n; i++ {
		min := int(math.Pow10(18))
		for _, lighter := range a {
			r := GridDistanceEuclidean(p[i], p[lighter])
			if min > r {
				min = r
			}
		}
		if ans < min {
			ans = min
		}
	}
	fprintf("%.12f\n", math.Sqrt(float64(ans)))
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
