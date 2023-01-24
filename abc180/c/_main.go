package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	N := strToInt(readLine())

	//fmt.Printf("%d\n", N)

	ans := make([]int, 0)
	for i := 1; i*i <= N; i++ {
		if N%i == 0 {
			ans = append(ans, i)
			if i != N/i {
				ans = append(ans, N/i)
			}
			//fmt.Printf("%d\n", N/i)
		}
	}
	sort.Ints(ans)

	for _, an := range ans {
		fmt.Printf("%d\n", an)
	}
}

var reader = bufio.NewReader(os.Stdin)

func readLine() string {
	r, _, err := reader.ReadLine()
	if err != nil {
		log.Fatal(err)
	}
	return string(r)
}

func strToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return i
}
