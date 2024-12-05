package day04

import (
	"os"
	"slices"
	"strings"
)

const target = "XMAS"

func parseInputs() []string {
    inputs, err := os.ReadFile("inputs/day04.txt")
    if err != nil {
        println(err.Error())
    }

    lines := strings.Split(string(inputs), "\n")
    lines = lines[:len(lines)-1] // remove empty "line" due to final \n

    return lines
}

func count(list []string, word string) int {
    total := 0
    for _, w := range list {
        if w == word {total++}
    }
    return total
}

func countXmases(lines []string, y int, x int) int {
    // y is the line index the X appears in
    // x is the character index the X appears in y line
    // this implementation disappoints me.. I will refactor!

    xmases := 0

    var words []string

    // left to right
    if len(lines[y]) >= x+4 {
        words = append(words, lines[y][x:x+4])
    }
    // right to left
    if x >= 3 {
        word := []byte(lines[y][x-3:x+1])
        slices.Reverse(word)
        words = append(words, string(word))
    }
    // top to bottom
    if len(lines) > y+3 {
        word := []byte{lines[y][x], lines[y+1][x], lines[y+2][x], lines[y+3][x]}
        words = append(words, string(word))
    }
    // bottom to top
    if y >= 3 {
        word := []byte{lines[y][x], lines[y-1][x], lines[y-2][x], lines[y-3][x]}
        words = append(words, string(word))
    }
    // upper-left diagonal
    if y >= 3 && x >= 3 {
        word := []byte{lines[y][x], lines[y-1][x-1], lines[y-2][x-2], lines[y-3][x-3]}
        words = append(words, string(word))
    }
    // upper-right diagonal
    if y >= 3 && x+4 <= len(lines[y]) {
        word := []byte{lines[y][x], lines[y-1][x+1], lines[y-2][x+2], lines[y-3][x+3]}
        words = append(words, string(word))
    }
    // lower-left diagonal
    if len(lines) > y+3 && x >= 3 {
        word := []byte{lines[y][x], lines[y+1][x-1], lines[y+2][x-2], lines[y+3][x-3]}
        words = append(words, string(word))
    }
    // lower-right diagonal
    if len(lines) > y+3 && x+4 <= len(lines[y]) {
        word := []byte{lines[y][x], lines[y+1][x+1], lines[y+2][x+2], lines[y+3][x+3]}
        words = append(words, string(word))
    }

    xmases += count(words, target)

    return xmases
}

func findXmas(lines []string) int {
    total := 0

    for i, line := range lines {
        for j, char := range line {
            if string(char) == "X" {
                total += countXmases(lines, i, j)
            }
        }
    }

    return total
}

func isMasX(lines []string, y int, x int) bool {
    // y is the line index the X appears in
    // x is the character index the X appears in y line
    // this implementation disappoints me.. I will refactor!

    masx := false

    if (y >= 1 && x >= 1) && (y < len(lines)-1 && x < len(lines[y])-1) {
        // coordinates are within required bounds
        cross1 := string([]byte{lines[y-1][x-1], lines[y][x], lines[y+1][x+1]})
        cross2 := string([]byte{lines[y+1][x-1], lines[y][x], lines[y-1][x+1]})

        if (cross1 == "MAS" || cross1 == "SAM") && (cross2 == "MAS" || cross2 == "SAM") {
            masx = true
        }
    }

    return masx
}

func countMasXes(lines []string) int {
    total := 0

    for i, line := range lines {
        for j, char := range line {
            if string(char) == "A" {
                if isMasX(lines, i, j) {
                    total++
                }
            }
        }
    }

    return total
}

func Run() (any, error) {
    lines := parseInputs()
    result1 := findXmas(lines)
    result2 := countMasXes(lines)
    return []int{result1, result2}, nil
}
