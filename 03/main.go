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
	var sum int
	items := map[rune]bool{}
	for _, line := range lines {
		if line == "" {
			continue
		}
		clear(items)

		for _, c := range line[:len(line)/2] {
			items[c] = true
		}
		for _, c := range line[len(line)/2:] {
			if !items[c] {
				continue
			}
			sum += priority(c)
			break
		}
	}
	fmt.Println("Part 1:", sum)
}

func priority(c rune) int {
	if 'a' <= c && c <= 'z' {
		return int(c - 'a' + 1)
	}
	if 'A' <= c && c <= 'Z' {
		return int(c - 'A' + 27)
	}
	panic("invalid character")
}

func part2(lines []string) {
	var sum int
	items1 := map[rune]bool{}
	items3 := map[rune]int{}
	for i := 0; i < len(lines); i += 3 {
		clear(items3)

		for j := 0; j < 3; j++ {
			clear(items1)
			for _, c := range lines[i+j] {
				items1[c] = true
			}
			for c := range items1 {
				items3[c]++
			}
		}
		for c, v := range items3 {
			if v == 3 {
				sum += priority(c)
				break
			}
		}
	}
	fmt.Println("Part 2:", sum)
}
