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

func main() {
	n := nextInt()
	a := nextInts(n)
	dp := make([]int, n+2)
	for i := 1; i <= n+1; i++ {
		dp[i] = math.MaxInt64
	}
	dp[0] = 0
	dp[1] = a[0]
	for i := 1; i < n; i++ {
		idx := sort.Search(n+1, func(j int) bool {
			return a[i] <= dp[j]
		})
		if idx == 0 {
			continue
		}
		dp[idx] = a[i]
	}
	// fmt.Println(dp)
	for i := 1; i <= n+1; i++ {
		if dp[i] == math.MaxInt64 {
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
