package main

import (
	"fmt"
	"os"
	"regexp"
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

func toInts(ss []string) []int {
	ns := make([]int, len(ss))
	for i, s := range ss {
		n, err := strconv.Atoi(s)
		catch(err)
		ns[i] = n
	}
	return ns
}

var reJob = regexp.MustCompile(`(\d+)-(\d+),(\d+)-(\d+)`)

func part1(lines []string) {
	var count int
	for _, line := range lines {
		job := toInts(reJob.FindStringSubmatch(line)[1:])
		if job[0] <= job[2] && job[3] <= job[1] || job[2] <= job[0] && job[1] <= job[3] {
			count++
		}
	}
	fmt.Println("Part 1:", count)
}

func part2(lines []string) {
	var count int
	for _, line := range lines {
		job := toInts(reJob.FindStringSubmatch(line)[1:])
		if job[0] > job[3] || job[2] > job[1] {
			continue
		}
		count++
	}
	fmt.Println("Part 2:", count)
}
