package day03

import (
	"os"
	"regexp"
	"strconv"
	"strings"
)

func parseInputs() []string {
    inputs, err := os.ReadFile("inputs/day03.txt")
    if err != nil {
        println(err.Error())
    }

    lines := strings.Split(string(inputs), "\n")
    lines = lines[:len(lines)-1] // remove empty "line" due to final \n

    return lines
}

func extractOps(lines []string) []string {
    ops := []string{}
    for _, line := range lines {
        regex := regexp.MustCompile(`(mul\(\d{1,3},\d{1,3}\))`)
        matches := regex.FindAllString(line, -1)
        ops = append(ops, matches...)
    }
    return ops
}

func extractEnabledOps(line string) []string {
    regex := regexp.MustCompile(`(?:don't\(\).*?do\(\))|(?:don't\(\).*)`) // matches sections of diabled operations

    filtered := regex.ReplaceAllString(line, "")

    regex = regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
    matches := regex.FindAllString(filtered, -1)

    return matches
}

func processOps(ops []string) int {
    total := 0
    for _, op := range ops {
        regex := regexp.MustCompile(`\((\d{1,3}),(\d{1,3})\)`)
        args := regex.FindAllStringSubmatch(op, -1)
        arg1, _ := strconv.Atoi(args[0][1])
        arg2, _ := strconv.Atoi(args[0][2])
        total += arg1 * arg2
    }
    return total
}

func Run() (any, error) {
    lines := parseInputs()
    ops := extractOps(lines)
    result1 := processOps(ops)
    enabledOps := extractEnabledOps(strings.Join(lines, ""))
    result2 := processOps(enabledOps)
    return []int{result1,result2}, nil
}
