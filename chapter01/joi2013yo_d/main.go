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
	d := nextInt() // 2 <= d <= 200
	n := nextInt() // 1 <= n <= 200
	t := nextInts(d)
	a := make([]int, n)
	b := make([]int, n)
	c := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = nextInt()
		b[i] = nextInt()
		c[i] = nextInt()
	}

	dp := make([][]int, d)
	for i := 0; i < d; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < d; i++ {
		for j := 0; j < n; j++ {
			if i == 0 {
				dp[i][j] = 0
				continue
			}
			if t[i] < a[j] || b[j] < t[i] {
				continue
			}
			for k := 0; k < n; k++ {
				if t[i-1] < a[k] || b[k] < t[i-1] {
					continue
				}
				dp[i][j] = max(dp[i][j], dp[i-1][k]+abs(c[j]-c[k]))
			}
		}
	}

	var ans int
	for _, v := range dp[d-1] {
		ans = max(ans, v)
	}

	// for _, v := range dp {
	// 	fmt.Println(v)
	// }
	fmt.Println(ans)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
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
