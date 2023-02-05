package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type adList struct {
	l     map[int][]int
	paths [][]int
	path  []int
}

func initAdlist(v_count int) adList {
	return adList{
		l:     make(map[int][]int, v_count+1),
		paths: make([][]int, 0, 2),
		path:  make([]int, 0, 2),
	}
}
func (al *adList) appendPaths() {
	tmp := make([]int, len(al.path))
	copy(tmp, al.path)
	al.paths = append(al.paths, tmp)
}
func (al adList) printWithSpace() {
	for _, l := range al.l {
		for j, v := range l {
			if j >= 1 {
				fmt.Printf(" ")
			}
			fmt.Printf("%d", v)
		}
		fmt.Printf("\n")
	}
}
func main() {
	sc.Buffer(make([]byte, 128), 500000)
	sc.Split(bufio.ScanWords)
	n := scanInt()
	m := scanInt()
	// fmt.Printf("%d %d\n", n, m)

	al := initAdlist(n)
	for i := 0; i < m; i++ {
		a, b := scanInt(), scanInt()
		al.l[a] = append(al.l[a], b)
		al.l[b] = append(al.l[b], a)
	}
	al.printWithSpace()
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
