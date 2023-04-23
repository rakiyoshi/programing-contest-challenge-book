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

type Point struct {
	x, y int
}

func main() {
	n := nextInt()
	a := make([]Point, n)
	b := make([]Point, n)
	for i := 0; i < n; i++ {
		a[i] = Point{nextInt(), nextInt()}
	}
	for i := 0; i < n; i++ {
		b[i] = Point{nextInt(), nextInt()}
	}
	sort.Slice(a, func(i, j int) bool {
		if a[i].x != a[j].x {
			return a[i].x < a[j].x
		}
		return a[i].y < a[j].y
	})
	sort.Slice(b, func(i, j int) bool {
		if b[i].x != b[j].x {
			return b[i].x < b[j].x
		}
		return b[i].y < b[j].y
	})

	var ans int
	for i := 0; i < n; i++ {
		// fmt.Println(a, b[i])

		idx := sort.Search(len(a), func(j int) bool {
			return a[j].x >= b[i].x
		})
		if idx == 0 {
			continue
		}
		maxi, maxv := -1, -1
		for j := 0; j < idx; j++ {
			if a[j].y < b[i].y {
				if maxv <= a[j].y {
					maxv = a[j].y
					maxi = j
				}
			}
		}
		if maxi == -1 {
			continue
		}

		ans++
		copy(a[maxi:], a[maxi+1:])
		a = a[:len(a)-1]
	}

	fmt.Println(ans)
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
