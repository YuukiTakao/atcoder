package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type key struct {
	k1, k2 string
}
type tmap map[key]bool

func (t tmap) insert(k1, k2 string, v bool) {
	t[key{k1, k2}] = v
}
func (t tmap) at(k1, k2 string) bool {
	return t[key{k1, k2}]
}
func (t tmap) erase(k1, k2 string) {
	delete(t, key{k1, k2})
}

func main() {
	sc.Buffer(make([]byte, 128), 500000)
	sc.Split(bufio.ScanWords)
	_ = scanInt()
	q := scanInt()
	// fmt.Printf("%d %d\n", n, q)

	m := make(tmap, q)

	for i := 0; i < q; i++ {
		t := scanInt()
		a := scanText()
		b := scanText()

		if t == 1 {
			m.insert(a, b, true)
		}
		if t == 2 {
			m.erase(a, b)
		}
		if t == 3 {
			if m.at(a, b) && m.at(b, a) {
				fmt.Println("Yes")
			} else {
				fmt.Println("No")
			}
		}
	}
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
