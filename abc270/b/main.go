package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func bufInit() {
	sc.Buffer(make([]byte, 128), 500000)
	sc.Split(bufio.ScanWords)
}
func main() {
	bufInit()
	x := scanInt()
	y := scanInt()
	z := scanInt()
	// fmt.Printf("%d %d %d\n", x, y, z)

	if 0 < x {
		if x < y || y < 0 {
			fmt.Printf("%d\n", abs(x))
		} else {
			if z < y {
				if z < 0 {
					fmt.Printf("%d\n", abs(z+z)+x)
				} else {
					fmt.Printf("%d\n", x)
				}
			} else {
				fmt.Printf("-1\n")
			}
		}
	} else { // x <= 0
		if x > y || y > 0 { // 壁が目標よりも遠いか、プラス側にあるからまっすぐ目標にいける
			fmt.Printf("%d\n", abs(x))
		} else { // 目標と原点の間に壁
			if z > y { // カベよりハンマーの方が近い
				if z > 0 {
					fmt.Printf("%d\n", abs(z+z)+abs(x))
				} else {
					fmt.Printf("%d\n", abs(x))
				}
			} else {
				fmt.Printf("-1\n")
			}
		}
	}

}

func abs(n int) int {
	if n < 0 {
		n *= -1
	}
	return n
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
