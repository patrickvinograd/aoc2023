package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	wliterals := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	nliterals := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	var total int = 0
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		var firstindex int = len(line)
		var firstval int = 0
		for i, lit := range wliterals {
			fi := strings.Index(line, lit)
			if fi >= 0 && fi <= firstindex {
				firstindex = fi
				firstval = i
			}
		}
		for i, lit := range nliterals {
			fi := strings.Index(line, lit)
			if fi >= 0 && fi <= firstindex {
				firstindex = fi
				firstval = i
			}
		}

		var lastindex int = 0
		var lastval int = 0
		for i, lit := range wliterals {
			li := strings.LastIndex(line, lit)
			if li >= lastindex {
				lastindex = li
				lastval = i
			}
		}
		for i, lit := range nliterals {
			li := strings.LastIndex(line, lit)
			if li >= lastindex {
				lastindex = li
				lastval = i
			}
		}
		//fmt.Fprintln(os.Stdout, firstval, firstindex, lastval, lastindex)
		total += (firstval * 10) + lastval
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}

	fmt.Println(total)
}
