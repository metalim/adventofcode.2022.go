package main

import (
	"fmt"
	"os"
	"strconv"
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

func simulate1(lines []string) int {
	visited := map[Pos]bool{}
	var h, t Pos
	visited[t] = true
	for _, line := range lines {
		move := strings.Fields(line)
		distance, err := strconv.Atoi(move[1])
		catch(err)

		for ; distance > 0; distance-- {
			oldH := h
			switch move[0] {
			case "U":
				h.y++
			case "D":
				h.y--
			case "L":
				h.x--
			case "R":
				h.x++
			}
			dx := h.x - t.x
			dy := h.y - t.y
			if dx < -1 || dx > 1 || dy < -1 || dy > 1 {
				t = oldH
				visited[t] = true
			}
		}
	}
	draw(visited)
	return len(visited)
}

func simulate(lines []string, ropeLen int) int {
	maxDist := ropeLen - 1
	visited := map[Pos]bool{}
	var h, t Pos
	visited[t] = true
	for _, line := range lines {
		move := strings.Fields(line)
		distance, err := strconv.Atoi(move[1])
		catch(err)

		for ; distance > 0; distance-- {
			switch move[0] {
			case "U":
				h.y++
			case "D":
				h.y--
			case "L":
				h.x--
			case "R":
				h.x++
			}
			dx := h.x - t.x
			dy := h.y - t.y
			if dx < -maxDist || dx > maxDist || dy < -maxDist || dy > maxDist {
				// move the tail
				t.x += sign(dx)
				t.y += sign(dy)
				visited[t] = true
			}
		}
	}
	draw(visited)
	return len(visited)
}

func simulateRope(lines []string, ropeLen int) int {
	visited := map[Pos]bool{}
	knots := make([]Pos, ropeLen)
	knotsOld := make([]Pos, ropeLen)
	visited[knots[ropeLen-1]] = true
	for _, line := range lines {
		move := strings.Fields(line)
		distance, err := strconv.Atoi(move[1])
		catch(err)

		for ; distance > 0; distance-- {
			copy(knotsOld, knots)
			switch move[0] {
			case "U":
				knots[0].y++
			case "D":
				knots[0].y--
			case "L":
				knots[0].x--
			case "R":
				knots[0].x++
			}
			for i := 1; i < ropeLen; i++ {
				dx := knots[i].x - knots[i-1].x
				dy := knots[i].y - knots[i-1].y
				if dx < -1 || dx > 1 || dy < -1 || dy > 1 {
					knots[i] = knotsOld[i-1]
				}
			}
			visited[knots[ropeLen-1]] = true
		}
	}
	draw(visited)
	return len(visited)
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

func draw(visited map[Pos]bool) {
	var min, max Pos
	for pos := range visited {
		if pos.x < min.x {
			min.x = pos.x
		}
		if pos.y < min.y {
			min.y = pos.y
		}
		if pos.x > max.x {
			max.x = pos.x
		}
		if pos.y > max.y {
			max.y = pos.y
		}
	}
	for y := max.y; y >= min.y; y-- {
		for x := min.x; x <= max.x; x++ {
			if visited[Pos{x, y}] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
func part1(lines []string) {
	fmt.Println("Part 1:", simulate1(lines))
	fmt.Println("Part 1.1:", simulate(lines, 2))
	fmt.Println("Part 1.2:", simulateRope(lines, 2))
}

func part2(lines []string) {
	fmt.Println("Part 2.1:", simulate(lines, 10))
	fmt.Println("Part 2.2:", simulateRope(lines, 10))
}
