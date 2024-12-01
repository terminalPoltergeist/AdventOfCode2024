package main

import (
	"github.com/terminalPoltergeist/aoc24/internal/args"
	"github.com/terminalPoltergeist/aoc24/internal/puzzles"
)

func main() {
    if err := args.ParseArgs(); err != nil {
        println(err.Error())
    } else {
        puzzles.Solve(args.Puzzle)
    }
}
