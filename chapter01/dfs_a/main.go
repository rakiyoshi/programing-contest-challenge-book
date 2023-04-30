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
	start := Point{}
	goal := Point{}
	c := make([][]byte, h)
	for i := 0; i < h; i++ {
		c[i] = make([]byte, w)
		s := next()
		for j := 0; j < w; j++ {
			c[i][j] = s[j]
			if c[i][j] == 's' {
				start = Point{i, j}
			} else if c[i][j] == 'g' {
				goal = Point{i, j}
			}
		}
	}

	visited := make([][]bool, h)
	for i := 0; i < h; i++ {
		visited[i] = make([]bool, w)
	}
	visited[start.x][start.y] = true
	if dfs(start, goal, c, h, w, &visited) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func dfs(start, goal Point, c [][]byte, h, w int, visited *[][]bool) bool {
	deltas := []Point{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}

	for _, delta := range deltas {
		x := start.x + delta.x
		y := start.y + delta.y
		if goal.x == x && goal.y == y {
			return true
		}
		if x < 0 || h <= x || y < 0 || w <= y || c[x][y] == '#' || (*visited)[x][y] {
			continue
		}
		(*visited)[x][y] = true
		if dfs(Point{x, y}, goal, c, h, w, visited) {
			return true
		}
	}
	return false
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
