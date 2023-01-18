package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	n := scanInt()
	d := scanInt()
	// fmt.Printf("%d %d\n", n, d)
	ans := 0
	for i := 0; i < n; i++ {
		x := scanInt()
		y := scanInt()

		dd := float64(x*x + y*y)
		// fmt.Printf("x=%d\n", x)
		// fmt.Printf("y=%d\n", y)
		if float64(d) >= math.Sqrt(dd) {
			ans++
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
