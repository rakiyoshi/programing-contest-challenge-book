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
	k := nextInt()

	fmt.Println(combination(n+k-1, k))
}

func combination(n, r int) int {
	if n-r < r {
		r = n - r
	}
	c := 1
	for i := 0; i < r; i++ {
		c *= n - i
		c %= mod
		c *= inv(i + 1)
		c %= mod
	}
	return c
}

func inv(a int) int {
	_, x, _ := extGCD(a, mod)
	if x < 0 {
		return x + mod
	}
	return x % mod
}

func extGCD(a, mod int) (int, int, int) {
	if mod == 0 {
		return a, 1, 0
	}
	d, x, y := extGCD(mod, a%mod)
	return d, y, x - (a/mod)*y
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
