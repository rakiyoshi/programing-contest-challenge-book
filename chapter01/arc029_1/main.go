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
	t := nextInts(n)
	var ans int
	switch n {
	case 1:
		ans = t[0]
	case 2:
		ans = max(t[0], t[1])
	case 3:
		ans = min(
			max(t[0]+t[1], t[2]),
			max(t[1]+t[2], t[0]),
			max(t[2]+t[0], t[1]),
		)
	case 4:
		ans = math.MaxInt64
		for i := 0; i < 16; i++ {
			var left, right int
			for j := 0; j < 4; j++ {
				if i&(1<<j) != 0 {
					left += t[j]
				} else {
					right += t[j]
				}
			}
			ans = min(ans, max(right, left))
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

func min(params ...int) int {
	res := math.MaxInt64
	for _, v := range params {
		if v < res {
			res = v
		}
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
