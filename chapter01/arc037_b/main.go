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
	edges := make(map[int][]int)
	for i := 0; i < m; i++ {
		u := nextInt() - 1
		v := nextInt() - 1
		edges[u] = append(edges[u], v)
		edges[v] = append(edges[v], u)
	}
	var count int
	visited := make([]bool, n)
	for i := 0; i < n; i++ {
		if visited[i] {
			continue
		}
		if dfs(i, -1, edges, visited) {
			count++
		}
	}
	fmt.Println(count)
}

func dfs(start, prev int, edges map[int][]int, visited []bool) bool {
	visited[start] = true
	res := true
	for _, next := range edges[start] {
		if next == prev {
			continue
		}
		if visited[next] {
			res = false
			continue
		}
		res = res && dfs(next, start, edges, visited)
	}
	return res
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
