/*
Напиши код на Go для решения задачи.
Входные данные в файле указываемом аргументом в командной строке.
Выведи ответ и время решения после решения каждой части.
Каждая часть должна решаться за несколько секунд максимум. Вторая часть задачи МОЖЕТ требовать особого подхода и не решаться перебором вариантов.
Если программа не сработает, я вставлю вывод и возможные комментарии. В ответ просто выдай исправленную версию.

https://adventofcode.com/2022/day/14

--- Day 14: Regolith Reservoir ---
The distress signal leads you to a giant waterfall! Actually, hang on - the signal seems like it's coming from the waterfall itself, and that doesn't make any sense. However, you do notice a little path that leads behind the waterfall.

Correction: the distress signal leads you behind a giant waterfall! There seems to be a large cave system here, and the signal definitely leads further inside.

As you begin to make your way deeper underground, you feel the ground rumble for a moment. Sand begins pouring into the cave! If you don't quickly figure out where the sand is going, you could quickly become trapped!

Fortunately, your familiarity with analyzing the path of falling material will come in handy here. You scan a two-dimensional vertical slice of the cave above you (your puzzle input) and discover that it is mostly air with structures made of rock.

Your scan traces the path of each solid rock structure and reports the x,y coordinates that form the shape of the path, where x represents distance to the right and y represents distance down. Each path appears as a single line of text in your scan. After the first point of each path, each point indicates the end of a straight horizontal or vertical line to be drawn from the previous point. For example:

498,4 -> 498,6 -> 496,6
503,4 -> 502,4 -> 502,9 -> 494,9
This scan means that there are two paths of rock; the first path consists of two straight lines, and the second path consists of three straight lines. (Specifically, the first path consists of a line of rock from 498,4 through 498,6 and another line of rock from 498,6 through 496,6.)

The sand is pouring into the cave from point 500,0.

Drawing rock as #, air as ., and the source of the sand as +, this becomes:


  4     5  5
  9     0  0
  4     0  3
0 ......+...
1 ..........
2 ..........
3 ..........
4 ....#...##
5 ....#...#.
6 ..###...#.
7 ........#.
8 ........#.
9 #########.
Sand is produced one unit at a time, and the next unit of sand is not produced until the previous unit of sand comes to rest. A unit of sand is large enough to fill one tile of air in your scan.

A unit of sand always falls down one step if possible. If the tile immediately below is blocked (by rock or sand), the unit of sand attempts to instead move diagonally one step down and to the left. If that tile is blocked, the unit of sand attempts to instead move diagonally one step down and to the right. Sand keeps moving as long as it is able to do so, at each step trying to move down, then down-left, then down-right. If all three possible destinations are blocked, the unit of sand comes to rest and no longer moves, at which point the next unit of sand is created back at the source.

So, drawing sand that has come to rest as o, the first unit of sand simply falls straight down and then stops:

......+...
..........
..........
..........
....#...##
....#...#.
..###...#.
........#.
......o.#.
#########.
The second unit of sand then falls straight down, lands on the first one, and then comes to rest to its left:

......+...
..........
..........
..........
....#...##
....#...#.
..###...#.
........#.
.....oo.#.
#########.
After a total of five units of sand have come to rest, they form this pattern:

......+...
..........
..........
..........
....#...##
....#...#.
..###...#.
......o.#.
....oooo#.
#########.
After a total of 22 units of sand:

......+...
..........
......o...
.....ooo..
....#ooo##
....#ooo#.
..###ooo#.
....oooo#.
...ooooo#.
#########.
Finally, only two more units of sand can possibly come to rest:

......+...
..........
......o...
.....ooo..
....#ooo##
...o#ooo#.
..###ooo#.
....oooo#.
.o.ooooo#.
#########.
Once all 24 units of sand shown above have come to rest, all further sand flows out the bottom, falling into the endless void. Just for fun, the path any new sand takes before falling forever is shown here with ~:

.......+...
.......~...
......~o...
.....~ooo..
....~#ooo##
...~o#ooo#.
..~###ooo#.
..~..oooo#.
.~o.ooooo#.
~#########.
~..........
~..........
~..........
Using your scan, simulate the falling sand. How many units of sand come to rest before sand starts flowing into the abyss below?

--- Part Two ---
You realize you misread the scan. There isn't an endless void at the bottom of the scan - there's floor, and you're standing on it!

You don't have time to scan the floor, so assume the floor is an infinite horizontal line with a y coordinate equal to two plus the highest y coordinate of any point in your scan.

In the example above, the highest y coordinate of any point is 9, and so the floor is at y=11. (This is as if your scan contained one extra rock path like -infinity,11 -> infinity,11.) With the added floor, the example above now looks like this:

        ...........+........
        ....................
        ....................
        ....................
        .........#...##.....
        .........#...#......
        .......###...#......
        .............#......
        .............#......
        .....#########......
        ....................
<-- etc #################### etc -->
To find somewhere safe to stand, you'll need to simulate falling sand until a unit of sand comes to rest at 500,0, blocking the source entirely and stopping the flow of sand into the cave. In the example above, the situation finally looks like this after 93 units of sand come to rest:

............o............
...........ooo...........
..........ooooo..........
.........ooooooo.........
........oo#ooo##o........
.......ooo#ooo#ooo.......
......oo###ooo#oooo......
.....oooo.oooo#ooooo.....
....oooooooooo#oooooo....
...ooo#########ooooooo...
..ooooo.......ooooooooo..
#########################
Using your scan, simulate the falling sand until the source of the sand becomes blocked. How many units of sand come to rest?


*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// Position represents a coordinate in the cave
type Position struct {
	x int
	y int
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <input_file>")
		return
	}

	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	// Parse input and get rock positions
	rockPaths, minX, maxX, maxY, err := parseInput(file)
	if err != nil {
		fmt.Printf("Error parsing input: %v\n", err)
		return
	}

	// For part 1
	startPart1 := time.Now()
	part1Answer := simulateSand(rockPaths, minX, maxX, maxY, false)
	durationPart1 := time.Since(startPart1)
	fmt.Printf("Part 1 Answer: %d\n", part1Answer)
	fmt.Printf("Part 1 Time: %v\n", durationPart1)

	// For part 2
	startPart2 := time.Now()
	part2Answer := simulateSand(rockPaths, minX, maxX, maxY, true)
	durationPart2 := time.Since(startPart2)
	fmt.Printf("Part 2 Answer: %d\n", part2Answer)
	fmt.Printf("Part 2 Time: %v\n", durationPart2)
}

// parseInput reads the input and returns rock paths, minX, maxX, maxY
func parseInput(file *os.File) ([]Position, int, int, int, error) {
	scanner := bufio.NewScanner(file)
	rockSet := make(map[Position]struct{})
	minX := 500
	maxX := 500
	maxY := 0

	for scanner.Scan() {
		line := scanner.Text()
		points := strings.Split(line, " -> ")
		var path []Position
		for _, p := range points {
			coords := strings.Split(p, ",")
			if len(coords) != 2 {
				return nil, 0, 0, 0, fmt.Errorf("invalid point: %s", p)
			}
			x, err := strconv.Atoi(coords[0])
			if err != nil {
				return nil, 0, 0, 0, fmt.Errorf("invalid x coordinate: %s", coords[0])
			}
			y, err := strconv.Atoi(coords[1])
			if err != nil {
				return nil, 0, 0, 0, fmt.Errorf("invalid y coordinate: %s", coords[1])
			}
			path = append(path, Position{x, y})
			if x < minX {
				minX = x
			}
			if x > maxX {
				maxX = x
			}
			if y > maxY {
				maxY = y
			}
		}
		// Fill in the rocks between consecutive points
		for i := 0; i < len(path)-1; i++ {
			from := path[i]
			to := path[i+1]
			if from.x == to.x {
				// Vertical line
				startY := from.y
				endY := to.y
				if startY > endY {
					startY, endY = endY, startY
				}
				for y := startY; y <= endY; y++ {
					rockSet[Position{from.x, y}] = struct{}{}
				}
			} else if from.y == to.y {
				// Horizontal line
				startX := from.x
				endX := to.x
				if startX > endX {
					startX, endX = endX, startX
				}
				for x := startX; x <= endX; x++ {
					rockSet[Position{x, from.y}] = struct{}{}
				}
			} else {
				return nil, 0, 0, 0, fmt.Errorf("non-straight line from %v to %v", from, to)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, 0, 0, 0, err
	}

	// Convert rockSet to slice
	var rocks []Position
	for pos := range rockSet {
		rocks = append(rocks, pos)
	}

	return rocks, minX, maxX, maxY, nil
}

// simulateSand runs the simulation for part 1 or part 2
func simulateSand(rocks []Position, minX, maxX, maxY int, part2 bool) int {
	// For part 2, we need to adjust the grid to include the floor
	var floorY int
	if part2 {
		floorY = maxY + 2
	}

	// To handle part 2, we might need to expand the grid horizontally
	// The maximum possible spread is floorY to the left and right
	if part2 {
		minX -= floorY
		maxX += floorY
	}

	width := maxX - minX + 1
	height := maxY + 1
	if part2 {
		height = floorY + 1
	}

	// Initialize grid
	grid := make([][]byte, height)
	for y := 0; y < height; y++ {
		grid[y] = make([]byte, width)
		for x := 0; x < width; x++ {
			grid[y][x] = '.' // air
		}
	}

	// Fill rocks
	for _, rock := range rocks {
		if rock.x < minX || rock.x > maxX || rock.y < 0 || rock.y >= height {
			continue
		}
		grid[rock.y][rock.x-minX] = '#' // rock
	}

	// For part 2, add the floor
	if part2 {
		for x := 0; x < width; x++ {
			grid[floorY][x] = '#' // floor
		}
	}

	sourceX := 500 - minX
	sourceY := 0
	grid[sourceY][sourceX] = '+' // source

	count := 0
	for {
		// Start sand at source
		sandX := sourceX
		sandY := sourceY

		for {
			if !part2 && sandY > maxY {
				// Part 1: Sand falls into the abyss
				return count
			}

			// Try to move down
			if sandY+1 < height && grid[sandY+1][sandX] == '.' {
				sandY++
				continue
			}

			// Try to move down-left
			if sandX-1 >= 0 && sandY+1 < height && grid[sandY+1][sandX-1] == '.' {
				sandX--
				sandY++
				continue
			}

			// Try to move down-right
			if sandX+1 < width && sandY+1 < height && grid[sandY+1][sandX+1] == '.' {
				sandX++
				sandY++
				continue
			}

			// Sand comes to rest
			if sandY == sourceY && sandX == sourceX {
				// Source is blocked
				count++
				return count
			}

			grid[sandY][sandX] = 'o'
			count++
			break
		}
	}
}
