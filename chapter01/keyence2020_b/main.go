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

type Robot struct {
	x, l int
}

func main() {
	n := nextInt()
	robots := make([]Robot, n)
	for i := 0; i < n; i++ {
		robots[i] = Robot{nextInt(), nextInt()}
	}
	sort.Slice(robots, func(i, j int) bool {
		return robots[i].x+robots[i].l < robots[j].x+robots[j].l
	})

	count := 1
	x := robots[0].x + robots[0].l
	for i := 1; i < n; i++ {
		if robots[i].x-robots[i].l < x {
			continue
		}
		count++
		x = robots[i].x + robots[i].l
	}

	fmt.Println(count)
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
