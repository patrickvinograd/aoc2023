package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Coord struct {
	row int
	col int
}

var gears = make(map[Coord][]int)

func linkGear(row int, col int, val int) {
	var c = Coord{row, col}
	gears[c] = append(gears[c], val)
}

func checkHalo(lines []string, row int, col_start int, col_end int, val int) bool {
	for y := row - 1; y <= row+1; y++ {
		for x := col_start - 1; x <= col_end+1; x++ {
			if y >= 0 && y < len(lines) && x >= 0 && x < len(lines[y]) {
				c := lines[y][x]
				if !isDigit(c) && c != 46 {
					if c == 42 {
						linkGear(y, x, val)
					}
				}
			}
		}
	}
	return false
}

func isDigit(b byte) bool {
	return (b >= 48 && b <= 57)
}

func score(lines []string) int {
	fmt.Println(lines)
	for y, line := range lines {
		x := 0
		chomp := false
		chompstart := -1
		for x < len(line) {
			if chomp == false && isDigit(line[x]) {
				chomp = true
				chompstart = x
			} else if chomp == true && !isDigit(line[x]) {
				// fmt.Println("found token", line[chompstart:x])
				val, _ := strconv.Atoi(line[chompstart:x])
				// fmt.Println("score", val)
				checkHalo(lines, y, chompstart, x-1, val)
				// fmt.Println(touching)
				chomp = false
			}
			x++
		}
		if chomp == true {
			// fmt.Println("found token", line[chompstart:len(line)])
			val, _ := strconv.Atoi(line[chompstart:len(line)])
			// fmt.Println("score", val)
			checkHalo(lines, y, chompstart, len(line)-1, val)
			// fmt.Println(touching)
			// do search
		}
	}
	// fmt.Println(gears)
	var total int
	for _, value := range gears {
		if len(value) == 2 {
			total += (value[0] * value[1])
		}
	}
	return total
}

func main() {

	var total int = 0
	var lines []string
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	total = score(lines)

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}

	fmt.Println(total)
}
