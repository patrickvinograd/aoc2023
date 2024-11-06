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

type Combo struct {
	Red   int
	Green int
	Blue  int
}

var Constraint Combo = Combo{12, 13, 14}

func checkCubes(line string) int {
	x := strings.Split(line, ":")
	var id int
	fmt.Sscanf(x[0], "Game %d", &id)
	pStrs := strings.Split(x[1], ";")
	pulls := make([]Combo, len(pStrs))
	for i, p := range pStrs {
		cStrs := strings.Split(p, ",")
		var pull Combo
		for _, c := range cStrs {
			var count int
			var color string
			fmt.Sscanf(c, " %d %s", &count, &color)
			if color == "red" {
				pull.Red = count
			} else if color == "green" {
				pull.Green = count
			} else if color == "blue" {
				pull.Blue = count
			}
		}
		pulls[i] = pull
	}

	possible := true
	for _, pull := range pulls {
		if pull.Red > Constraint.Red || pull.Green > Constraint.Green || pull.Blue > Constraint.Blue {
			possible = false
		}
	}
	//fmt.Println(id, pulls, possible)
	if possible {
		return id
	} else {
		return 0
	}
}

func main() {

	var total int = 0
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		total += checkCubes(line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}

	fmt.Println(total)
}
