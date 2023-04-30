package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

type Point struct {
	x, y int
}

func main() {
	n := nextInt() // 2 <= n <= 100
	points := make([]Point, n)
	for i := 0; i < n; i++ {
		points[i] = Point{nextInt(), nextInt()}
	}

	var ans float64
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			ans = max(ans, distance(points[i], points[j]))
		}
	}

	fmt.Println(ans)
}

func max(x, y float64) float64 {
	if x < y {
		return y
	}
	return x
}

func distance(a, b Point) float64 {
	return math.Hypot(float64(a.x-b.x), float64(a.y-b.y))
}

func init() {
	if len(os.Args) > 1 && os.Args[1] == "debug" {
		if len(os.Args) == 2 {
			fmt.Fprintf(os.Stderr, "filename is required")
			os.Exit(1)
		}
		debug(os.Args[2])
	}

	sc.Split(bufio.ScanWords)

	buf := make([]byte, 10*1024)
	sc.Buffer(buf, math.MaxInt32)
}

func debug(filename string) {
	testFile, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: no such file", filename)
		os.Exit(1)
	}
	sc = bufio.NewScanner(testFile)
}

func next() string {
	sc.Scan()
	return sc.Text()
}

func nextInt() int {
	n, err := strconv.Atoi(next())
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}
	return n
}
