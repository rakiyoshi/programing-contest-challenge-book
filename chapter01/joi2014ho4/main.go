package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

type Path struct {
	to, time int
}

func main() {
	n := nextInt()
	m := nextInt()
	x := nextInt()
	h := nextInts(n)
	edges := make(map[int][]Path)
	for i := 0; i < m; i++ {
		a := nextInt() - 1
		b := nextInt() - 1
		t := nextInt()
		edges[a] = append(edges[a], Path{b, t})
		edges[b] = append(edges[b], Path{a, t})
	}

	pq := PriorityQueue{&Cost{0, 0, 0}}
	costs := make([]Cost, n)
	checked := make([]bool, n)
	for i := 1; i < n; i++ {
		costs[i] = Cost{i, i, math.MaxInt64}
		heap.Push(&pq, &costs[i])
	}
	for len(pq) > 0 {
		current := heap.Pop(&pq).(*Cost)
		checked[current.costIdx] = true
		if current.value == math.MaxInt64 {
			continue
		}
		tall := max(0, x-costs[current.costIdx].value)
		for _, next := range edges[current.costIdx] {
			if checked[next.to] || h[current.costIdx]-next.time < 0 {
				continue
			}
			var nextCost int
			if h[next.to] < tall-next.time { // go down
				nextCost = current.value + tall - h[next.to]
			} else if tall-next.time < 0 { // go up
				nextCost = current.value + next.time + (next.time - tall)
			} else {
				nextCost = current.value + next.time
			}
			if nextCost < costs[next.to].value {
				pq.update(&costs[next.to], nextCost)
			}
		}
	}
	// fmt.Printf("%+v\n", costs)
	if costs[n-1].value == math.MaxInt64 {
		fmt.Println(-1)
	} else {
		if x <= costs[n-1].value {
			fmt.Println(costs[n-1].value + h[n-1])
		} else {
			fmt.Println(h[n-1] - (x - costs[n-1].value) + costs[n-1].value)
		}
	}
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

type Cost struct {
	heapIdx int
	costIdx int
	value   int
}

type PriorityQueue []*Cost

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].value < pq[j].value
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].heapIdx = i
	pq[j].heapIdx = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Cost)
	item.heapIdx = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*pq = old[:n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Cost, cost int) {
	item.value = cost
	heap.Fix(pq, item.heapIdx)
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
