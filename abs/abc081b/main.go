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
	// fmt.Printf("%d\n", n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = scanInt()
	}

	evenCnt := 0
	finished := 0
	for i := 0; true; i++ {
		for k := 0; k < n; k++ {
			if a[k]%2 == 1 {
				break
			}
		}
		if finished == 1 {
			break
		}
		for j := 0; j < n; j++ {
			a[j] = a[j] / 2
		}
		evenCnt++
	}

	fmt.Printf("%d\n", evenCnt)
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
