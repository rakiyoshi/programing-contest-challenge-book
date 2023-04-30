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
	d := nextInt()
	pfs := getPrimeFactorsMap(d)
	var two, three, five int
	for k, v := range pfs {
		switch k {
		case 2:
			two = v
		case 3:
			three = v
		case 5:
			five = v
		default:
			fmt.Println(0)
			return
		}
	}
	dp := make([][][][]float64, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([][][]float64, two+1)
		for j := 0; j <= two; j++ {
			dp[i][j] = make([][]float64, three+1)
			for k := 0; k <= three; k++ {
				dp[i][j][k] = make([]float64, five+1)
			}
		}
	}

	dp[0][0][0][0] = 1
	for i := 0; i < n; i++ {
		for j := 0; j <= two; j++ {
			for k := 0; k <= three; k++ {
				for l := 0; l <= five; l++ {
					dp[i+1][j][k][l] += dp[i][j][k][l]                           // 1
					dp[i+1][min(j+1, two)][k][l] += dp[i][j][k][l]               // 2
					dp[i+1][j][min(k+1, three)][l] += dp[i][j][k][l]             // 3
					dp[i+1][min(j+2, two)][k][l] += dp[i][j][k][l]               // 4
					dp[i+1][j][k][min(l+1, five)] += dp[i][j][k][l]              // 5
					dp[i+1][min(j+1, two)][min(k+1, three)][l] += dp[i][j][k][l] // 6
				}
			}
		}
	}
	fmt.Println(dp[n][two][three][five] / math.Pow(6, float64(n)))
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
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
