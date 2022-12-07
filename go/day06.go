package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Bytes()
		fmt.Println(FindStartOfPacket(line, 4), FindStartOfPacket(line, 14))
	}
}

func FindStartOfPacket(line []byte, numConsecutive int) int {
	for i := 0; i < len(line)-numConsecutive; i += 1 {
		if IsStartOfPacket(line[i : i+numConsecutive]) {
			return i + numConsecutive
		}
	}
	return -1
}

func IsStartOfPacket(packet []byte) bool {
	for i := 0; i < len(packet); i += 1 {
		for j := i + 1; j < len(packet); j += 1 {
			if packet[i] == packet[j] {
				return false
			}
		}
	}
	return true
}
