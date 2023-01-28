package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	sc.Buffer(make([]byte, 128), 500000)
	sc.Split(bufio.ScanWords)
	h := scanInt()
	w := scanInt()
	// fmt.Printf("%d %d\n", h, w)

	grid := make([][]int, h)
	minCell := int(10e17)
	for i := 0; i < h; i++ {
		grid[i] = make([]int, w)

		for j := 0; j < w; j++ {
			grid[i][j] = scanInt()
			if minCell > grid[i][j] {
				minCell = grid[i][j]
			}
		}
	}

	ans := 0
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if grid[i][j]-minCell < 0 {
				log.Fatalf("error:%s", "最小のマスが取れてない")
			}
			ans += grid[i][j] - minCell
		}
	}
	fmt.Printf("%d\n", ans)
}

var sc = bufio.NewScanner(os.Stdin)

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
