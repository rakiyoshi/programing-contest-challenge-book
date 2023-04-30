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
	k := nextInt()
	a := nextInts(n)

	if k == 0 {
		fmt.Println(0)
		return
	}
	if k == sum(a...) || n == 1 {
		fmt.Println(1)
		return
	}

	sumA := a[0]
	if a[0] == 1 {
		sumA += a[1]
		a = a[1:]
		n--
	}
	b := make([]int, n+1)
	for i := 0; i <= n; i++ {
		b[i] = i
	}

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, i+2)
		for j := 1; j < i+2; j++ {
			dp[i][j] = 500000*2000 + 1
		}
	}
	dp[0][1] = 1

	for i := 1; i < n; i++ {
		for j := 1; j <= i+1; j++ {
			if j != i+1 {
				dp[i][j] = dp[i-1][j]
			}
			// dp[i-1][j-1]/sum(a[:i-1]) < dp[i][j]/sum(a[:i])
			// dp[i-1][j-1] * sum(a[:i]) < dp[i][j] * sum(a[:i-1])
			wins := dp[i-1][j-1]*(sumA+a[i])/sumA + 1
			dp[i][j] = min(dp[i][j], wins)
		}
		sumA += a[i]
	}

	var ans int
	for i := 1; i <= n; i++ {
		if dp[n-1][i] <= k {
			ans = i
		}
	}

	// fmt.Println(dp)
	fmt.Println(ans)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func sum(a ...int) int {
	var res int
	for _, v := range a {
		res += v
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
