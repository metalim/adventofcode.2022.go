package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

var Start = Point{500, 0}
var print bool

func catch(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	flag.BoolVar(&print, "print", false, "print the grid")
	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Println("Usage: go run main.go input.txt")
		os.Exit(1)
	}

	bs, err := os.ReadFile(flag.Arg(0))
	catch(err)

	input := parseInput(string(bs))
	part1(input)
	input = parseInput(string(bs))
	part2(input)
}

type Point [2]int

func (p Point) Add(q Point) Point {
	return Point{p[0] + q[0], p[1] + q[1]}
}

func (p Point) Sub(q Point) Point {
	return Point{p[0] - q[0], p[1] - q[1]}
}

func (p Point) Sign() Point {
	return Point{sign(p[0]), sign(p[1])}
}

func sign(x int) int {
	if x < 0 {
		return -1
	}
	if x > 0 {
		return 1
	}
	return 0
}

func (p Point) Min(q Point) Point {
	return Point{min(p[0], q[0]), min(p[1], q[1])}
}

func (p Point) Max(q Point) Point {
	return Point{max(p[0], q[0]), max(p[1], q[1])}
}

type Grid struct {
	Data   map[Point]rune
	TL, BR Point
}

func (g *Grid) IsFree(p Point) bool {
	_, ok := g.Data[p]
	return !ok
}

func (g *Grid) Draw(p Point, v rune) {
	g.Data[p] = v
	g.TL = g.TL.Min(p)
	g.BR = g.BR.Max(p)
}

func (g *Grid) Print() {
	if !print {
		return
	}
	for y := g.TL[1]; y <= g.BR[1]; y++ {
		for x := g.TL[0]; x <= g.BR[0]; x++ {
			if v, ok := g.Data[Point{x, y}]; ok {
				fmt.Printf("%c", v)
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	catch(err)
	return i
}

func toPoint(s string) Point {
	parts := strings.Split(s, ",")
	return Point{toInt(parts[0]), toInt(parts[1])}
}

func parseInput(input string) Grid {
	lines := strings.Split(input, "\n")
	if len(lines[len(lines)-1]) == 0 {
		lines = lines[:len(lines)-1]
	}

	grid := Grid{
		Data: map[Point]rune{},
		TL:   Point{math.MaxInt, math.MaxInt},
		BR:   Point{math.MinInt, math.MinInt},
	}

	for _, line := range lines {
		points := strings.Split(line, " -> ")

		p0 := toPoint(points[0])
		grid.Draw(p0, '#')
		for _, p := range points[1:] {
			p1 := toPoint(p)
			dir := p1.Sub(p0).Sign()
			for p0 != p1 {
				grid.Draw(p0, '#')
				p0 = p0.Add(dir)
			}
			grid.Draw(p1, '#')
		}
	}
	grid.Draw(Start, '+')
	return grid
}

var dir = []Point{{0, 1}, {-1, 1}, {1, 1}}

func part1(grid Grid) {
	timeStart := time.Now()

	bottom := grid.BR[1]
	count := 0
	for {
		p := Start
	FALL:
		for p[1] <= bottom {
			for _, d := range dir {
				np := p.Add(d)
				if grid.IsFree(np) {
					p = np
					continue FALL
				}
			}
			break FALL
		}
		if p[1] > grid.BR[1] {
			break
		}
		grid.Draw(p, 'o')
		count++
	}
	grid.Print()

	fmt.Printf("Part 1: %d\t\tin %v\n", count, time.Since(timeStart))
}

func part2(grid Grid) {
	timeStart := time.Now()

	bottom := grid.BR[1]
	count := 0
	for {
		p := Start
	FALL:
		for p[1] <= bottom {
			for _, d := range dir {
				np := p.Add(d)
				if grid.IsFree(np) {
					p = np
					continue FALL
				}
			}
			break FALL
		}
		grid.Draw(p, 'o')
		count++
		if p == Start {
			break
		}
	}
	grid.Print()

	fmt.Printf("Part 2: %d\t\tin %v\n", count, time.Since(timeStart))
}
