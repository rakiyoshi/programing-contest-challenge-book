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
	h := nextInt()
	w := nextInt()
	s := Point{
		nextInt() - 1,
		nextInt() - 1,
	}
	g := Point{
		nextInt() - 1,
		nextInt() - 1,
	}
	maze := make([][]byte, h)
	for i := 0; i < h; i++ {
		maze[i] = []byte(next())
	}
	deltas := []Point{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}

	distance := make([][]int, h)
	for i := 0; i < h; i++ {
		distance[i] = make([]int, w)
	}

	queue := []Point{s}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		for _, delta := range deltas {
			x := current.x + delta.x
			y := current.y + delta.y
			if x < 0 || h <= x || y < 0 || w <= y || maze[x][y] == '#' || (distance[x][y] != 0 && distance[x][y] <= distance[current.x][current.y]+1) {
				continue
			}
			distance[x][y] = distance[current.x][current.y] + 1
			queue = append(queue, Point{x, y})
		}
	}
	fmt.Println(distance[g.x][g.y])
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

func nextInts(n int) []int {
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = nextInt()
	}
	return a
}
