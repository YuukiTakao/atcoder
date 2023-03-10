package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	sc.Buffer(make([]byte, 128), 500000)
	sc.Split(bufio.ScanWords)

	l := scanInt()
	r := scanInt()
	nums := make([]int, 101)
	for i := l; i <= r; i++ {
		nums[i]++
	}

	l2 := scanInt()
	r2 := scanInt()
	for i := l2; i <= r2; i++ {
		nums[i]++
	}

	// fmt.Printf("nums=%v\n", nums)
	max := 0
	sum := 0
	for i := 1; i <= 100; i++ {
		if nums[i-1] == 2 && nums[i] == 2 {
			// fmt.Printf("sum=%d\n", sum)
			sum++
			if max < sum {
				max = sum
			}
		} else {
			sum = 0
		}
	}
	fmt.Printf("%d\n", max)
}

func maxOf(vars ...int) int {
	max := int(math.Pow10(18)) * -1
	for _, v := range vars {
		if max < v {
			max = v
		}
	}
	return max
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
