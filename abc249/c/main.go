package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func calc(ans int, ss []int, s []string, k int) int {
	alps := make(map[rune]int, 26)
	for _, idx := range ss {
		for _, char := range s[idx] {
			alps[char]++
		}
	}
	max := 0
	for _, v := range alps {
		if v == k {
			max++
		}
	}
	if ans < max {
		ans = max
	}
	return ans
}

func bitSearch(n, k int, s []string) int {
	ans := 0
	for bit := 0; bit < 1<<n; bit++ {
		ss := make([]int, 0)
		for i := 0; i < n; i++ {
			if bit&(1<<i) > 0 {
				ss = append(ss, i)
			}
		}
		ans = calc(ans, ss, s, k)
	}
	return ans
}

func main() {
	sc.Buffer(make([]byte, 128), 500000)
	sc.Split(bufio.ScanWords)
	n := scanInt()
	k := scanInt()
	// fmt.Printf("%d %d\n", n, k)

	s := make([]string, n)
	for i := 0; i < n; i++ {
		s[i] = scanText()
	}

	fmt.Printf("%d\n", bitSearch(n, k, s))
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
