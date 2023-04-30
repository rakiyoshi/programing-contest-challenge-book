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
	w := nextInt()
	n := nextInt()
	k := nextInt()
	a := make([]int, n)
	b := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = nextInt()
		b[i] = nextInt()
	}
	dp := make([][]map[int]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]map[int]int, k+1)
		for j := 0; j <= k; j++ {
			dp[i][j] = make(map[int]int)
		}
	}
	dp[0][0] = map[int]int{0: 0}
	for i := 0; i < n; i++ {
		for j := 0; j < k; j++ {
			for ww, v := range dp[i][j] {
				dp[i+1][j][ww] = max(dp[i+1][j][ww], v)
				if ww+a[i] <= w {
					dp[i+1][j+1][ww+a[i]] = max(dp[i+1][j+1][ww+a[i]], v+b[i])
				}
			}
		}
		for ww, v := range dp[i][k] {
			dp[i+1][k][ww] = max(dp[i+1][k][ww], v)
		}
	}
	var ans int
	for j := range dp[n] {
		for _, v := range dp[n][j] {
			ans = max(ans, v)
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
