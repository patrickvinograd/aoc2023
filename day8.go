package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Node struct {
	Left  string
	Right string
}

var desert = make(map[string]Node)

func parseTurns(line string) []string {
	return strings.Split(line, "")
}

func parseMapLine(line string) {
	var key, left, right string
	fmt.Sscanf(line, "%3s = (%3s, %3s)", &key, &left, &right)
	desert[key] = Node{left, right}
}

func walk(turns []string) int {
	steps := 0
	current := "AAA"
	for i := 0; true; i++ {
		turn := turns[i%len(turns)]
		var next string
		if turn == "L" {
			next = desert[current].Left
		} else {
			next = desert[current].Right
		}
		fmt.Println(current, next)
		current = next
		steps++
		if current == "ZZZ" {
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
	for scanner.Scan() {
		line := scanner.Text()
		parseMapLine(line)
	}
	fmt.Println(turns)
	fmt.Println(desert)

	fmt.Println(walk(turns))
	// fmt.Println(total)
}
