package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func DivisorEnumeration(n int) []int {
	divs := make([]int, 0, 64)
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			divs = append(divs, i)
			if i != n/i {
				divs = append(divs, n/i)
			}
		}
	}
	return divs
}

func filterUnder10(arr []int) []int {
	ret := make([]int, 0, 10)
	for i := 0; i < len(arr); i++ {
		if arr[i] <= 9 {
			ret = append(ret, arr[i])
		}
	}
	return ret
}

func main() {
	bufInit()
	defer wr.Flush()
	n := scanInt()

	divs := filterUnder10(DivisorEnumeration(n))

	// fprintf("%v\n", divs)
	for i := 0; i <= n; i++ {
		nums := make(map[int]bool, 10)
		for _, j := range divs {
			if i%(n/j) == 0 {
				nums[j] = true
			}
		}
		if len(nums) == 0 {
			fprintf("%s", "-")
		} else {
			min := 100
			for num := range nums {
				if min > num {
					min = num
				}
			}
			fprintf("%d", min)
		}
	}
	fprintln()
}

func minOf(vars ...int) int {
	min := int(math.Pow10(18))
	for _, v := range vars {
		if min > v {
			min = v
		}
	}
	return min
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
