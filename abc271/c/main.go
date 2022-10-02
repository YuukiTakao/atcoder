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
	N := strToInt(readLine())
	fmt.Printf("%d\n", N)
	ai := strings.Split(readLine(), " ")

	comics := make([]int, N)
	for i, a := range ai {
		//fmt.Printf("%d\n", strToInt(a))
		comics[i] = strToInt(a)
	}

	sort.Ints(comics)
	fmt.Printf("%v\n", comics)
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
