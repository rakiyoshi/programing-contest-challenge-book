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
	n := nextInt()      // 1 <= n <= 50000
	m := nextInt()      // 1 <= m <= 50000
	a := make([]int, n) // 1 <= a[i] <= 100
	b := make([]int, n) // 1 <= b[i] <= 100
	c := make([]int, n) // 1 <= c[i] <= 100
	w := make([]int, n) // 1 <= c[i] <= 10**9
	for i := 0; i < n; i++ {
		a[i] = nextInt()
		b[i] = nextInt()
		c[i] = nextInt()
		w[i] = nextInt()
	}
	x := make([]int, m) // 1 <= x[i] <= 100
	y := make([]int, m) // 1 <= y[i] <= 100
	z := make([]int, m) // 1 <= z[i] <= 100
	for i := 0; i < m; i++ {
		x[i] = nextInt()
		y[i] = nextInt()
		z[i] = nextInt()
	}

	job := make([][][]int, 101)
	for i := 0; i <= 100; i++ {
		job[i] = make([][]int, 101)
		for j := 0; j <= 100; j++ {
			job[i][j] = make([]int, 101)
		}
	}
	for i := 0; i < n; i++ {
		job[a[i]][b[i]][c[i]] = max(job[a[i]][b[i]][c[i]], w[i])
	}
	for i := 0; i <= 100; i++ {
		for j := 0; j <= 100; j++ {
			for k := 0; k <= 100; k++ {
				job[i][j][k] = max(
					job[i][j][k],
					job[max(0, i-1)][j][k],
					job[i][max(0, j-1)][k],
					job[i][j][max(0, k-1)],
					job[max(0, i-1)][max(0, j-1)][k],
					job[max(0, i-1)][j][max(0, k-1)],
					job[i][max(0, j-1)][max(0, k-1)],
					job[max(0, i-1)][max(0, j-1)][max(0, k-1)],
				)
			}
		}
	}
	for i := 0; i < m; i++ {
		fmt.Println(job[x[i]][y[i]][z[i]])
	}
}

func max(a ...int) int {
	var res int
	for _, v := range a {
		if res < v {
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
