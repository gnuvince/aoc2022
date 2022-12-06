package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var crates [][]byte
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		var this_crate []byte
		for i := 1; i < len(line); i += 4 {
			this_crate = append(this_crate, line[i])
		}
		crates = append(crates, this_crate)
	}
	crates = crates[:len(crates)-1] // Remove the crate indexes

	var transposed [][]byte
	var transposed2 [][]byte
	for col := 0; col < len(crates[0]); col += 1 {
		var this_column []byte
		var this_column2 []byte
		for row := len(crates) - 1; row >= 0; row -= 1 {
			b := crates[row][col]
			if b != 32 {
				this_column = append(this_column, b)
				this_column2 = append(this_column2, b)
			}
		}
		transposed = append(transposed, this_column)
		transposed2 = append(transposed2, this_column2)
	}

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		qty, _ := strconv.Atoi(parts[1])
		from, _ := strconv.Atoi(parts[3])
		to, _ := strconv.Atoi(parts[5])

		// Part1
		for i := 0; i < qty; i += 1 {
			a := &transposed[from-1]
			b := &transposed[to-1]
			*b = append(*b, (*a)[len(*a)-1])
			*a = (*a)[:len(*a)-1]
		}

		// Part2
		{
			a := &transposed2[from-1]
			b := &transposed2[to-1]
			cutPoint := len(*a) - qty
			*b = append(*b, (*a)[cutPoint:]...)
			*a = (*a)[:cutPoint]
		}
	}

	var answer1 []byte
	for _, column := range transposed {
		answer1 = append(answer1, column[len(column)-1])
	}

	var answer2 []byte
	for _, column := range transposed2 {
		answer2 = append(answer2, column[len(column)-1])
	}
	fmt.Println(string(answer1), string(answer2))
}
