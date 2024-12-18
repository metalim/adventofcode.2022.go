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
