package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func PrimeFactorize(n int) map[int]int {
	primeMap := make(map[int]int, 100000)
	for p := 2; p*p <= n; p++ {
		if n%p != 0 {
			continue
		}
		e := 0
		for n%p == 0 {
			e++
			n /= p
		}
		if e >= 3 {
			continue
		}
		primeMap[p] += e
	}
	if n != 1 {
		primeMap[n] += 1
	}
	return primeMap
}

// O(N log log N)
func enumPrimesByEratosthenes(x int) []int {
	deleted := make([]bool, 300000+1)
	primes := make([]int, 0, x+1)
	for i := 2; i <= x; i++ {
		deleted[i] = false
	}
	for i := 2; i*i <= 300000; i++ {
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
	bufInit()
	defer wr.Flush()
	n := scanInt()
	// primes := enumPrimesByEratosthenes(1000000000000)

	for i := 1; i <= n; i++ {
		f := PrimeFactorize(i)
		if len(f) != 3 {
			continue
		}
		fprintf("%d, %v\n", i, f)
	}
}

var wr *bufio.Writer
var sc = bufio.NewScanner(os.Stdin)

func fprintf(f string, a ...interface{}) {
	fmt.Fprintf(wr, f, a...)
}
func fprintln(a ...interface{}) {
	fmt.Fprintln(wr, a...)
}
func fprint(f string, a ...interface{}) {
	fmt.Fprint(wr, a...)
}
func bufInit() {
	sc.Buffer(make([]byte, 128), 500000)
	sc.Split(bufio.ScanWords)
	wr = bufio.NewWriter(os.Stdout)
}
func scanInts(n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = scanInt()
	}
	return a
}
func scanInts2(n int) ([]int, []int) {
	a, b := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		a[i], b[i] = scanInt(), scanInt()
	}
	return a, b
}
func scanInt() int {
	sc.Scan()
	return atoi(sc.Text())
}
func scanFloat() float64 {
	sc.Scan()
	f, e := strconv.ParseFloat(sc.Text(), 64)
	if e != nil {
		panic(e)
	}
	return f
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
func scanBytes() []byte {
	sc.Scan()
	return sc.Bytes()
}
