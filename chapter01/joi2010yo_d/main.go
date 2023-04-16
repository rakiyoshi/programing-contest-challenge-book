package main

import (
	"bufio"
	"fmt"
	"math"
	"math/bits"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	n := nextInt()
	k := nextInt()
	a := nextInts(n)
	pat := make(map[string]struct{})
	combs := combinations(a, k)
	for _, comb := range combs {
		sets := permute(comb)
		for _, set := range sets {
			var s string
			for _, v := range set {
				s = s + fmt.Sprint(v)
			}
			pat[s] = struct{}{}
		}
	}
	fmt.Println(len(pat))
}

func combinations(sets []int, r int) (subsets [][]int) {
	n := len(sets)
	if r > n {
		fmt.Fprintf(os.Stderr, "%v", fmt.Errorf("len(sets) must be less than r: len(sets)=%d, r=%d", len(sets), r))
		os.Exit(1)
	}
	for subsetBits := uint(1); subsetBits < (1 << len(sets)); subsetBits++ {
		if bits.OnesCount(subsetBits) != r {
			continue
		}
		var subset []int
		for i := range sets {
			if (subsetBits>>i)&1 == 1 {
				subset = append(subset, sets[i])
			}
		}
		subsets = append(subsets, subset)
	}
	return
}

func permute(a_ []int) [][]int {
	a := makeCopy(a_)
	result := make([][]int, 0, factorial(len(a)))
	result = append(result, makeCopy(a))

	n := len(a)
	p := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		p[i] = i
	}
	for i := 1; i < n; {
		p[i]--
		var j int
		if i%2 == 1 {
			j = p[i]
		}
		a[j], a[i] = a[i], a[j]
		result = append(result, makeCopy(a))
		i = 1
		for p[i] == 0 {
			p[i] = i
			i++
		}
	}
	return result
}

func factorial(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

func makeCopy(a []int) []int {
	tmp := make([]int, len(a))
	copy(tmp, a)
	return tmp
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
