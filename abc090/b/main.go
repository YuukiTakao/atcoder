package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func isPalindromicNumber(n int) bool {
	str := strconv.Itoa(n)
	div := len(str) / 2
	slen := len(str)
	for i := 0; i < div; i++ {
		if str[i] != str[slen-1-i] {
			return false
		}
	}
	return true
}
func main() {
	sc.Buffer(make([]byte, 128), 500000)
	sc.Split(bufio.ScanWords)
	a := scanInt()
	b := scanInt()
	// fmt.Printf("%d %d\n", a, b)

	ans := 0
	for i := a; i <= b; i++ {
		if isPalindromicNumber(i) {
			ans++
		}
	}
	fmt.Printf("%d\n", ans)
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
