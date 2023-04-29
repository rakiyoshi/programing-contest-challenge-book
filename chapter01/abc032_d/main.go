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

func main() {
	n := nextInt()
	m := nextInt()
	v := make([]int, n)
	w := make([]int, n)
	for i := 0; i < n; i++ {
		v[i] = nextInt()
		w[i] = nextInt()
	}

	if n <= 30 {
		fmt.Println(solveA(n, m, v, w))
	} else if ltAll(1001, w) {
		fmt.Println(solveB(n, m, v, w))
	} else {
		fmt.Println(solveC(n, m, v, w))
	}
}

func ltAll(n int, a []int) bool {
	for _, v := range a {
		if n <= v {
			return false
		}
	}
	return true
}

func solveA(n, m int, v, w []int) int {
	type Item struct {
		v int
		w int
	}
	if len(v) == 1 {
		return v[0]
	}
	var itemsA, itemsB []Item
	for i := 0; i < pow(2, n/2); i++ {
		var value, weight int
		for j := 0; j < n/2; j++ {
			if i&(1<<j) != 0 {
				value += v[j]
				weight += w[j]
			}
		}
		if weight <= m {
			itemsA = append(itemsA, Item{value, weight})
		}
	}
	for i := 0; i < pow(2, n-n/2); i++ {
		var value, weight int
		for j := n / 2; j < n; j++ {
			if i&(1<<(j-n/2)) != 0 {
				value += v[j]
				weight += w[j]
			}
		}
		if weight != 0 && weight <= m {
			itemsB = append(itemsB, Item{value, weight})
		}
	}

	sort.Slice(itemsA, func(i, j int) bool {
		if itemsA[i].w != itemsA[j].w {
			return itemsA[i].w < itemsA[j].w
		}
		return itemsA[i].v > itemsA[j].v
	})

	tmp := -1
	for i := 0; i < len(itemsA); i++ {
		for i < len(itemsA) && itemsA[i].v <= tmp {
			copy(itemsA[i:], itemsA[i+1:])
			itemsA = itemsA[:len(itemsA)-1]
		}
		if i != len(itemsA) {
			tmp = itemsA[i].v
		}
	}
	sort.Slice(itemsB, func(i, j int) bool {
		if itemsB[i].w != itemsB[j].w {
			return itemsB[i].w < itemsB[j].w
		}
		return itemsB[i].v > itemsB[j].v
	})

	tmp = -1
	for i := 0; i < len(itemsB); i++ {
		for i < len(itemsB) && itemsB[i].v <= tmp {
			copy(itemsB[i:], itemsB[i+1:])
			itemsB = itemsB[:len(itemsB)-1]
		}
		if i != len(itemsB) {
			tmp = itemsB[i].v
		}
	}

	var ans int
	for _, item := range itemsA {
		ans = max(ans, item.v)
	}
	for _, item := range itemsB {
		ans = max(ans, item.v)
	}
	for _, item := range itemsB {
		idx := sort.Search(len(itemsA), func(i int) bool {
			return m-itemsA[i].w < item.w
		})
		if idx == 0 {
			continue
		}
		idx--
		if itemsA[idx].w+item.w <= m {
			ans = max(ans, itemsA[idx].v+item.v)
		}
	}
	return ans
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

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func solveB(n, m int, v, w []int) int {
	dp := make(map[int]int)
	dp[0] = 0
	for i := 0; i < n; i++ {
		next := make(map[int]int)
		for weight, value := range dp {
			if _, ok := next[weight]; (!ok || next[weight] < value) && weight <= m {
				next[weight] = value
			}
			if _, ok := next[weight+w[i]]; (!ok || next[weight+w[i]] < value+v[i]) && weight+w[i] <= m {
				next[weight+w[i]] = value + v[i]
			}
		}
		dp = next
	}
	var ans int
	for _, v := range dp {
		if ans < v {
			ans = v
		}
	}
	return ans
}

func solveC(n, m int, v, w []int) int {
	dp := make(map[int]int)
	dp[0] = 0
	for i := 0; i < n; i++ {
		next := make(map[int]int)
		for value, weight := range dp {
			if _, ok := next[value]; (!ok || weight < next[value]) && weight <= m {
				next[value] = weight
			}
			if _, ok := next[value+v[i]]; (!ok || weight+w[i] < next[value+v[i]]) && weight+w[i] <= m {
				next[value+v[i]] = weight + w[i]
			}
		}
		dp = next
	}
	var ans int
	for v := range dp {
		if ans < v {
			ans = v
		}
	}
	return ans
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
