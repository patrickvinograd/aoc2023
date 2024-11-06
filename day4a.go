package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

// type Card struct {
// 	Red   int
// 	Green int
// 	Blue  int
// }

var cardmap = make(map[int]int)

func checkTicket(line string) {
	x := strings.Split(line, ":")
	var id int
	fmt.Sscanf(x[0], "Card %d", &id)
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

	score := 0
	for _, x := range havers {
		if slices.Contains(winners, x) {
			score++
		}
	}
	// fmt.Println(id, score)
	cardmap[id] = score
}

func recurseTicket(i int) int {
	total := 0
	total += cardmap[i]
	for x := 1; x <= cardmap[i]; x++ {
		total += recurseTicket(i + x)
	}
	// fmt.Println(i, total)
	return total
}

func countTickets() int {
	total := 0
	for key, _ := range cardmap {
		total += recurseTicket(key)
		total += 1 // for the original card
	}
	return total
}

func main() {

	// var total float64 = 0
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		checkTicket(line)
	}
	fmt.Println(cardmap)
	total := countTickets()

	fmt.Println(total)
}
