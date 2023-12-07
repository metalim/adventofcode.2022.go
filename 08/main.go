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

type Pos struct {
	x, y int
}

func part1(lines []string) {
	visible := map[Pos]bool{}
	for y, line := range lines {
		// LR
		tallest := '0' - 1
		for x, c := range line {
			if c > tallest {
				tallest = c
				visible[Pos{x, y}] = true
			}
		}

		// RL
		tallest = '0' - 1
		for x := len(line) - 1; x >= 0; x-- {
			c := rune(line[x])
			if c > tallest {
				tallest = c
				visible[Pos{x, y}] = true
			}
		}
	}

	// UD
	for x := 0; x < len(lines[0]); x++ {
		tallest := '0' - 1
		for y := 0; y < len(lines); y++ {
			c := rune(lines[y][x])
			if c > tallest {
				tallest = c
				visible[Pos{x, y}] = true
			}
		}

		// DU
		tallest = '0' - 1
		for y := len(lines) - 1; y >= 0; y-- {
			c := rune(lines[y][x])
			if c > tallest {
				tallest = c
				visible[Pos{x, y}] = true
			}
		}
	}

	fmt.Println("Part 1:", len(visible))
}

func part2(lines []string) {
}
