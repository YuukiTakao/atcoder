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
	n := scanInt()

	i := n
	s := make([]int, 0)
	for {
		s = append(s, i)
		if i == 0 {
			break
		}
		fmt.Printf("i=%b (i-1)=%b (i-1)&n=%b (i-1)&n=%d\n", i, (i - 1), (i-i)&n, (i-1)&n)
		i = (i - 1) & n
	}
	// sort.Ints(s)
	// for _, v := range s {
	// 	fmt.Println(v)
	// }

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
