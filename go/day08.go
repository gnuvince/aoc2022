package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var grid []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}
	fmt.Println(part1(grid), part2(grid))
}

func part1(grid []string) int {
	total_visible := 0
	for y := 0; y < len(grid); y += 1 {
		for x := 0; x < len(grid[y]); x += 1 {
			if left(grid, y, x) || right(grid, y, x) || down(grid, y, x) || up(grid, y, x) {
				total_visible += 1
			}
		}
	}
	return total_visible
}

func left(grid []string, y, x int) bool {
	for i := 0; i < x; i += 1 {
		if grid[y][i] >= grid[y][x] {
			return false
		}
	}
	return true
}

func right(grid []string, y, x int) bool {
	for i := x + 1; i < len(grid[y]); i += 1 {
		if grid[y][i] >= grid[y][x] {
			return false
		}
	}
	return true
}

func up(grid []string, y, x int) bool {
	for i := 0; i < y; i += 1 {
		if grid[i][x] >= grid[y][x] {
			return false
		}
	}
	return true
}

func down(grid []string, y, x int) bool {
	for i := y + 1; i < len(grid); i += 1 {
		if grid[i][x] >= grid[y][x] {
			return false
		}
	}
	return true
}

func part2(grid []string) int {
	best_scenic := 0
	for y := 0; y < len(grid); y += 1 {
		for x := 0; x < len(grid[y]); x += 1 {
			scenic := scenic_left(grid, y, x) * scenic_right(grid, y, x) * scenic_up(grid, y, x) * scenic_down(grid, y, x)
			if scenic > best_scenic {
				best_scenic = scenic
			}
		}
	}
	return best_scenic
}

func scenic_left(grid []string, y, x int) int {
	score := 0
	subject := grid[y][x]
	for i := x - 1; i >= 0; i -= 1 {
		score += 1
		if subject <= grid[y][i] {
			break
		}
	}
	return score
}

func scenic_right(grid []string, y, x int) int {
	score := 0
	subject := grid[y][x]
	for i := x + 1; i < len(grid); i += 1 {
		score += 1
		if subject <= grid[y][i] {
			break
		}
	}
	return score
}

func scenic_up(grid []string, y, x int) int {
	score := 0
	subject := grid[y][x]
	for i := y - 1; i >= 0; i -= 1 {
		score += 1
		if subject <= grid[i][x] {
			break
		}
	}
	return score
}

func scenic_down(grid []string, y, x int) int {
	score := 0
	subject := grid[y][x]
	for i := y + 1; i < len(grid); i += 1 {
		score += 1
		if subject <= grid[i][x] {
			break
		}
	}
	return score
}
