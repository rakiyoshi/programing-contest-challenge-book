package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	s := next()
	t := next()
	if len(s) < len(t) {
		fmt.Println("UNRESTORABLE")
		return
	}
	for i := len(s) - len(t); i >= 0; i-- {
		for j := 0; j < len(t); j++ {
			if s[i+j] != '?' && s[i+j] != t[j] {
				break
			}
			if j == len(t)-1 {
				for k := 0; k < len(s); k++ {
					if i <= k && k < i+len(t) {
						fmt.Print(string(t[k-i]))
					} else {
						if s[k] == '?' {
							fmt.Print("a")
						} else {
							fmt.Print(string(s[k]))
						}
					}
				}
				fmt.Println()
				return
			}
		}
	}
	fmt.Println("UNRESTORABLE")
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
