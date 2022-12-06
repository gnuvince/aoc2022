package main

import (
	"bufio"
	"fmt"
	"os"
)

type scores struct {
	part1 int
	part2 int
}

func main() {
	total1 := 0
	total2 := 0
	results := map[string]scores{
		"A X": {4, 3},
		"A Y": {8, 4},
		"A Z": {3, 8},
		"B X": {1, 1},
		"B Y": {5, 5},
		"B Z": {9, 9},
		"C X": {7, 2},
		"C Y": {2, 6},
		"C Z": {6, 7},
	}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		score := results[scanner.Text()]
		total1 += score.part1
		total2 += score.part2
	}
	fmt.Println(total1, total2)
}
