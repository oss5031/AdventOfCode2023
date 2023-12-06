package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("usage: %s <file>\n", os.Args[0])
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("%s not found\n", os.Args[1])
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	part1, part2 := solve(lines)
	fmt.Printf("%d\n%d\n", part1, part2)
}

func parse(entryStr string) entry {
	fields := strings.Fields(entryStr)
	dst, _ := strconv.Atoi(fields[0])
	src, _ := strconv.Atoi(fields[1])
	rng, _ := strconv.Atoi(fields[2])
	return entry{src, src + (rng - 1), dst}
}

type entry struct {
	min, max, dst int
}

func solve(lines []string) (part1 int, part2 int) {
	var seeds []int
	for _, seedStr := range strings.Fields(lines[0][strings.Index(lines[0], ":")+1:]) {
		s, _ := strconv.Atoi(seedStr)
		seeds = append(seeds, s)
	}

	var maps [][]entry
	for i := 1; i < len(lines); i++ {
		var entries []entry
		if len(lines[i]) == 0 {
			for j := 2; j+i < len(lines); j++ {
				curr := lines[i+j]
				if len(curr) == 0 {
					i += j - 1
					break
				}
				entries = append(entries, parse(curr))
			}
			sort.Slice(entries, func(i, j int) bool { return entries[i].min < entries[j].min })
			maps = append(maps, entries)
		}
	}

	getLoc := func(seed int) int {
		val := seed
		for _, m := range maps {
			for _, e := range m {
				if val >= e.min && val <= e.max {
					val = e.dst + (val - e.min)
					break
				}
			}
		}
		return val
	}

	part1 = math.MaxInt
	for _, s := range seeds {
		loc := getLoc(s)
		if part1 > loc {
			part1 = loc
		}
	}

	part2 = math.MaxInt
	for i := 0; i < len(seeds)-1; i += 2 {
		for j := 0; j < seeds[i+1]; j++ {
			loc := getLoc(seeds[i] + j)
			if part2 > loc {
				part2 = loc
			}
		}
	}

	return
}
