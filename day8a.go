package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

// inlined from https://github.com/TheAlgorithms/Go/tree/master/math/gcd
func iterative(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// inlined from https://github.com/TheAlgorithms/Go/tree/master/math/lcm
func lcm(a, b int64) int64 {
	return int64(math.Abs(float64(a*b)) / float64(iterative(a, b)))
}

type Node struct {
	Left  string
	Right string
}

var desert = make(map[string]Node)

func parseTurns(line string) []string {
	return strings.Split(line, "")
}

func parseMapLine(line string) string {
	var key, left, right string
	fmt.Sscanf(line, "%3s = (%3s, %3s)", &key, &left, &right)
	desert[key] = Node{left, right}
	return key
}

func done(positions []string) bool {
	for _, p := range positions {
		if p[len(p)-1:] != "Z" {
			return false
		}
	}
	return true
}

func walk(start string, turns []string) int64 {
	current := start
	var steps int64 = 0
	for i := 0; true; i++ {
		turn := turns[i%len(turns)]
		if turn == "L" {
			current = desert[current].Left
		} else {
			current = desert[current].Right
		}
		steps++
		if current[len(current)-1:] == "Z" {
			return steps
		}
	}
	return 0
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	// line := scanner.Text()
	turns := parseTurns(scanner.Text())
	scanner.Scan() // skip blank line
	var starts []string
	for scanner.Scan() {
		line := scanner.Text()
		id := parseMapLine(line)
		if id[len(id)-1:] == "A" {
			starts = append(starts, id)
		}
	}
	fmt.Println(turns)
	fmt.Println(desert)
	fmt.Println(starts)
	var total int64 = 1
	for _, p := range starts {
		steps := walk(p, turns)
		total = lcm(total, steps)
		fmt.Println(p, steps)
	}
	fmt.Println(total)
}
