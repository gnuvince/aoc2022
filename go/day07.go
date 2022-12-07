package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type filesize struct {
	path []string
	name string
	size int
}

func main() {
	var cwd []string
	var filesizes []filesize
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		if parts[0] == "$" {
			if line == "$ cd /" {
				// ignore
			} else if line == "$ cd .." {
				cwd = cwd[:len(cwd)-1]
			} else if parts[1] == "cd" {
				cwd = append(cwd, parts[2])
			}
		} else if parts[0] == "dir" {
			// ignore
		} else {
			size, _ := strconv.Atoi(parts[0])
			thisCwd := make([]string, len(cwd))
			copy(thisCwd, cwd)
			filesizes = append(filesizes, filesize{thisCwd, parts[1], size})
		}
	}
	dirSizes := make(map[string]int)
	for _, f := range filesizes {
		for i := 0; i <= len(f.path); i += 1 {
			dir := strings.Join(f.path[:i], "/")
			dirSizes[dir] += f.size
		}
	}
	fmt.Println(part1(dirSizes))
	fmt.Println(part2(dirSizes))
}

func part1(dirsizes map[string]int) int {
	total := 0
	for _, size := range dirsizes {
		if size <= 100000 {
			total += size
		}
	}
	return total
}

func part2(dirsizes map[string]int) int {
	free_space := 70000000 - dirsizes[""]
	needed_space := 30000000 - free_space
	best := 2 << 32
	for _, size := range dirsizes {
		if size >= needed_space && size < best {
			best = size
		}
	}
	return best
}
