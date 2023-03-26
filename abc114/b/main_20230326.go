package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func abs(n int) int {
	if n < 0 {
		n *= -1
	}
	return n
}
func main() {
	sc.Buffer(make([]byte, 128), 500000)
	sc.Split(bufio.ScanWords)
	s := scanText()
	// fmt.Printf("%s\n", s)

	min := int(10e10)
	for i := 0; i+2 < len(s); i++ {
		num := atoi(string(s[i]))*100 + atoi(string(s[i+1]))*10 + atoi(string(s[i+2]))
		sub := abs(int(753) - int(num))
		if min > sub {
			min = sub
		}
	}
	fmt.Printf("%d\n", min)
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
