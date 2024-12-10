package main

import (
	"flag"
	"fmt"
	"maps"
	"os"
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

type Row []rune
type Grid []Row
type Input struct {
	grid       Grid
	start, end Point
}

func parseInput(input string) Input {
	lines := strings.Split(input, "\n")
	if len(lines[len(lines)-1]) == 0 {
		lines = lines[:len(lines)-1]
	}
	grid := []Row{}
	var start, end Point
	for y, line := range lines {
		grid = append(grid, Row(line))
		for x, c := range line {
			if c == 'S' {
				start = Point{y, x}
			} else if c == 'E' {
				end = Point{y, x}
			}
		}
	}

	grid[start[0]][start[1]] = 'a'
	grid[end[0]][end[1]] = 'z'

	return Input{grid, start, end}
}

type Point [2]int

func (p Point) Add(q Point) Point {
	return Point{p[0] + q[0], p[1] + q[1]}
}

var dirs = []Point{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

func bfs(grid Grid, next map[Point]struct{}, end Point) (int, bool) {
	visited := maps.Clone(next)
	steps := 0
	cur := map[Point]struct{}{}
	H, W := len(grid), len(grid[0])
	for ; len(next) > 0; steps++ {
		cur, next = next, cur
		clear(next)
		for p := range cur {
			if p == end {
				return steps, true
			}
			for _, dir := range dirs {
				np := p.Add(dir)
				if np[0] < 0 || np[0] >= H || np[1] < 0 || np[1] >= W {
					continue
				}
				if _, ok := visited[np]; ok {
					continue
				}
				if grid[np[0]][np[1]] > grid[p[0]][p[1]]+1 {
					continue
				}
				visited[np] = struct{}{}
				next[np] = struct{}{}
			}
		}
	}
	return steps, false
}

func part1(input Input) {
	timeStart := time.Now()

	steps, found := bfs(input.grid, map[Point]struct{}{input.start: {}}, input.end)
	if !found {
		panic("path not found")
	}

	fmt.Printf("Part 1: %d\t\tin %v\n", steps, time.Since(timeStart))
}

func part2(input Input) {
	timeStart := time.Now()
	start := map[Point]struct{}{}
	for y, row := range input.grid {
		for x, c := range row {
			if c == 'a' {
				start[Point{y, x}] = struct{}{}
			}
		}
	}
	steps, found := bfs(input.grid, start, input.end)
	if !found {
		panic("path not found")
	}

	fmt.Printf("Part 2: %d\t\tin %v\n", steps, time.Since(timeStart))
}
