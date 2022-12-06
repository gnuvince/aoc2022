package main

import (
	"fmt"
	"os"
)

func main() {
	part1 := 0
	part2 := 0
	for {
		var a, b, c, d int
		n, err := fmt.Fscanf(os.Stdin, "%d-%d,%d-%d\n", &a, &b, &c, &d)
		if n != 4 || err != nil {
			break
		}
		if (a <= c && b >= d) || (c <= a && d >= b) {
			part1 += 1
		}

		if a <= c && b >= c || a <= d && b >= d || c <= a && d >= a || c <= b && d >= b {
			part2 += 1
		}
	}
	fmt.Println(part1, part2)
}
