package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
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

func solve(lines []string) (part1, part2 int) {
	times, t := parse(lines[0])
	dist, d := parse(lines[1])
	part1 = 1
	for i := 0; i < len(times); i++ {
		upper, lower, _ := solveQuadratic(1.0, -float64(times[i]), float64(dist[i])+1)
		part1 *= int(math.Floor(upper)-math.Ceil(lower)) + 1
	}
	upper, lower, _ := solveQuadratic(1.0, -float64(t), float64(d)+1)
	part2 = int(math.Floor(upper)-math.Ceil(lower)) + 1
	return
}

func parse(s string) (vals []int, val int) {
	s = s[strings.Index(s, ":")+1:]

	for _, f := range strings.Fields(s) {
		v, _ := strconv.Atoi(f)
		vals = append(vals, v)
	}

	s = strings.ReplaceAll(s, " ", "")
	val, _ = strconv.Atoi(s)
	return
}

func solveQuadratic(a, b, c float64) (x1, x2 float64, has bool) {
	d := math.Pow(b, 2) - (4 * a * c)
	has = d >= 0
	if !has {
		return
	}
	if d > 0 {
		x1 = (-b + math.Sqrt(d)) / (2 * a)
		x2 = (-b - math.Sqrt(d)) / (2 * a)
	} else {
		fmt.Println(a, b, c)
		x1 = -b / (2 * a)
		x2 = x1
	}
	return
}
