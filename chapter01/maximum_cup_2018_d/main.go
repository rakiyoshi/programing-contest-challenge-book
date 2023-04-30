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
	l := nextInt()
	x := nextInt()
	a := nextInts(n)

	dp := map[int]int{0: 0}
	for i := 0; i < n; i++ {
		next := make(map[int]int)
		for place, laps := range dp {
			if l, ok := next[place]; !ok || laps < l {
				next[place] = laps
			}
			if x <= laps+(place+a[i])/m {
				continue
			}
			next[(place+a[i])%m] = laps + (place+a[i])/m
		}
		dp = next
	}

	if _, ok := dp[l]; ok {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
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
