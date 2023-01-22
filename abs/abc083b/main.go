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
	n := scanInt()
	a := scanInt()
	b := scanInt()
	// fmt.Printf("%d %d %d\n", n, a, b)

	sum := 0
	for i := 1; i <= n; i++ {
		// fmt.Printf("%d\n", digitSum(i))
		dsum := digitSum(i)
		if a <= dsum && dsum <= b {
			sum += i
		}
	}
	fmt.Printf("%d\n", sum)
}

var sc = bufio.NewScanner(os.Stdin)

func digitSum(n int) int {
	nstr := strconv.Itoa(n)
	sum := 0
	for _, v := range nstr {
		sum += int(v - '0')
	}
	return sum
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
