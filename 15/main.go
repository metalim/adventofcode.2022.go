package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

const MULTIPLIER = 4e6

var Y int = 2e6
var MAX int = MULTIPLIER

func catch(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	flag.IntVar(&Y, "y", 2e6, "y coordinate")
	flag.IntVar(&MAX, "max", MULTIPLIER, "max coordinate")
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

type Point [2]int
type Input [][2]Point

func (p Point) Min(other Point) Point {
	return Point{min(p[0], other[0]), min(p[1], other[1])}
}

func (p Point) Max(other Point) Point {
	return Point{max(p[0], other[0]), max(p[1], other[1])}
}

func (p Point) Add(other Point) Point {
	return Point{p[0] + other[0], p[1] + other[1]}
}

func parseInput(input string) Input {
	lines := strings.Split(input, "\n")
	if len(lines[len(lines)-1]) == 0 {
		lines = lines[:len(lines)-1]
	}
	var result Input
	for _, line := range lines {
		var s, b Point
		_, err := fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &s[0], &s[1], &b[0], &b[1])
		catch(err)
		result = append(result, [2]Point{s, b})
	}
	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func manhattanDistance(a, b Point) int {
	return abs(a[0]-b[0]) + abs(a[1]-b[1])
}

func part1(pairs Input) {
	timeStart := time.Now()
	line := map[int]struct{}{}
	for _, pair := range pairs {
		s, b := pair[0], pair[1]
		sensorRange := manhattanDistance(s, b)
		distToLine := abs(s[1] - Y)
		if distToLine <= sensorRange {
			for x := s[0] - sensorRange + distToLine; x <= s[0]+sensorRange-distToLine; x++ {
				line[x] = struct{}{}
			}
		}
	}

	for _, pair := range pairs {
		b := pair[1]
		if b[1] == Y {
			delete(line, b[0])
		}
	}

	fmt.Printf("Part 1: %d\t\tin %v\n", len(line), time.Since(timeStart))
}

type Sensor struct {
	Point
	Range int
}

func part2(pairs Input) {
	timeStart := time.Now()

	sensors := []Sensor{}
	for _, pair := range pairs {
		s, b := pair[0], pair[1]
		sensors = append(sensors, Sensor{s, manhattanDistance(s, b)})
	}

	var found bool
	var beacon Point
	ranges := make([][2]int, len(sensors))
SCAN:
	for y := 0; y <= MAX; y++ {
		for i, sensor := range sensors {
			distToLine := abs(sensor.Point[1] - y)
			if distToLine > sensor.Range {
				continue
			}
			minX := sensor.Point[0] - sensor.Range + distToLine
			maxX := sensor.Point[0] + sensor.Range - distToLine
			ranges[i] = [2]int{minX, maxX}
		}
		for _, r := range ranges {
			x0, x1 := r[0]-1, r[1]+1
			if x0 >= 0 && x0 <= MAX && !isCovered(x0, ranges) {
				found = true
				beacon = Point{x0, y}
				break SCAN
			}
			if x1 >= 0 && x1 <= MAX && !isCovered(x1, ranges) {
				found = true
				beacon = Point{x1, y}
				break SCAN
			}
		}
	}

	if found {
		fmt.Printf("Found beacon: %d,%d\n", beacon[0], beacon[1])
		freq := beacon[0]*MULTIPLIER + beacon[1]
		fmt.Printf("Part 2: %d\t\tin %v\n", freq, time.Since(timeStart))
	} else {
		fmt.Println("Not found")
	}
}

func isCovered(x int, ranges [][2]int) bool {
	for _, r := range ranges {
		if x >= r[0] && x <= r[1] {
			return true
		}
	}
	return false
}
