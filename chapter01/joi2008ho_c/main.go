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
	m := nextInt()
	p := nextInts(n)
	p = append(p, 0)

	// s1 + s2 + s3 + s4 <= m
	// s1 + s2 <= m - s3 - s4
	var left []int
	for i := 0; i < n+1; i++ {
		for j := i; j < n+1; j++ {
			left = append(left, p[i]+p[j])
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(left)))

	var ans int
	for i := 0; i < n+1; i++ {
		for j := i; j < n+1; j++ {
			a := sort.Search(len(left), func(l int) bool {
				return left[l] <= m-p[i]-p[j]
			})
			if a == len(left) {
				continue
			} else {
				ans = max(ans, left[a]+p[i]+p[j])
			}
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
