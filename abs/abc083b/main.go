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
	fmt.Printf("%d %d %d\n", n, a, b)

	sum := 0
	for i := 0; i < n; i++ {
		fmt.Printf("%s\n", i+'0')
	}

	// fmt.Printf("%d\n", digit(1))
	// fmt.Printf("%d\n", digit(12))
	// fmt.Printf("%d\n", digit(111))
	// fmt.Printf("%d\n", digit(1111))

	fmt.Printf("%d\n", sum)
}

var sc = bufio.NewScanner(os.Stdin)

// func digit_sum(n string) int {
// 	digit := 0
// 	for n > 0 {
// 		n /= 10
// 		digit++
// 	}
// 	return digit
// }

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
