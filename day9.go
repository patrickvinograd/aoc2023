package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Left  string
	Right string
}

var desert = make(map[string]Node)

func parseLine(line string) []int {
	vstrs := strings.Fields(line)
	result := make([]int, len(vstrs))
	for i, s := range vstrs {
		result[i], _ = strconv.Atoi(s)
	}
	return result
}

func allZeroes(data []int) bool {
	for _, x := range data {
		if x != 0 {
			return false
		}
	}
	return true
}

func calcLine(prev []int) []int {
	var result = make([]int, len(prev)-1)
	for i := 1; i < len(prev); i++ {
		result[i-1] = prev[i] - prev[i-1]
	}
	return result
}

func extrapolateAndScore(data [][]int) int {
	var total int = 0
	for _, row := range data {
		total += row[len(row)-1]
	}
	return total
}

func process(data []int) int {

	var steps [][]int
	steps = append(steps, data)
	for true {
		prev := steps[len(steps)-1]
		next := calcLine(prev)
		steps = append(steps, next)
		if allZeroes(next) {
			break
		}
	}
	//fmt.Println(steps)
	return extrapolateAndScore(steps)
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	total := 0
	for scanner.Scan() {
		data := parseLine(scanner.Text())
		total += process(data)
	}
	fmt.Println(total)
}
