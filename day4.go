package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Combo struct {
	Red   int
	Green int
	Blue  int
}

var Constraint Combo = Combo{12, 13, 14}

func checkTicket(line string) float64 {
	x := strings.Split(line, ":")
	var id int
	fmt.Sscanf(x[0], "Game %d", &id)
	xn := strings.Split(x[1], "|")
	winstrs := strings.Fields(xn[0])
	winners := make([]int, len(winstrs))
	for i, s := range winstrs {
		winners[i], _ = strconv.Atoi(s)
	}
	havestrs := strings.Fields(xn[1])
	havers := make([]int, len(havestrs))
	for i, s := range havestrs {
		havers[i], _ = strconv.Atoi(s)
	}

	fmt.Println(winners, havers)

	count := 0
	for _, x := range havers {
		if slices.Contains(winners, x) {
			count++
		}
	}
	if count > 0 {
		return math.Pow(2, float64(count-1))
	} else {
		return 0
	}
	// fmt.Println(count)
	// return count
}

func main() {

	var total float64 = 0
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		total += checkTicket(line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}

	fmt.Println(total)
}
