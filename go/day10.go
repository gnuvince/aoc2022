package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	cycle := 0
	x := 1
	p1 := 0
	var crt [6][40]byte
	crt_y, crt_x := 0, 0
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		instr_cycles := 0
		dx := 0
		switch parts[0] {
		case "noop":
			instr_cycles = 1
		case "addx":
			dx, _ = strconv.Atoi(parts[1])
			instr_cycles = 2
		}
		sprite_idx := x - 1
		for i := 0; i < instr_cycles; i += 1 {
			if crt_x >= sprite_idx && crt_x < sprite_idx+3 {
				crt[crt_y][crt_x] = '#'
			} else {
				crt[crt_y][crt_x] = ' '
			}
			crt_x += 1
			cycle += 1
			if cycle%40 == 0 {
				crt_y += 1
				crt_x = 0
			}
			if cycle%40 == 20 {
				p1 += cycle * x
			}
		}
		x += dx
	}
	fmt.Println(p1)
	for _, row := range crt {
		fmt.Println(string(row[:]))
	}
}
