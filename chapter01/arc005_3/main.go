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
	x, y  int
	count int
}

func main() {
	h := nextInt()
	w := nextInt()
	c := make([]string, h)
	var s, g Point
	for i := 0; i < h; i++ {
		c[i] = next()
		for j := 0; j < w; j++ {
			switch c[i][j] {
			case 's':
				s = Point{i, j, 0}
			case 'g':
				g = Point{i, j, 0}
			}
		}
	}

	deltas := []Point{
		{0, 1, 0},
		{1, 0, 0},
		{0, -1, 0},
		{-1, 0, 0},
	}
	counts := make([][]int, h)
	for i := 0; i < h; i++ {
		counts[i] = make([]int, w)
		for j := 0; j < w; j++ {
			counts[i][j] = -1
		}
	}
	queue := []Point{s}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		for _, delta := range deltas {
			x := current.x + delta.x
			y := current.y + delta.y
			count := current.count
			if x < 0 || h <= x || y < 0 || w <= y {
				continue
			}
			if c[x][y] == '#' {
				count++
			}
			if count > 2 ||
				(counts[x][y] != -1 && counts[x][y] <= count) {
				continue
			}
			counts[x][y] = count
			queue = append(queue, Point{x, y, count})
		}
	}

	if counts[g.x][g.y] == -1 {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
	}
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
