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
	if p[i].second == p[j].second {
		return p[i].first < p[j].first
	}
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
	t := make([][]byte, n)
	players := make(Pairs, n)
	for i := 0; i < n; i++ {
		t[i] = make([]byte, n)
		t[i] = []byte(scanText())
		winCnt := 0
		for _, v := range t[i] {
			if v == 'o' {
				winCnt++
			}
		}
		players[i].first = i + 1
		players[i].second = winCnt
	}
	sort.Slice(players, players.Less)
	// fprintf("%v\n", players)
	for i, v := range players {
		fprintf("%d", v.first)
		if i < len(players)-1 {
			fprintf(" ")
		}
	}
	fprintln()
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
