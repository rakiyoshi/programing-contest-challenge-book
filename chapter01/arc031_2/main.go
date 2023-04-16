package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

type Point struct {
	x, y int
}

func main() {
	a := make([]string, 10)
	for i := 0; i < 10; i++ {
		a[i] = next()
	}
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if a[i][j] == 'x' {
				var visited [10][10]bool
				dfs(Point{i, j}, &visited, a)
				if check(a, visited) {
					fmt.Println("YES")
					return
				}
			}
		}
	}
	fmt.Println("NO")
}

func check(a []string, visited [10][10]bool) bool {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if a[i][j] == 'o' && !visited[i][j] {
				return false
			}
		}
	}
	return true
}

func dfs(start Point, visited *[10][10]bool, a []string) {
	deltas := []Point{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}
	for _, delta := range deltas {
		x := start.x + delta.x
		y := start.y + delta.y
		if x < 0 || 10 <= x || y < 0 || 10 <= y ||
			a[x][y] == 'x' ||
			(*visited)[x][y] {
			continue
		}
		(*visited)[x][y] = true
		if a[x][y] == 'x' {
			dfs(Point{x, y}, visited, a)
		} else {
			dfs(Point{x, y}, visited, a)
		}
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
