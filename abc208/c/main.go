package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func sortedMap(n int, m map[int]int, sorted []int) {
	for i := 0; i < n; i++ {
		m[sorted[i]]++
	}
}
func deepCopySlice(s []int) []int {
	tmp := make([]int, len(s))
	copy(tmp, s)
	return tmp
}
func main() {
	sc.Buffer(make([]byte, 128), 500000)
	sc.Split(bufio.ScanWords)
	n := scanInt()
	k := scanInt()
	// fmt.Printf("%d %d\n", n, k)
	// ソート済みのキーを用意する
	// map[国民番号]お菓子所持数
	// 配る

	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = scanInt()
	}

	sorted := deepCopySlice(a)
	sort.Ints(sorted)

	snack := make(map[int]int, n)
	if k < n {
		for i := 0; i < k; i++ {
			snack[sorted[i]]++
		}
		for i := 0; i < n; i++ {
			fmt.Printf("%d\n", snack[a[i]])
		}
	} else {
		perPerson := k / n
		rem := k % n
		for i := 0; i < n; i++ {
			snack[sorted[i]] += perPerson
			if i+1 <= rem {
				snack[sorted[i]]++
			}
		}
		for i := 0; i < n; i++ {
			fmt.Printf("%d\n", snack[a[i]])
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
