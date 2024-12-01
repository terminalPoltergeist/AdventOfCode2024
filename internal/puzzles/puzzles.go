package puzzles

import (
	"fmt"

	"github.com/terminalPoltergeist/aoc24/pkg/day01"
)

var Solvers map[int]func() (any, error) = map[int]func() (any, error){
    1 : day01.Run,
}

func Solve(puzzle int) {
    if f := Solvers[puzzle]; f != nil {
        result, _ := f()

        fmt.Printf("The solution to puzzle %d is %v\n", puzzle, result)
    }
}
