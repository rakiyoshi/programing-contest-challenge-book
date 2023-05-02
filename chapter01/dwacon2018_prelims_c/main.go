package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

const mod = int(1e9 + 7)

func main() {
	n := nextInt()
	m := nextInt()
	a := nextInts(n)
	b := nextInts(m)

	dp := make([][]int, 1001)
	for i := range dp {
		dp[i] = make([]int, 101)
	}
	dp[0][0] = 1
	for n := 0; n <= 1000; n++ {
		for k := 1; k <= 100; k++ {
			if n-k >= 0 {
				dp[n][k] = dp[n-k][k]
			}
			dp[n][k] = (dp[n][k] + dp[n][k-1]) % mod
		}
	}

	ansA := solve(a, sum(b), dp)
	ansB := solve(b, sum(a), dp)
	// fmt.Println(ansA, ansB)
	fmt.Println(ansA * ansB % mod)
}

func solve(kills []int, deaths int, dpg [][]int) int {
	var group []int
	for i := range kills {
		if i == 0 || kills[i] != kills[i-1] {
			group = append(group, 1)
		} else {
			group[len(group)-1]++
		}
	}

	dp := make([][]int, len(group))
	for i := range dp {
		dp[i] = make([]int, deaths+1)
	}
	for g := 0; g < len(group); g++ {
		for k := 0; k <= deaths; k++ {
			if g == 0 {
				dp[g][k] = dpg[k][group[g]]
				// fmt.Printf("dpg[%d][%d]=%d\n", k, group[g], dpg[k][group[g]])
				continue
			}
			for i := 0; i <= k; i++ {
				dp[g][k] += dp[g-1][k-i] * dpg[i][group[g]]
				dp[g][k] %= mod
			}
		}
	}

	// for _, v := range dp {
	// 	fmt.Println(v)
	// }
	return dp[len(group)-1][deaths]
}

func sum(a []int) int {
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
