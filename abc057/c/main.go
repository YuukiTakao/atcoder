package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func DivisorEnumeration(n int) []int {
	ans := make([]int, 0, 64)
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			ans = append(ans, i)
			if i != n/i {
				ans = append(ans, n/i)
			}
		}
	}
	return ans
}

func F(a, b int) int {
	da := digit(a)
	db := digit(b)

	if da >= db {
		return da
	} else {
		return db
	}
}

func digit(n int) int {
	digit := 0
	tmp := n
	if n < 0 {
		tmp *= -1
	}
	for tmp > 0 {
		tmp /= 10
		digit++
	}
	return digit
}
func main() {
	sc.Buffer(make([]byte, 128), 500000)
	sc.Split(bufio.ScanWords)
	n := scanInt()
	// fmt.Printf("%d\n", n)

	if n == 1 {
		fmt.Printf("%d\n", digit(n))
		return
	}
	min := int(10e17)
	divisors := DivisorEnumeration(n)
	for _, v := range divisors {
		for _, v2 := range divisors {
			if v*v2 == n {
				// fmt.Printf("v=%d v2=%d\n", v, v2)
				tmp := F(v, v2)
				if min > tmp {
					min = tmp
				}
			}
		}
	}
	fmt.Printf("%d\n", min)
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
