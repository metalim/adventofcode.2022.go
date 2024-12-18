package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"time"
)

type Point struct {
	x, y int
}

type Sensor struct {
	pos      Point
	beacon   Point
	distance int
}

func manhattan(a, b Point) int {
	return abs(a.x-b.x) + abs(a.y-b.y)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func parseInput(filename string) ([]Sensor, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	sensors := []Sensor{}
	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile(`Sensor at x=([-]?\d+), y=([-]?\d+): closest beacon is at x=([-]?\d+), y=([-]?\d+)`)

	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindStringSubmatch(line)
		if matches == nil {
			continue
		}
		sx, _ := strconv.Atoi(matches[1])
		sy, _ := strconv.Atoi(matches[2])
		bx, _ := strconv.Atoi(matches[3])
		by, _ := strconv.Atoi(matches[4])
		sensor := Sensor{
			pos:    Point{x: sx, y: sy},
			beacon: Point{x: bx, y: by},
		}
		sensor.distance = manhattan(sensor.pos, sensor.beacon)
		sensors = append(sensors, sensor)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return sensors, nil
}

type Interval struct {
	start, end int
}

func mergeIntervals(intervals []Interval) []Interval {
	if len(intervals) == 0 {
		return intervals
	}
	// Sort intervals by start
	for i := 0; i < len(intervals)-1; i++ {
		for j := 0; j < len(intervals)-i-1; j++ {
			if intervals[j].start > intervals[j+1].start {
				intervals[j], intervals[j+1] = intervals[j+1], intervals[j]
			}
		}
	}

	merged := []Interval{intervals[0]}
	for _, current := range intervals[1:] {
		last := &merged[len(merged)-1]
		if current.start <= last.end+1 {
			if current.end > last.end {
				last.end = current.end
			}
		} else {
			merged = append(merged, current)
		}
	}
	return merged
}

func part1(sensors []Sensor, targetY int) int {
	intervals := []Interval{}

	for _, sensor := range sensors {
		dy := abs(sensor.pos.y - targetY)
		if dy > sensor.distance {
			continue
		}
		remaining := sensor.distance - dy
		interval := Interval{
			start: sensor.pos.x - remaining,
			end:   sensor.pos.x + remaining,
		}
		intervals = append(intervals, interval)
	}

	merged := mergeIntervals(intervals)

	// Count total covered positions
	total := 0
	for _, interval := range merged {
		total += interval.end - interval.start + 1
	}

	// Subtract positions where a beacon is present on targetY
	beaconsOnTarget := make(map[int]bool)
	for _, sensor := range sensors {
		if sensor.beacon.y == targetY {
			beaconsOnTarget[sensor.beacon.x] = true
		}
	}

	for x := range beaconsOnTarget {
		for _, interval := range merged {
			if x >= interval.start && x <= interval.end {
				total--
				break
			}
		}
	}

	return total
}

func part2(sensors []Sensor, maxCoord int) (int, int) {
	// For each sensor, consider the perimeter at distance+1
	for _, sensor := range sensors {
		dist := sensor.distance + 1
		for dx := 0; dx <= dist; dx++ {
			dy := dist - dx
			candidates := []Point{
				{sensor.pos.x + dx, sensor.pos.y + dy},
				{sensor.pos.x + dx, sensor.pos.y - dy},
				{sensor.pos.x - dx, sensor.pos.y + dy},
				{sensor.pos.x - dx, sensor.pos.y - dy},
			}
			for _, p := range candidates {
				if p.x < 0 || p.x > maxCoord || p.y < 0 || p.y > maxCoord {
					continue
				}
				covered := false
				for _, other := range sensors {
					if manhattan(p, other.pos) <= other.distance {
						covered = true
						break
					}
				}
				if !covered {
					return p.x, p.y
				}
			}
		}
	}
	return -1, -1
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <input_file>", os.Args[0])
	}
	filename := os.Args[1]

	sensors, err := parseInput(filename)
	if err != nil {
		log.Fatalf("Error parsing input: %v", err)
	}

	// Часть 1
	start1 := time.Now()
	result1 := part1(sensors, 2000000)
	elapsed1 := time.Since(start1)
	fmt.Printf("Часть 1: %d (Время: %s)\n", result1, elapsed1)

	// Часть 2
	start2 := time.Now()
	x, y := part2(sensors, 4000000)
	if x != -1 && y != -1 {
		tuningFrequency := x*4000000 + y
		elapsed2 := time.Since(start2)
		fmt.Printf("Часть 2: x=%d, y=%d, частота настройки=%d (Время: %s)\n", x, y, tuningFrequency, elapsed2)
	} else {
		fmt.Println("Часть 2: Решение не найдено")
	}
}
