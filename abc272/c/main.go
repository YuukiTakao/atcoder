package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func main() {
	sc.Buffer(make([]byte, 128), 500000)
	sc.Split(bufio.ScanWords)
	n := scanInt()
	a := scanInts(n)

	sort.Sort(sort.Reverse(sort.IntSlice(a)))

	odd := make([]int, 0, 2)
	even := make([]int, 0, 2)
	ans := make([]int, 0, 2)
	for i := 0; i < n; i++ {
		if a[i]%2 == 0 {
			even = append(even, a[i])
		} else {
			odd = append(odd, a[i])
		}
		if len(even) == 2 {
			ans = append(ans, even[0]+even[1])
		}
		if len(odd) == 2 {
			ans = append(ans, odd[0]+odd[1])
		}
	}
	if len(ans) == 0 {
		fmt.Printf("-1\n")
	} else {
		fmt.Printf("%d\n", maxOf(ans...))
	}
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
