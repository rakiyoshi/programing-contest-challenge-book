package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

type Circle struct {
	l, r int
}

func main() {
	n := nextInt()
	circles := make([]Circle, n)
	for i := range circles {
		x := nextInt()
		r := nextInt()
		circles[i] = Circle{x - r, x + r}
	}
	sort.Slice(circles, func(i, j int) bool {
		if circles[i].r != circles[j].r {
			return circles[i].r < circles[j].r
		}
		return circles[i].l < circles[j].l
	})

	dp := make([]int, n+2)
	for i := 1; i <= n+1; i++ {
		dp[i] = math.MinInt64
	}
	dp[0] = math.MaxInt64
	for i := 0; i < n; i++ {
		idx := sort.Search(n+2, func(j int) bool {
			return dp[j] <= circles[i].l
		})
		dp[idx] = circles[i].l
	}
	for i := 1; i <= n+1; i++ {
		if dp[i] == math.MinInt64 {
			fmt.Println(i - 1)
			return
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

	buf := make([]byte, 1024*1024)
	sc.Buffer(buf, math.MaxInt64)
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
