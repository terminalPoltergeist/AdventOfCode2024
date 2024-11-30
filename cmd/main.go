package main

import (
	"github.com/terminalPoltergeist/aoc24/internal/args"
	"github.com/terminalPoltergeist/aoc24/internal/puzzles"
)

func main() {
    if err := args.ParseArgs(); err != nil {
        println(err.Error())
    }

    puzzles.Solve(args.Puzzle)
}
