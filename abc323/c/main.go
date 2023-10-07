package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func (p Pairs) Len() int      { return len(p) }
func (p Pairs) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p Pairs) Less(i, j int) bool {
	return p[i].second > p[j].second
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
	m := scanInt()
	a := make(Pairs, m)
	for i := 0; i < m; i++ {
		a[i].first = i
		a[i].second = scanInt()
	}
	s := make([][]byte, n)
	for i := 0; i < n; i++ {
		s[i] = []byte(scanText())
	}
	scores := make([]int, n)
	maxScore := 0
	for i := 0; i < n; i++ {
		sum := 0
		for j, v := range s[i] {
			if v == 'o' {
				sum += a[j].second
			}
		}
		sum += (i + 1)
		scores[i] = sum
		if maxScore < scores[i] {
			maxScore = scores[i]
		}
		// fprintf("scores[%d]=%d\n", i, scores[i])
	}
	sort.Slice(a, a.Less)

	// fprintf("maxScore=%d scores=%d\n", maxScore, scores)
	for i := 0; i < n; i++ {
		pSum := scores[i]
		// fprintf("pSum=%d maxScore=%d\n", pSum, maxScore)
		if pSum == maxScore {
			fprintf("%d\n", 0)
			continue
		}
		cnt := 0
		// fprintf("a=%v\n", a)
		for _, v := range a {
			if s[i][v.first] == 'x' {
				pSum += v.second
				cnt++
			}
			// fprintf("i=%d v=%d pSum=%d maxScore=%d\n", i, v, pSum, maxScore)

			if pSum >= maxScore {
				fprintf("%d\n", cnt)
				break
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
