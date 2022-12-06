package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	total1 := 0
	for _, line := range lines {
		a := line[:len(line)/2]
		b := line[len(line)/2:]
		total1 += common1(a, b)
	}

	total2 := 0
	for i := 0; i < len(lines); i += 3 {
		total2 += common2(lines[i], lines[i+1], lines[i+2])
	}
	fmt.Println(total1, total2)
}

func common1(a, b string) int {
	for _, c := range a {
		for _, d := range b {
			if c == d {
				return value(c)
			}
		}
	}
	return 0
}

func common2(x, y, z string) int {
	for _, a := range x {
		for _, b := range y {
			if a == b {
				for _, c := range z {
					if c == a {
						return value(c)
					}
				}
			}
		}
	}
	return 0
}

func value(c rune) int {
	if c >= 'a' && c <= 'z' {
		return int(c - 'a' + 1)
	} else if c >= 'A' && c <= 'Z' {
		return int(c - 'A' + 27)
	} else {
		return 0
	}
}
