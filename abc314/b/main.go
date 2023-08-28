package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Person struct {
	id  int
	cnt int
}

func main() {
	bufInit()
	defer wr.Flush()
	n := scanInt()

	rlt := make(map[int][]int, n)
	cnt := make(map[int]int, n)
	for i := 0; i < n; i++ {
		c := scanInt()
		cnt[i+1] = c
		for j := 0; j < c; j++ {
			a := scanInt()
			rlt[a] = append(rlt[a], i+1)
			// fprintf("rlt[%d]=%d\n", a, rlt[a])
		}
	}
	result := scanInt()
	// fprintf("result=%d\n", result)
	persons := make([]Person, 0, n)
	for _, a := range rlt[result] {
		// fprintf("i=%d a=%d cnt[a]=%d\n", i, a, cnt[a])
		persons = append(persons, Person{a, cnt[a]})
	}
	// fprintf("persons=%v\n", persons)
	if len(persons) == 0 {
		fprintf("%d\n", 0)
		fprintln("")
		return
	}
	sort.Slice(persons, func(i, j int) bool {
		return persons[i].cnt < persons[j].cnt
	})
	min := persons[0].cnt
	// fprintf("min=%d\n", min)
	ans := make([]int, 0, n)
	for _, p := range persons {
		if p.cnt == min {
			ans = append(ans, p.id)
		}
	}
	sort.Ints(ans)
	fprintf("%d\n", len(ans))
	for _, v := range ans {
		fprintf("%d ", v)
	}
	fprintln("")
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
