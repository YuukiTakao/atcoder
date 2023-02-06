package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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

// O(N log log N)
func enumPrimesByEratosthenes(x int) []int {
	deleted := make([]bool, 300000+1)
	primes := make([]int, 0, x+1)
	for i := 2; i <= x; i++ {
		deleted[i] = false
	}
	for i := 2; i*i <= 30000000; i++ {
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
func main() {
	sc.Buffer(make([]byte, 128), 500000)
	sc.Split(bufio.ScanWords)
	a := scanInt()
	b := scanInt()
	// fmt.Printf("%d %d\n", a, b)
	// fmt.Printf("%d\n", greatestCommonDivisor(a, b))
	fmt.Printf("%d\n", len(enumPrimesByEratosthenes(greatestCommonDivisor(a, b)))+1)

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
