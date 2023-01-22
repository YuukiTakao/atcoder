package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc.Buffer(make([]byte, 4096), 1000000)
	sc.Split(bufio.ScanWords)
	n := scanInt()
	a := scanInt()
	b := scanInt()
	s := scanText()
	// fmt.Printf("%d %d %d\n", n, a, b)
	// fmt.Printf("%s\n", s)

	minCost := a*n + b*n
	rcnt := 0
	for i := 0; i < n; i++ {
		// rotated := s[i:] + s[0:i]
		var sb strings.Builder
		sb.WriteString(s[i:])
		sb.WriteString(s[0:i])
		rotated := sb.String()

		diff := 0
		for left, right := 0, n-1; left < right; left, right = left+1, right-1 {
			if rotated[left] != rotated[right] {
				diff++
			}
		}
		cost := rcnt*a + diff*b
		if minCost > cost {
			minCost = cost
		}
		rcnt++
	}
	fmt.Printf("%d\n", minCost)
}

func minOf(vars ...int) int {
	min := 2100000000
	for _, v := range vars {
		if min > v {
			min = v
		}
	}
	return min
}

var sc = bufio.NewScanner(os.Stdin)

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
