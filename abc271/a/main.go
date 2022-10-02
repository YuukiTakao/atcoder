package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	N := strToInt(readLine())
	if N <= 15 {
		fmt.Printf("0%X\n", N)
	} else {
		fmt.Printf("%X\n", N)
	}

	//hex := baseNInt(N, 16)
	//fmt.Printf("%x\n", hex)
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

func baseNInt(n int, base int) int {
	result := ""
	for n > 0 {
		result = intToStr(n%base) + result
		n = n / base
	}
	oct := 0
	if len(result) > 0 {
		oct = strToInt(result)
	}
	return oct
}

func intToStr(i int) string {
	s := strconv.Itoa(i)
	return s
}
