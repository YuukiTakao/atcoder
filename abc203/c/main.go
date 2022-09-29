package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	line1 := strings.Split(readLine(), " ")
	N := strToInt(line1[0])
	K := strToInt(line1[1])
	//fmt.Printf("%d %d\n", N, K)

	villages := make([][2]int, N)
	for i := 0; i < N; i++ {
		line2 := strings.Split(readLine(), " ")
		Ai := strToInt(line2[0])
		Bi := strToInt(line2[1])

		villages[i] = [2]int{Ai, Bi}
	}

	sort.Slice(villages, func(i, j int) bool {
		return villages[i][0] < villages[j][0]
	})

	//fmt.Printf("%v\n", villages)
	ans := K
	for _, village := range villages {
		A := village[0]
		B := village[1]

		if ans >= A {
			ans += B
		} else {
			break
		}
	}
	//fmt.Printf("ans %d K %d\n", ans, K)
	fmt.Println(ans)
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

func intToStr(i int) string {
	s := strconv.Itoa(i)
	return s
}
