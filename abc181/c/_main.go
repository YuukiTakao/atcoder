package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// ここみながらやった
// https://www.geisya.or.jp/~mwm48961/kou2/line_brief1.htm
func main() {
	N := strToInt(readLine())

	//fmt.Printf("%d\n", N)

	points := make([][2]int, N)
	for i := 0; i < N; i++ {
		pi := strings.Split(readLine(), " ")
		x := strToInt(pi[0])
		y := strToInt(pi[1])
		//fmt.Printf("%d %d\n", x, y)
		points[i] = [2]int{x, y}
	}
	//fmt.Printf("points %v\n", points)

	ans := false
	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			for k := j + 1; k < N; k++ {
				x1 := points[i][0]
				y1 := points[i][1]
				//fmt.Printf("x1 %d y1 %d\n", x1, y1)

				x2 := points[j][0]
				y2 := points[j][1]
				//fmt.Printf("x2 %d y2 %d\n", x2, y2)

				x3 := points[k][0]
				y3 := points[k][1]
				//fmt.Printf("x3 %d y3 %d\n", x3, y3)

				if (y3-y1)*(x2-x1) == (y2-y1)*(x3-x1) {
					ans = true
				}
			}
		}
	}
	if ans {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
	//println(ans)
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
