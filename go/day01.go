package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var sums []int
	curr := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			sums = append(sums, curr)
			curr = 0
		} else {
			weight, _ := strconv.Atoi(line)
			curr += weight
		}
	}
	sort.Ints(sums)
	fmt.Println(sums[len(sums)-1],
		sums[len(sums)-3]+sums[len(sums)-2]+sums[len(sums)-1])
}
