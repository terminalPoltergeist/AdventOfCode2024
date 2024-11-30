package args

import (
	"flag"
	"fmt"
	"strconv"
	"time"
)

// AOC 2024 start date
var START = time.Date(2024, 12, 1, 0, 0, 0, 0, time.Local)

var Puzzle int

func GetDay() (int, error) {
    if START.After(time.Now()) {
        timeUntilStart := time.Since(START)
        return -1, fmt.Errorf("AOC 2024 hasn't started yet. Check back in %s.", timeUntilStart.Abs().Round(time.Second).String())
    }

    return time.Now().Day(), nil
}

func ParseArgs() error {
    day, err := GetDay()

    Puzzle = day

    if err != nil {
        return err
    }

    puzzle := flag.String("puzzle", "0","Which puzzle to run. Comma separated numbers starting at 1.")

    flag.Parse()

    Puzzle, err = strconv.Atoi(*puzzle)

    if err != nil {
        return err
    }

    return nil
}
