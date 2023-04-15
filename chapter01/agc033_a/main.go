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
	var queue []Point
	s := make([][]byte, h)
	for i := 0; i < h; i++ {
		s[i] = []byte(next())
		for j := 0; j < w; j++ {
			if s[i][j] == '#' {
				queue = append(queue, Point{i, j})
			}
		}
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

	var ans int
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		for _, delta := range deltas {
			x := current.x + delta.x
			y := current.y + delta.y
			if x < 0 || h <= x || y < 0 || w <= y || s[x][y] == '#' || (distance[x][y] != 0 && distance[x][y] <= distance[current.x][current.y]+1) {
				continue
			}
			distance[x][y] = distance[current.x][current.y] + 1
			s[x][y] = '#'
			ans = distance[x][y]
			queue = append(queue, Point{x, y})
		}
	}
	fmt.Println(ans)
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
