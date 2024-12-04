package puzzles

import (
	"fmt"

	"github.com/terminalPoltergeist/aoc24/pkg/day01"
	"github.com/terminalPoltergeist/aoc24/pkg/day02"
	"github.com/terminalPoltergeist/aoc24/pkg/day03"
)

var Solvers map[int]func() (any, error) = map[int]func() (any, error){
    1 : day01.Run,
    2 : day02.Run,
    3 : day03.Run,
}

func Solve(puzzle int) {
    if f, ok := Solvers[puzzle]; ok {
        result, _ := f()

        fmt.Printf("The solution to puzzle %d is %v\n", puzzle, result)
    } else if puzzle == 0 {
        for k := range Solvers {
            result, _ := Solvers[k]()
            fmt.Printf("The solution to puzzle %d is %v\n", k, result)
        }
    }
}
