package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	s := next()

	// pattern: 2 ^ (len(s)-1)
	end := pow(2, len(s)-1)
	var ans int
	for pat := 0; pat < end; pat++ {
		formula := []byte{s[0]}
		for i := 1; i < len(s); i++ {
			if (pat & (1 << (i - 1))) != 0 {
				formula = append(formula, '+')
			}
			formula = append(formula, s[i])
		}
		params := strings.Split(string(formula), "+")
		for _, p := range params {
			n, err := strconv.ParseInt(p, 10, 64)
			if err != nil {
				log.Fatalf("error: %s", p)
			}
			ans += int(n)
		}
	}
	fmt.Println(ans)
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
