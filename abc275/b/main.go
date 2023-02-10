package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Buffer(make([]byte, 128), 500000)
	sc.Split(bufio.ScanWords)
	mod := 998244353
	a := scanInt() % mod
	b := scanInt() % mod
	c := scanInt() % mod
	d := scanInt() % mod
	e := scanInt() % mod
	f := scanInt() % mod
	abc := (a * b % mod) * c % mod
	def := (d * e % mod) * f % mod
	fmt.Printf("%d\n", (abc-def+mod)%mod)

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
