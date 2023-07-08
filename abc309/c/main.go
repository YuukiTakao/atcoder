package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func (p Pairs) Len() int      { return len(p) }
func (p Pairs) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p Pairs) Less(i, j int) bool {
	return p[i].first < p[j].first
}

type Pairs []Pair
type Pair struct {
	first  int
	second int
}

func main() {
	bufInit()
	defer wr.Flush()
	n := scanInt()
	k := scanInt()

	p := make(Pairs, n)
	sum := 0
	for i := 0; i < n; i++ {
		a := scanInt()
		b := scanInt()
		p[i].first = a
		p[i].second = b
		sum += b
	}
	sort.Sort(p)
	if k >= sum {
		fprintf("%d\n", 1)
		return
	}
	for i := 0; i < n; i++ {
		sum -= p[i].second
		if k >= sum {
			fprintf("%d\n", p[i].first+1)
			return
		}
	}
	// fprintf("%d %d\n", sum, k)
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
