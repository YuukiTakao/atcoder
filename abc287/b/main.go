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
	m := scanInt()
	// fmt.Printf("%d %d\n", n, m)

	S := make([]string, n)
	for i := 0; i < n; i++ {
		S[i] = scanText()
	}

	T := make([]string, m)
	for i := 0; i < m; i++ {
		T[i] = scanText()
	}

	ans := make(map[int]int)
	// ans := 0
	for i := 0; i < n; i++ {
		// fmt.Printf("S[i][n-4:]=%v\n", S[i][3:])
		for j := 0; j < m; j++ {

			if S[i][3:] == T[j] {
				// fmt.Printf("s[i]=%s t[i]=%s\n", S[i][3:], T[j])
				ans[i]++
				// ans++
			}
		}
	}
	// fmt.Printf("%d\n", ans)
	fmt.Printf("%d\n", len(ans))
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
