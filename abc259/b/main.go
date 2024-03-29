package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func AngleToRadian(angle float64) float64 {
	return angle * (math.Pi / 180)
}
func RadianToAngle(radian float64) float64 {
	return radian * (180 / math.Pi)
}

// 2次元座標のユークリッド距離を求める
func EuclideanDistance(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}

func main() {
	bufInit()
	defer wr.Flush()
	a := scanInt()
	b := scanInt()
	d := scanInt()

	eDistance := EuclideanDistance(0, 0, float64(a), float64(b))
	sumRad := AngleToRadian(float64(d)) + math.Atan2(float64(b), float64(a))

	newX := math.Cos(sumRad) * eDistance
	newY := math.Sin(sumRad) * eDistance

	fprintf("%.20f %.20f\n", newX, newY)
}

var wr *bufio.Writer
var sc = bufio.NewScanner(os.Stdin)

func fprintf(f string, a ...interface{}) {
	fmt.Fprintf(wr, f, a...)
}
func fprintln(a ...interface{}) {
	fmt.Fprintln(wr, a...)
}
func fprint(f string, a ...interface{}) {
	fmt.Fprint(wr, a...)
}
func bufInit() {
	sc.Buffer(make([]byte, 128), 500000)
	sc.Split(bufio.ScanWords)
	wr = bufio.NewWriter(os.Stdout)
}
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
