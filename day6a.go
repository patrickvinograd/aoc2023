package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// type Card struct {
// 	Red   int
// 	Green int
// 	Blue  int
// }

func parseLine(line string, field string) []int {
	vstrs := strings.Fields(strings.Split(line, ":")[1])
	result := make([]int, len(vstrs))
	for i, s := range vstrs {
		result[i], _ = strconv.Atoi(s)
	}
	return result
}

func winners(time int, distance int) int {
	count := 0
	for i := 0; i <= time; i++ {
		if i*(time-i) > distance {
			count++
		}
	}
	fmt.Println(time, distance, count)
	return count
}

func main() {

	// var total float64 = 0
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	// line := scanner.Text()
	times := parseLine(scanner.Text(), "Time")
	scanner.Scan()
	distances := parseLine(scanner.Text(), "Distance")
	fmt.Println(times, distances)

	// total := winners(71530, 940200) // sample
	total := winners(58819676, 434104122191218)
	fmt.Println(total)
}
