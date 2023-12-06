package main

import (
	"fmt"
	"os"
	"strings"
)

func catch(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go input.txt")
		os.Exit(1)
	}

	bs, err := os.ReadFile(os.Args[1])
	catch(err)
	lines := strings.Split(string(bs), "\n")

	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	var score int
	for _, line := range lines {
		switch line {
		case "B X":
			score += 0 + 1
		case "C Y":
			score += 0 + 2
		case "A Z":
			score += 0 + 3
		case "A X":
			score += 3 + 1
		case "B Y":
			score += 3 + 2
		case "C Z":
			score += 3 + 3
		case "C X":
			score += 6 + 1
		case "A Y":
			score += 6 + 2
		case "B Z":
			score += 6 + 3
		}
	}
	fmt.Println("Part 1:", score)
}

func part2(lines []string) {
	var score int
	for _, line := range lines {
		switch line {
		case "B X":
			score += 0 + 1
		case "C X":
			score += 0 + 2
		case "A X":
			score += 0 + 3
		case "A Y":
			score += 3 + 1
		case "B Y":
			score += 3 + 2
		case "C Y":
			score += 3 + 3
		case "C Z":
			score += 6 + 1
		case "A Z":
			score += 6 + 2
		case "B Z":
			score += 6 + 3
		}
	}
	fmt.Println("Part 2:", score)
}
