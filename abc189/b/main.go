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
	line1 := strings.Split(readLine(), " ")
	N := strToInt(line1[0])
	X := strToInt(line1[1])
	//fmt.Printf("%d %d\n", N, X)

	xt := X * 100

	allAlc := 0
	ans := 0
	for i := 1; i <= N; i++ {
		vi := strings.Split(readLine(), " ")
		V := strToInt(vi[0])
		P := strToInt(vi[1])
		//fmt.Printf("%d %d\n", V, P)

		vt := V * 100
		pt := P * 100
		allAlc += vt * pt / 10000
		//fmt.Printf("allAlc %d\n", allAlc)

		if xt < allAlc {
			ans = i
			break
		}
	}
	if ans > 0 {
		fmt.Printf("%d\n", ans)
	} else {
		fmt.Printf("%d\n", -1)
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
