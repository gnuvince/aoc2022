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
	head := point{0, 0}
	tail := point{0, 0}
	tail2 := point{0, 0}
	tail3 := point{0, 0}
	tail4 := point{0, 0}
	tail5 := point{0, 0}
	tail6 := point{0, 0}
	tail7 := point{0, 0}
	tail8 := point{0, 0}
	tail9 := point{0, 0}
	visited := make(map[point]int)
	visited9 := make(map[point]int)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		instruction := strings.Split(scanner.Text(), " ")
		dir := instruction[0]
		steps, _ := strconv.Atoi(instruction[1])
		for i := 0; i < steps; i += 1 {
			switch dir {
			case "U":
				head.y += 1
			case "D":
				head.y -= 1
			case "R":
				head.x += 1
			case "L":
				head.x -= 1
			}
			tail = updateTail(head, tail)
			tail2 = updateTail(tail, tail2)
			tail3 = updateTail(tail2, tail3)
			tail4 = updateTail(tail3, tail4)
			tail5 = updateTail(tail4, tail5)
			tail6 = updateTail(tail5, tail6)
			tail7 = updateTail(tail6, tail7)
			tail8 = updateTail(tail7, tail8)
			tail9 = updateTail(tail8, tail9)

			visited[tail] += 1
			visited9[tail9] += 1
		}
	}
	fmt.Println(len(visited), len(visited9))
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
