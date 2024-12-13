/*
Напиши код на Go для решения следующей задачи.
Входные данные в файле указываемом аргументом в командной строке.
Выведи ответ и время решения после решения каждой части.
Входные данные такие, что перебор всех вариантов невозможен.

https://adventofcode.com/2022/day/15

--- Day 15: Beacon Exclusion Zone ---
You feel the ground rumble again as the distress signal leads you to a large network of subterranean tunnels. You don't have time to search them all, but you don't need to: your pack contains a set of deployable sensors that you imagine were originally built to locate lost Elves.

The sensors aren't very powerful, but that's okay; your handheld device indicates that you're close enough to the source of the distress signal to use them. You pull the emergency sensor system out of your pack, hit the big button on top, and the sensors zoom off down the tunnels.

Once a sensor finds a spot it thinks will give it a good reading, it attaches itself to a hard surface and begins monitoring for the nearest signal source beacon. Sensors and beacons always exist at integer coordinates. Each sensor knows its own position and can determine the position of a beacon precisely; however, sensors can only lock on to the one beacon closest to the sensor as measured by the Manhattan distance. (There is never a tie where two beacons are the same distance to a sensor.)

It doesn't take long for the sensors to report back their positions and closest beacons (your puzzle input). For example:

Sensor at x=2, y=18: closest beacon is at x=-2, y=15
Sensor at x=9, y=16: closest beacon is at x=10, y=16
Sensor at x=13, y=2: closest beacon is at x=15, y=3
Sensor at x=12, y=14: closest beacon is at x=10, y=16
Sensor at x=10, y=20: closest beacon is at x=10, y=16
Sensor at x=14, y=17: closest beacon is at x=10, y=16
Sensor at x=8, y=7: closest beacon is at x=2, y=10
Sensor at x=2, y=0: closest beacon is at x=2, y=10
Sensor at x=0, y=11: closest beacon is at x=2, y=10
Sensor at x=20, y=14: closest beacon is at x=25, y=17
Sensor at x=17, y=20: closest beacon is at x=21, y=22
Sensor at x=16, y=7: closest beacon is at x=15, y=3
Sensor at x=14, y=3: closest beacon is at x=15, y=3
Sensor at x=20, y=1: closest beacon is at x=15, y=3
So, consider the sensor at 2,18; the closest beacon to it is at -2,15. For the sensor at 9,16, the closest beacon to it is at 10,16.

Drawing sensors as S and beacons as B, the above arrangement of sensors and beacons looks like this:

               1    1    2    2
     0    5    0    5    0    5
 0 ....S.......................
 1 ......................S.....
 2 ...............S............
 3 ................SB..........
 4 ............................
 5 ............................
 6 ............................
 7 ..........S.......S.........
 8 ............................
 9 ............................
10 ....B.......................
11 ..S.........................
12 ............................
13 ............................
14 ..............S.......S.....
15 B...........................
16 ...........SB...............
17 ................S..........B
18 ....S.......................
19 ............................
20 ............S......S........
21 ............................
22 .......................B....
This isn't necessarily a comprehensive map of all beacons in the area, though. Because each sensor only identifies its closest beacon, if a sensor detects a beacon, you know there are no other beacons that close or closer to that sensor. There could still be beacons that just happen to not be the closest beacon to any sensor. Consider the sensor at 8,7:

               1    1    2    2
     0    5    0    5    0    5
-2 ..........#.................
-1 .........###................
 0 ....S...#####...............
 1 .......#######........S.....
 2 ......#########S............
 3 .....###########SB..........
 4 ....#############...........
 5 ...###############..........
 6 ..#################.........
 7 .#########S#######S#........
 8 ..#################.........
 9 ...###############..........
10 ....B############...........
11 ..S..###########............
12 ......#########.............
13 .......#######..............
14 ........#####.S.......S.....
15 B........###................
16 ..........#SB...............
17 ................S..........B
18 ....S.......................
19 ............................
20 ............S......S........
21 ............................
22 .......................B....
This sensor's closest beacon is at 2,10, and so you know there are no beacons that close or closer (in any positions marked #).

None of the detected beacons seem to be producing the distress signal, so you'll need to work out where the distress beacon is by working out where it isn't. For now, keep things simple by counting the positions where a beacon cannot possibly be along just a single row.

So, suppose you have an arrangement of beacons and sensors like in the example above and, just in the row where y=10, you'd like to count the number of positions a beacon cannot possibly exist. The coverage from all sensors near that row looks like this:

                 1    1    2    2
       0    5    0    5    0    5
 9 ...#########################...
10 ..####B######################..
11 .###S#############.###########.
In this example, in the row where y=10, there are 26 positions where a beacon cannot be present.

Consult the report from the sensors you just deployed. In the row where y=2000000, how many positions cannot contain a beacon?

--- Part Two ---
Your handheld device indicates that the distress signal is coming from a beacon nearby. The distress beacon is not detected by any sensor, but the distress beacon must have x and y coordinates each no lower than 0 and no larger than 4000000.

To isolate the distress beacon's signal, you need to determine its tuning frequency, which can be found by multiplying its x coordinate by 4000000 and then adding its y coordinate.

In the example above, the search space is smaller: instead, the x and y coordinates can each be at most 20. With this reduced search area, there is only a single position that could have a beacon: x=14, y=11. The tuning frequency for this distress beacon is 56000011.

Find the only possible position for the distress beacon. What is its tuning frequency?

*/

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
