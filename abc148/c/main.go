package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	line := strings.Split(readLine(), " ")
	A := strToInt(line[0])
	B := strToInt(line[1])
	//fmt.Printf("%d %d\n", A, B)

	max := A
	min := B
	if A < B {
		max = B
		min = A
	}
	ans := leastCommonMultiple(max, min)
	fmt.Printf("%d\n", ans)

}

// 最小公倍数
func leastCommonMultiple(a, b int) int {
	return (a * b) / greatestCommonDevisor(a, b)
}

// 最大公約数
func greatestCommonDevisor(a, b int) int {
	if b == 0 {
		return a
	}
	return greatestCommonDevisor(b, a%b)
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
