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
	d := nextInt()
	g := nextInt()
	g /= 100
	p := make([]int, d)
	c := make([]int, d)
	for i := 0; i < d; i++ {
		p[i] = nextInt()
		c[i] = nextInt() / 100
	}

	ans := math.MaxInt64
	for i := 0; i < pow(2, d); i++ {
		var count int
		var score int
		rem := d
		for j := 0; j < d; j++ {
			if i&(1<<j) != 0 {
				count += p[j]
				score += (j+1)*p[j] + c[j]
			} else {
				rem = j
			}
		}
		if g <= score {
			ans = min(ans, count)
			continue
		}
		if rem == d {
			continue
		}
		if score+(rem+1)*(p[rem]-1) < g {
			continue
		}
		ans = min(ans, count+(g-score+rem)/(rem+1))
	}
	fmt.Println(ans)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func pow(x, n int) int {
	switch n {
	case 0:
		return 1
	case 1:
		return x
	default:
		return pow(x*x, n/2) * pow(x, n%2)
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
