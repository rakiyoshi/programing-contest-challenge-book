package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

const mod = 1000000007

func main() {
	d := nextInt()
	n := next()

	dp := make([][][2]int, len(n))
	for i := 0; i < len(n); i++ {
		dp[i] = make([][2]int, d)
	}
	var prev int
	for i := 0; i < len(n); i++ {
		a := int(n[i] - '0')

		if i == 0 {
			dp[i][a%d][0] = 1
			for j := a - 1; j >= 0; j-- {
				dp[i][j%d][1] += 1
			}
			prev = a % d
			continue
		}

		dp[i][(a+prev)%d][0] = 1
		for j := 0; j < a; j++ {
			dp[i][(prev+j)%d][1] += 1
			dp[i][(prev+j)%d][1] %= mod
		}
		for j := 0; j < d; j++ {
			for k := 0; k < 10; k++ {
				dp[i][(j+k)%d][1] += dp[i-1][j][1]
				dp[i][(j+k)%d][1] %= mod
			}
		}

		prev = (prev + a) % d
	}
	// for _, v := range dp {
	// 	fmt.Println(v)
	// }
	fmt.Println(dp[len(n)-1][0][0] + dp[len(n)-1][0][1] - 1)
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
