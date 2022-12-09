package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x, y int
}

func main() {
	var knots [10]point
	visited_tail := make(map[point]int)
	visited_tail9 := make(map[point]int)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		instruction := strings.Split(scanner.Text(), " ")
		dir := instruction[0]
		steps, _ := strconv.Atoi(instruction[1])
		for i := 0; i < steps; i += 1 {
			switch dir {
			case "U":
				knots[0].y += 1
			case "D":
				knots[0].y -= 1
			case "R":
				knots[0].x += 1
			case "L":
				knots[0].x -= 1
			}
			for i := 1; i < 10; i += 1 {
				knots[i] = updateTail(knots[i-1], knots[i])
			}
			visited_tail[knots[1]] += 1
			visited_tail9[knots[9]] += 1
		}
	}
	fmt.Println(len(visited_tail), len(visited_tail9))
}

func updateTail(head, tail point) point {
	if abs(head.x-tail.x) <= 1 && abs(head.y-tail.y) <= 1 {
		return tail
	}

	var dx, dy int
	if abs(head.x-tail.x) == 2 && abs(head.y-tail.y) == 2 {
		dx = (head.x - tail.x) / 2
		dy = (head.y - tail.y) / 2
	} else if abs(head.x-tail.x) == 2 {
		dx = (head.x - tail.x) / 2
		dy = head.y - tail.y
	} else if abs(head.y-tail.y) == 2 {
		dx = head.x - tail.x
		dy = (head.y - tail.y) / 2
	}
	return point{tail.x + dx, tail.y + dy}
}

func abs(a int) int {
	if a < 0 {
		return -a
	} else {
		return a
	}
}
