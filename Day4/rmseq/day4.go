package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("usage: %s <path> <part>\n", os.Args[0])
		os.Exit(1)
	}

	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	switch os.Args[2] {
	case "1":
		fmt.Println(part1(string(data)))
	case "2":
		fmt.Println(part2(string(data)))
	default:
		fmt.Printf("usage: %s <path> <part>\n", os.Args[0])
		os.Exit(1)
	}
}

func part1(input string) int {
	var res int
	for _, ln := range strings.Split(input, "\n") {
		if m := matches(fields(ln)); m >= 1 {
			res += m << (m - 1)
		}
	}
	return res
}

func part2(input string) int {
	lns := strings.Split(input, "\n")
	cards, total := make([]int, len(lns)), 0
	for i := 0; i < len(lns); i++ {
		cards[i]++
		total += cards[i]
		for j := 1; j <= matches(fields(lns[i])); j++ {
			cards[i+j] += cards[i]
		}
	}
	return total
}

func fields(ln string) ([]string, []string) {
	f := strings.Split(strings.Split(ln, ":")[1], "|")
	return strings.Fields(f[0]), strings.Fields(f[1])
}

func matches(x, y []string) int {
	set := make(map[string]bool)
	for _, s := range y {
		set[s] = true
	}
	var res int
	for _, s := range x {
		if _, has := set[s]; has {
			res++
		}
	}
	return res
}
