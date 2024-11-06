package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Record struct {
	condition string
	unknowns  int
	sizeStr   string
	sizes     []int
}

func parseLine(line string) Record {
	x := strings.Split(line, " ")

	conditionStr := x[0]
	sizeStr := x[1]
	for i := 0; i < 4; i++ {
		conditionStr = conditionStr + "?" + x[0]
		sizeStr = sizeStr + "," + x[1]
	}
	vstrs := strings.Split(sizeStr, ",")
	sizes := make([]int, len(vstrs))
	for i, s := range vstrs {
		sizes[i], _ = strconv.Atoi(s)
	}
	unknowns := 0
	for _, c := range conditionStr {
		if c == 63 { // ?
			unknowns++
		}
	}

	return Record{conditionStr, unknowns, sizeStr, sizes}
}

func sizeString(conditionStr string) string {
	runs := make([]string, 0)
	inPart := false
	runLength := 0
	for i := 0; i < len(conditionStr); i++ {
		if conditionStr[i] == 35 { // #
			if inPart {
				runLength++
			} else {
				inPart = true
				runLength++
			}
		} else { // .
			if inPart {
				runs = append(runs, strconv.Itoa(runLength))
				runLength = 0
				inPart = false
			} else {
				inPart = false
			}
		}
	}
	if inPart {
		runs = append(runs, strconv.Itoa(runLength))
	}
	// fmt.Println(conditionStr, strings.Join(runs, ","))
	return strings.Join(runs, ",")
}

func permuteString(condition []byte, permutation int) string {
	ukIndex := 0
	for i := 0; i < len(condition); i++ {
		if condition[i] == 63 {
			if permutation&(0x01<<ukIndex) != 0 {
				condition[i] = 35 // #
			} else {
				condition[i] = 46 // .
			}
			ukIndex++
		}
	}
	return string(condition)
}

func permuteCount(data Record) int {
	matches := 0
	n := math.Pow(2, float64(data.unknowns))
	for i := 0; i < int(n); i++ {
		variant := permuteString([]byte(data.condition), i)
		if sizeString(variant) == data.sizeStr {
			matches++
			// fmt.Println(i, variant)
		}
		// fmt.Println(i, variant)
	}
	return matches
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	total := 0
	for scanner.Scan() {
		data := parseLine(scanner.Text())
		fmt.Println(data)
		fmt.Println(len(data.condition))
		// total += permuteCount(data)
	}
	fmt.Println(total)
}
