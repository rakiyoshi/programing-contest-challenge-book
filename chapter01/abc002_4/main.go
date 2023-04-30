package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	n := nextInt()
	m := nextInt()
	edges := make([][]bool, n)
	for i := 0; i < n; i++ {
		edges[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		x := nextInt() - 1
		y := nextInt() - 1
		edges[x][y] = true
	}
	var ans int
	end := pow(2, n)
	for i := 0; i < end; i++ {
		var candidate []int
		for j := 0; j < n; j++ {
			if i&(1<<j) != 0 {
				candidate = append(candidate, j)
			}
		}
		if isOk(candidate, edges) {
			ans = max(ans, len(candidate))
		}
	}
	fmt.Println(ans)
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func isOk(candidate []int, edges [][]bool) bool {
	for i := 0; i < len(candidate); i++ {
		for j := i + 1; j < len(candidate); j++ {
			if !edges[candidate[i]][candidate[j]] {
				return false
			}
		}
	}
	return true
}

func pow(x, n int) int {
	switch n {
	case 0:
		return 1
	case 1:
		return x
	default:
		return pow(x*x, n/2) * pow(x, n%2)
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
