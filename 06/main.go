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

func findUnique(line []rune, n int) int {
	seq := map[rune]int{}
	for i, c := range line {
		seq[c]++
		if i >= n {
			if seq[line[i-n]] > 1 {
				seq[line[i-n]]--
			} else {
				delete(seq, line[i-n])
			}
		}
		if len(seq) == n {
			return i + 1
		}
	}
	panic("not found")
}

func part1(lines []string) {
	line := []rune(lines[0])
	i := findUnique(line, 4)
	fmt.Println("Part 1:", i, string(line[i-4:i]))
}

func part2(lines []string) {
	line := []rune(lines[0])
	i := findUnique(line, 14)
	fmt.Println("Part 1:", i, string(line[i-14:i]))
}
