package day02

import (
	"math"
	"os"
	"strconv"
	"strings"
)

func parseInputs() [][]int {
    inputs, err := os.ReadFile("inputs/day02.txt")
    if err != nil {
        println(err.Error())
    }

    lines := strings.Split(string(inputs), "\n")
    lines = lines[:len(lines)-1] // remove empty "line" due to final \n

    reports := make([][]int, 1000)

    for i, line := range lines {
        levels := strings.Split(line, " ")
        reports[i] = make([]int, len(levels))
        for j, level := range levels {
            reports[i][j], _ = strconv.Atoi(level)
        }
    }

    return reports
}

func allIncreasing(levels []int) bool {
    cur := levels[0]
    for _, level := range levels[1:] {
        if level <= cur {
            return false
        }
        cur = level
    }
    return true
}

func allDecreasing(levels []int) bool {
    cur := levels[0]
    for _, level := range levels[1:] {
        if level >= cur {
            return false
        }
        cur = level
    }
    return true
}

func isGradual(levels []int) bool {
    cur := levels[0]
    for _, level := range levels[1:] {
        delta := int(math.Abs(float64(cur - level)))
        if  delta > 3 || delta < 1 {
            return false
        }
        cur = level
    }
    return true
}

func reportIsSafe(report []int) bool {
    safe := true

    if (!allIncreasing(report) && !allDecreasing(report)) {
        safe = false
    }

    if (!isGradual(report)) {
        safe = false
    }

    return safe
}

func reportIsSuperSafe(report []int) bool {
    for i, _ := range report {
        sub := []int{}
        sub = append(sub, report[:i]...)
        sub = append(sub, report[i+1:]...)
        if reportIsSafe(sub) {
            return true
        }
    }

    return false
}

func evaluateReports(reports [][]int) (int, int) {
    safeReports := 0
    superSafeReports := 0
    for _, report := range reports {
        if reportIsSafe(report) { safeReports++ }
        if reportIsSuperSafe(report) { superSafeReports++ }
    }
    
    return safeReports, superSafeReports
}

func Run() (any, error) {
    reports := parseInputs()
    safe, superSafe := evaluateReports(reports)
    return []int{safe, superSafe}, nil
}
