package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func PrimeFactorize(n int) map[int]int {
	primeMap := make(map[int]int, 1000)
	for p := 2; p*p <= n; p++ {
		if n%p != 0 {
			continue
		}

		e := 0
		for n%p == 0 {
			e++
			n /= p
		}
		primeMap[p] += e
	}
	if n != 1 {
		primeMap[n] += 1
	}
	return primeMap
}

func greatestCommonDivisor(a, b int) int {
	if a < b {
		tmp := a
		a = b
		b = tmp
	}
	r := a % b
	for r != 0 {
		a = b
		b = r
		r = a % b
	}
	return b
}
func main() {
	sc.Buffer(make([]byte, 128), 500000)
	sc.Split(bufio.ScanWords)
	a := scanInt()
	b := scanInt()
	// fmt.Printf("%d %d\n", a, b)

	// fmt.Printf("greatestCommonDivisor=%v\n", greatestCommonDivisor(a, b))
	fmt.Printf("%d\n", len(PrimeFactorize(greatestCommonDivisor(a, b)))+1)
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
