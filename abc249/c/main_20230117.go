package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func intToSlice(bit, n int) []int {
	s := make([]int, 0, n)

	return s
}

func main() {
	sc.Split(bufio.ScanWords)
	n := scanInt()
	k := scanInt()
	// fmt.Printf("%d %d\n", n, k)

	S := make([]string, n)
	for i := 0; i < n; i++ {
		S[i] = scanText()
	}

	ans := 0
	for bit := 0; bit < 1<<n; bit++ {
		alphabets := make(map[string]int, 26)
		for i := 0; i < n; i++ {
			fmt.Printf("bit=%d i=%d bit & (1 << i)=%d\n", bit, i, bit&(1<<i))
			if (bit & (1 << i)) != 0 {
				for _, c := range S[i] {
					// fmt.Printf("%d %d %c\n", i, j, c)
					alphabets[string(c)]++
				}
			}
		}

		max := 0
		for _, a := range alphabets {
			// fmt.Printf("char=%s count=%d\n", key, a)
			if a == k {
				max++
			}
		}
		if ans < max {
			ans = max
		}
	}

	fmt.Printf("%d\n", ans)
}

var sc = bufio.NewScanner(os.Stdin)

func scanInt() int {
	sc.Scan()
	i, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return i
}
func scanText() string {
	sc.Scan()
	return sc.Text()
}
