package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func catch(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Println("Usage: go run main.go input.txt")
		os.Exit(1)
	}

	bs, err := os.ReadFile(flag.Arg(0))
	catch(err)

	input := parseInput(string(bs))
	part1(input)
	part2(input)
}

type Input []string

func parseInput(input string) Input {
	lines := strings.Split(input, "\n")
	if len(lines[len(lines)-1]) == 0 {
		lines = lines[:len(lines)-1]
	}
	return Input(lines)
}

var reOp = regexp.MustCompile(`^(noop)|(addx)\s+(-?\d+)$`)

func process(input Input, tick func(), addx func(n int)) {
	for _, line := range input {
		m := reOp.FindStringSubmatch(line)
		if m == nil {
			panic(fmt.Sprintf("invalid line: %s", line))
		}
		switch {
		case m[1] == "noop":
			tick()

		case m[2] == "addx":
			tick()
			tick()
			v, err := strconv.Atoi(m[3])
			catch(err)
			addx(v)
		}
	}
}

func part1(input Input) {
	timeStart := time.Now()

	var sum int
	X := 1
	cycle := 1
	tick := func() {
		if (cycle-20)%40 == 0 {
			sum += X * cycle
		}
		cycle++
	}
	addx := func(n int) {
		X += n
	}
	process(input, tick, addx)

	fmt.Printf("Part 1: %d\t\tin %v\n", sum, time.Since(timeStart))
}

func part2(input Input) {
	timeStart := time.Now()

	screen := []rune{}
	X := 1
	cycle := 1
	tick := func() {
		if X <= cycle%40 && cycle%40 <= X+2 {
			screen = append(screen, '#')
		} else {
			screen = append(screen, '.')
		}
		cycle++
	}
	addx := func(n int) {
		X += n
	}
	process(input, tick, addx)

	fmt.Printf("Part 2: \t\tin %v\n", time.Since(timeStart))
	for i := 0; i < len(screen); i += 40 {
		fmt.Println(string(screen[i : i+40]))
	}
}
