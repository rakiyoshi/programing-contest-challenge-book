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

type Cut struct {
	a, b int
}

func main() {
	nextInt()
	m := nextInt()
	cuts := make([]Cut, m)
	for i := 0; i < m; i++ {
		cuts[i] = Cut{nextInt(), nextInt()}
	}
	sort.Slice(cuts, func(i, j int) bool {
		return cuts[i].b < cuts[j].b
	})
	count := 1
	x := cuts[0].b
	for i := 1; i < m; i++ {
		if cuts[i].a < x {
			continue
		}
		x = cuts[i].b
		count++
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
