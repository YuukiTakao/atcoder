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
	N := strToInt(readLine())

	ans := 0
	for i := 1; i <= N; i++ {
		octI := baseNInt(i, 8)

		iInclude7 := strings.Contains(intToStr(i), "7")
		octIInclude7 := strings.Contains(intToStr(octI), "7")

		if !iInclude7 && !octIInclude7 {
			ans += 1
		}
	}
	fmt.Printf("%d\n", ans)
}

func baseNInt(n int, base int) int {
	result := ""
	for n > 0 {
		result = intToStr(n%base) + result
		n = n / 8
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
