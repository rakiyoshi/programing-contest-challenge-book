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

	ans := 1
	pfs := getPrimeFactorsMap(m)
	for _, v := range pfs {
		ans *= solve(n, v)
		ans %= mod
	}
	fmt.Println(ans)
}

func solve(n, p int) int {
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, p+1)
	}
	for j := 0; j <= p; j++ {
		dp[0][j] = 1
	}
	for i := 0; i < n; i++ {
		dp[i][0] = 1
	}
	for i := 1; i < n; i++ {
		for j := 1; j <= p; j++ {
			if j-1-p >= 0 {
				dp[i][j] = dp[i][j-1] + dp[i-1][j] - dp[i-1][j-1-p]
			} else {
				dp[i][j] = dp[i][j-1] + dp[i-1][j]
			}
			dp[i][j] %= mod
		}
	}
	return dp[n-1][p]
}

func getPrimeFactorsMap(n int) (pfs map[int]int) {
	pfs = make(map[int]int)
	// Get the number of 2s that divide n
	for n%2 == 0 {
		pfs[2]++
		n = n / 2
	}

	// n must be odd at this point. so we can skip one element
	// (note i = i + 2)
	for i := 3; i*i <= n; i = i + 2 {
		// while i divides n, append i and divide n
		for n%i == 0 {
			pfs[i]++
			n = n / i
		}
	}

	// This condition is to handle the case when n is a prime number
	// greater than 2
	if n > 2 {
		pfs[n]++
	}

	return
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
