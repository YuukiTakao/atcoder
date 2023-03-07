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
	a := scanInt()
	b := scanInt()
	c := scanInt()
	d := scanInt()
	e := scanInt()
	f := scanInt()
	x := scanInt()

	tDistance := 0
	ttime := 0
	for {
		if ttime+a <= x {
			tDistance += a * b
		} else {
			tDistance += (x - ttime) * b
		}

		ttime += a + c
		if ttime >= x {
			break
		}
	}
	aDistance := 0
	atime := 0
	for {
		if atime+a <= x {
			aDistance += d * e
		} else {
			aDistance += (x - atime) * e
		}

		atime += d + f
		if atime >= x {
			break
		}
	}
	if tDistance > aDistance {
		fmt.Printf("Takahashi\n")
	} else if tDistance == aDistance {
		fmt.Printf("Draw\n")
	} else {
		fmt.Printf("Aoki\n")
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
