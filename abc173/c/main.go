package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	h := scanInt()
	w := scanInt()
	k := scanInt()
	// fmt.Printf("%d %d %d\n", h, w, k)

	maps := make([]string, h)
	for i := 0; i < h; i++ {
		maps[i] = scanText()
	}

	sum := 0
	hcount := make([]int, h+1)
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			// fmt.Printf("maps[j]=%s\n", maps[j])
			if string(maps[i][j]) == "#" {
				sum++
				hcount[j]++
			}
		}
	}
	// fmt.Printf("sum=%d\n", sum)
	wcount := make([]int, w+1)
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			if string(maps[j][i]) == "#" {
				wcount[j]++
			}
		}
	}

	ans := 0
	if sum == k {
		ans++
	}
	for i := 0; i < w; i++ {
		// fmt.Printf("hc=%d\n", hcount[i])
		if k == sum-hcount[i] {
			ans++
		}
	}
	for i := 0; i < h; i++ {
		// fmt.Printf("wc=%d\n", wcount[i])
		if k == sum-wcount[i] {
			ans++
		}
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if k == sum-wcount[i]-hcount[j] {
				ans++

			}
		}
	}
	fmt.Printf("%d\n", ans)
}

var sc = bufio.NewScanner(os.Stdin)

func scanInt() int {
	sc.Scan()
	i, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return i
}
func scanText() string {
	sc.Scan()
	return sc.Text()
}
