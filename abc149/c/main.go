package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

// O(N log log N)
func enumPrimesByEratosthenes(x int) []int {
	deleted := make([]bool, 300000+1)
	primes := make([]int, 0, x+1)
	for i := 2; i <= x; i++ {
		deleted[i] = false
	}
	for i := 2; i*i <= 100000000000; i++ {
		if deleted[i] {
			continue
		}
		if i > x {
			break
		}
		primes = append(primes, i)
		for j := i * 2; j <= x; j += i {
			deleted[j] = true
		}
	}
	return primes
}
func main() {
	sc.Buffer(make([]byte, 128), 500000)
	sc.Split(bufio.ScanWords)
	x := scanInt()
	// fmt.Printf("%d\n", x)
	primes := enumPrimesByEratosthenes(x + 100)
	sort.Ints(primes)
	for i := 0; i < x+100; i++ {
		if primes[i] >= x {
			fmt.Printf("%d\n", primes[i])
			return
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
