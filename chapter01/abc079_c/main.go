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
	a := n / 1000
	b := n / 100 % 10
	c := n / 10 % 10
	d := n % 10

	for i := 0; i < 8; i++ {
		res := a
		var op1, op2, op3 string
		if i&(1<<0) == 0 {
			res += b
			op1 = "+"
		} else {
			res -= b
			op1 = "-"
		}
		if i&(1<<1) == 0 {
			res += c
			op2 = "+"
		} else {
			res -= c
			op2 = "-"
		}
		if i&(1<<2) == 0 {
			res += d
			op3 = "+"
		} else {
			res -= d
			op3 = "-"
		}
		if res != 7 {
			continue
		}
		fmt.Printf("%d%s%d%s%d%s%d=7\n", a, op1, b, op2, c, op3, d)
		return
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
