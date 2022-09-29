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
	//g1x := g1(N)
	//g2x := g2(N)
	K := strToInt(line1[1])
	//fmt.Printf("%d %d \n", g1x, g2x)

	ai := N
	for i := 0; i < K; i++ {
		ai = f(ai)
	}

	fmt.Printf("%d\n", ai)
}

func f(x int) int {
	xs := intToStr(x)
	return g1(xs) - g2(xs)
}

func g1(x string) int {
	xSl := strings.Split(x, "")
	sortedX := strings.Join(strSortReverse(xSl), "")
	return strToInt(sortedX)
}

func g2(x string) int {
	xSl := strings.Split(x, "")
	sort.Strings(xSl)
	return strToInt(strings.Join(xSl, ""))
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
	I, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return I
}

func intToStr(i int) string {
	s := strconv.Itoa(i)
	return s
}

func strSortReverse(s []string) []string {
	sort.Sort(sort.Reverse(sort.StringSlice(s)))
	return s
}
