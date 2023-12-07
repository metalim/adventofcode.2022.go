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

type Dir struct {
	parent *Dir
	dirs   map[string]*Dir
	files  map[string]int
}

func (d *Dir) cd(name string) *Dir {
	if name == "/" {
		for d.parent != nil {
			d = d.parent
		}
		return d
	}

	if name == ".." {
		return d.parent
	}

	return d.dirs[name]
}

func (d *Dir) mkdir(path string) {
	d.dirs[path] = &Dir{
		parent: d,
		dirs:   map[string]*Dir{},
		files:  map[string]int{},
	}
}

func (d *Dir) touch(path string, size int) {
	d.files[path] = size
}

func (d *Dir) getSize() int {
	var total int
	for _, size := range d.files {
		total += size
	}
	for _, dir := range d.dirs {
		total += dir.getSize()
	}
	return total
}

func (d *Dir) sumDirsLessThan(maxSize int) int {
	var total int

	size := d.getSize()
	if size <= maxSize {
		total += size
	}

	for _, dir := range d.dirs {
		total += dir.sumDirsLessThan(maxSize)
	}
	return total
}

func (d *Dir) findDirAtLeast(minSize int) int {
	minDirSize := d.getSize()
	if minDirSize < minSize {
		return -1
	}
	for _, dir := range d.dirs {
		subdirSize := dir.findDirAtLeast(minSize)
		if subdirSize < 0 {
			continue
		}
		if minDirSize > subdirSize {
			minDirSize = subdirSize
		}
	}
	return minDirSize
}

func parseRoot(lines []string) *Dir {
	root := &Dir{
		dirs:  map[string]*Dir{},
		files: map[string]int{},
	}

	cwd := root
	for _, line := range lines {
		switch {
		case strings.HasPrefix(line, "$ cd "):
			cwd = cwd.cd(line[5:])
		case strings.HasPrefix(line, "$ ls"):
			// do nothing
		case strings.HasPrefix(line, "dir "):
			cwd.mkdir(line[4:])

		default:
			entry := strings.Fields(line)
			size, err := strconv.Atoi(entry[0])
			catch(err)
			cwd.touch(entry[1], size)
		}
	}
	return root
}
func part1(lines []string) {
	root := parseRoot(lines)
	fmt.Println("Part 1:", root.sumDirsLessThan(1e5))
}

func part2(lines []string) {
	const driveSize = 7e7
	const reqFree = 3e7

	root := parseRoot(lines)
	minFolderSize := root.findDirAtLeast(reqFree - driveSize + root.getSize())
	fmt.Println("Part 2:", minFolderSize)
}
