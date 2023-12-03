package main

import (
	"fmt"
	"os"
	"strconv"
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
		fmt.Printf("usage: %s <path> <part>", os.Args[0])
		os.Exit(1)
	}
}

func part1(input string) int {
	var res int
NextGame:
	for i, ln := range strings.Split(input, "\n") {
		for _, handStr := range strings.Split(strings.Split(ln, ":")[1], ";") {
			if !parseHand(handStr).IsPossible(12, 13, 14) {
				continue NextGame
			}
		}
		res += i + 1
	}
	return res
}

func part2(input string) int {
	var res int
	for _, ln := range strings.Split(input, "\n") {
		var reds, greens, blues int
		for _, handStr := range strings.Split(strings.Split(ln, ":")[1], ";") {
			hand := parseHand(handStr)
			reds, greens, blues = max(reds, hand.Reds), max(greens, hand.Greens), max(blues, hand.Blues)
		}
		res += reds * greens * blues
	}
	return res
}

var (
	handSetter = map[string]func(h *Hand, val int){
		"red":   func(h *Hand, val int) { h.Reds = val },
		"blue":  func(h *Hand, val int) { h.Blues = val },
		"green": func(h *Hand, val int) { h.Greens = val },
	}
)

// Pre: well-behaved input
func parseHand(s string) *Hand {
	h := &Hand{}
	for _, elem := range strings.Split(s, ",") {
		f := strings.Fields(elem)
		val, _ := strconv.Atoi(f[0])
		handSetter[f[1]](h, val)
	}
	return h
}

type Hand struct {
	Reds   int
	Blues  int
	Greens int
}

func (h *Hand) IsPossible(reds, greens, blues int) bool {
	return h.Reds <= reds && h.Blues <= greens && h.Greens <= blues
}
